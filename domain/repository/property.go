package repository

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Properties returns used properties in datastore
func Properties(ctx context.Context, client *datastore.Client) ([]*datastore.Key, error) {
	query := datastore.NewQuery("__property__").KeysOnly()
	return client.GetAll(ctx, query, nil)
}
