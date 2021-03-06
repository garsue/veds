package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/garsue/veds/application"
	"github.com/garsue/veds/application/handler"
)

var cnf = &application.Config{}

func init() {
	flag.StringVar(&cnf.ProjectID, "id", "", "Project ID")
	flag.StringVar(&cnf.Public, "public", "public", "Public directory name")
}

func main() {
	app, err := application.NewApp(cnf)
	if err != nil {
		log.Fatal(err)
	}

	mux := handler.NewHandler(app)
	log.Fatal(http.ListenAndServe(":8090", mux))
}
