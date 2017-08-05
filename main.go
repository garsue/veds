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
	flag.StringVar(&cnf.dsHost, "host", "", "Datastore emulator host")
	flag.StringVar(&cnf.projectID, "id", "", "Project ID")
}

func main() {
	flag.Parse()

	if err := start(cnf); err != nil {
		log.Fatal(err)
	}
	log.Println("Finish")
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

	query := datastore.NewQuery("__namespace__").KeysOnly()
	keys, err := client.GetAll(ctx, query, nil)
	if err != nil {
		return err
	}
	for i, key := range keys {
		log.Println(i, key.Name, key)
	}

	return nil
}
