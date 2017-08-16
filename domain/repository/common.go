package repository

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Entities returns all keys and entries
func Entities(ctx context.Context, client *datastore.Client, kind string) ([]*datastore.Key, []datastore.PropertyList, error) {
	query := datastore.NewQuery(kind)
	var dst []datastore.PropertyList
	keys, err := client.GetAll(ctx, query, &dst)
	if err != nil {
		return nil, nil, err
	}
	return keys, dst, err
}
