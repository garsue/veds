package service

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/garsue/veds/application"
	"github.com/garsue/veds/domain/repository"
)

// Properties returns used properties in datastore
func Properties(ctx context.Context, app *application.App) (map[string][]string, error) {
	client, err := datastore.NewClient(ctx, app.Config.ProjectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	properties, err := repository.Properties(ctx, client)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, property := range properties {
		kind := property.Parent.Name
		name := property.Name
		result[kind] = append(result[kind], name)
	}
	return result, nil
}
