package bosh

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Username     string
	Password     string
	DirectorHost string
}

func (client *Client) FetchDeployments() (Deployments, error) {
	var deployments Deployments

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments", client.DirectorHost), nil)
	if err != nil {
		return deployments, err
	}
	request.SetBasicAuth(client.Username, client.Password)

	httpClient := http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	response, err := httpClient.Do(request)
	if err != nil {
		return Deployments{}, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&deployments)
	return deployments, err
}
