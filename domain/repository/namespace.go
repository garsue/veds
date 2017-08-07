package repository

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Namespaces returns used namespaces in datastore
func Namespaces(ctx context.Context, client *datastore.Client) ([]*datastore.Key, error) {
	query := datastore.NewQuery("__namespace__").KeysOnly()
	return client.GetAll(ctx, query, nil)
}
