package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/svett/cf-cloud-watch/api"
	"github.com/svett/cf-cloud-watch/middleware"
)

var (
	addr      string
	loglevel  string
	gracetime time.Duration
)

func init() {
	// Commmand-line arguments
	flag.StringVar(&addr, "addr", ":8080", "Sets the server listening host and port (format address:port)")
	flag.StringVar(&loglevel, "loglevel", "debug", "Logging level")
	flag.DurationVar(&gracetime, "gracetime", 2*time.Second, "Sets the server grace time")
}

func main() {
	// Read configuration
	configure()

	// Router initialization
	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(middleware.NewLogger(logrus.StandardLogger()))
	server.Use(negroni.NewStatic(http.Dir("public")))

	// Routes configuration
	router := mux.NewRouter()
	router.Handle("/api/v1/deployments", &api.Bosh{})
	server.UseHandler(router)

	logrus.Printf("Server is listening on %s", addr)
	graceful.Run(addr, gracetime, server)
}

func configure() {
	logrus.Print("Reading configuration")

	flag.Parse()

	// Configure logging level
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(level)

	logrus.Print("Reading configuration completed")
}

func FileServer(name string) http.Handler {
	return http.FileServer(http.Dir(name))
}

func StaticServer(name string) http.Handler {
	fs := http.FileServer(http.Dir(name))
	return http.StripPrefix(fmt.Sprintf("/%s/", name), fs)
}
