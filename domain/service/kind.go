package service

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/garsue/veds/application"
	"github.com/garsue/veds/domain/repository"
)

// Kinds returns used keys in datastore
func Kinds(ctx context.Context, app *application.App) ([]string, error) {
	client, err := datastore.NewClient(ctx, app.Config.ProjectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	kinds, err := repository.Kinds(ctx, client)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, kind := range kinds {
		name := kind.Name
		result = append(result, name)
	}
	return result, nil
}
