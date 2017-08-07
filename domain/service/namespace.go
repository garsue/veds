package service

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/garsue/veds/application"
	"github.com/garsue/veds/domain/repository"
)

const (
	defaultNamespaceDisplayName = "(default)"
)

// Namespaces returns used namespaces in datastore
func Namespaces(ctx context.Context, app *application.App) ([]string, error) {
	client, err := datastore.NewClient(ctx, app.Config.ProjectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	keys, err := repository.Namespaces(ctx, client)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, key := range keys {
		name := key.Name
		if name == "" {
			name = defaultNamespaceDisplayName
		}
		result = append(result, name)
	}
	log.Println("size", len(result))
	return result, nil
}
