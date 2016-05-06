package api

import (
	"encoding/json"
	"net/http"

	"github.com/svett/cf-cloud-watch/bosh"
)

type Bosh struct {
}

func (b *Bosh) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	client := &bosh.Client{Username: "admin", Password: "admin", DirectorHost: "https://192.168.50.4:25555"}
	deployments, err := client.FetchDeployments()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(writer).Encode(&deployments); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
}
