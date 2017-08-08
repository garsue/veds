package repository

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Kinds returns used namespaces in datastore
func Kinds(ctx context.Context, client *datastore.Client) ([]*datastore.Key, error) {
	query := datastore.NewQuery("__kind__").KeysOnly()
	return client.GetAll(ctx, query, nil)
}
