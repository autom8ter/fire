package functions

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
)

type BucketHandlerFunc func(b *storage.BucketHandle) error
type ObjectHandlerFunc func(b *storage.ObjectHandle) error
type CollectionHandlerFunc func(b *firestore.CollectionRef) error
type DocumentHandlerFunc func(b *firestore.DocumentRef) error
