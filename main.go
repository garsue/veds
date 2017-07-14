package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/datastore"
)

type config struct {
	dsHost    string
	projectID string
}

var cnf config

func init() {
	flag.StringVar(&cnf.dsHost, "datastore-host", "", "Datastore emulator host")
	flag.StringVar(&cnf.projectID, "project-id", "", "Project ID")
}

func main() {
	flag.Parse()

	if err := start(cnf); err != nil {
		log.Fatal(err)
	}
}

func start(cnf config) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, cnf.projectID)
	if err != nil {
		return err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	log.Println(client)
	return nil
}
