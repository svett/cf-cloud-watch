# Cloudwatch

Presently, the home page shows a dashboard of the current BOSH deployments and detailed
information about them:

- stemcell version
- release verion
- date of deployment

# Prerequisite

- [Golang](http://golang.org/)

# Installation

```sh
go get github.com/svett/cf-cloud-watch/...
```

# Running

Presently, the Cloudwatch works with `bosh-lite` by using the following defaults:

- IP address: 192.168.50.4
- Credentials (username: `admin`, password: `admin`)

If you want to change them, you should change `bosh/bosh.go` file.

```sh
go run cmd/cf-cloud-watch/main.go
```

# Remaining work

- the dashboard is not auto-update
- the deployments do not have good UI/UX




