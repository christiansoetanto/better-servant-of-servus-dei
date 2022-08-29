package fstore

import (
	"cloud.google.com/go/firestore"
	"context"
)

func Init(ctx context.Context) (*firestore.Client, error) {
	projectID := "youtube-title-updater-340409"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
