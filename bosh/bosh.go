package bosh

import "time"

type Deployments []Deployment

type Deployment struct {
	Name       string      `json:"name"`
	DeployDate time.Time   `json:"deploy_date"`
	Releases   []*Release  `json:"releases"`
	Stemcells  []*Stemcell `json:"stemcells"`
}

type Release struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Stemcell struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
