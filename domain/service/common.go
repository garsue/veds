package service

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"

	"github.com/garsue/veds/application"
	"github.com/garsue/veds/domain/repository"
)

// Entities returns entries associated with the kind
func Entities(ctx context.Context, app *application.App, kind string) (map[string]interface{}, error) {
	client, err := datastore.NewClient(ctx, app.Config.ProjectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	keys, entities, err := repository.Entities(ctx, client, kind)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	for i, key := range keys {
		result[key.String()] = entities[i]
	}
	return result, nil
}
