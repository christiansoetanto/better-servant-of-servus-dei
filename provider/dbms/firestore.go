package dbms

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

type FirestoreDb interface {
	HelloWorld(ctx context.Context) error
}
type firestoreDb struct {
	c *firestore.Client
}

func getFirestoreDbObj(c *firestore.Client) FirestoreDb {
	return &firestoreDb{
		c: c,
	}
}

func (firestoreDb *firestoreDb) HelloWorld(ctx context.Context) error {
	iter := firestoreDb.c.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			return err
		}
		if err == iterator.Done {
			break
		}
		fmt.Println(doc.Data())
	}
	return nil
}
