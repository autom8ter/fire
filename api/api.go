//go:generate godocdown -o README.md

package api

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"context"
	"github.com/autom8ter/api/driver"
	"github.com/autom8ter/fire/functions"
	"io"
)

type Publisher interface {
	GetTopic(ctx context.Context, cat driver.Categorizer) (*pubsub.Topic, error)
	Publish(ctx context.Context, message driver.Message) (string, error)
}

type Documentor interface {
	Documents(ctx context.Context, cat driver.Categorizer) *firestore.DocumentIterator
	DocSnapshot(ctx context.Context, group driver.Grouping) (*firestore.DocumentSnapshot, error)
	Document(ctx context.Context, group driver.Grouping) *firestore.DocumentRef
	MarshalDocTo(ctx context.Context, group driver.Grouping, obj interface{}) error
	DocDataAt(ctx context.Context, group driver.Grouping, key string) (interface{}, error)
	DocData(ctx context.Context, group driver.Grouping) (map[string]interface{}, error)
	UpdateDocField(ctx context.Context, group driver.Grouping, key string, value string) error
	CreateDoc(ctx context.Context, group driver.Grouping, data map[string]interface{}) error
	DeleteDoc(ctx context.Context, group driver.Grouping) error
	SetDocData(ctx context.Context, group driver.Grouping, data map[string]interface{}, merge bool) error
}

type Objector interface {
	Object(ctx context.Context, group driver.Grouping) *storage.ObjectHandle
	CopyFromObject(ctx context.Context, from driver.Grouping, to driver.Grouping) *storage.Copier
	DeleteObject(ctx context.Context, group driver.Grouping) error
	UpdateObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error)
	ObjectAttributes(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error)
	GetObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (map[string]string, error)
	ObjectsBucketName(ctx context.Context, grp driver.Grouping) string
	ObjectWriter(ctx context.Context, grp driver.Grouping) *storage.Writer
	ObjectReader(ctx context.Context, grp driver.Grouping) (*storage.Reader, error)
	CopyObjectTo(ctx context.Context, dst io.Writer, grp driver.Grouping) error
	CopyToObjectFrom(ctx context.Context, from io.Reader, grp driver.Grouping) error
}

type Bucketor interface {
	Bucket(ctx context.Context, cat driver.Categorizer) *storage.BucketHandle
	CreateBucket(ctx context.Context, cat driver.Categorizer) error
	DeleteBucket(ctx context.Context, cat driver.Categorizer) error
	UpdateBucket(ctx context.Context, cat driver.Categorizer, attr storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
	BucketObject(ctx context.Context, cat driver.Categorizer) *storage.ObjectIterator
	Buckets(ctx context.Context) *storage.BucketIterator
	HandleBucket(ctx context.Context, cat driver.Categorizer, fn functions.BucketHandlerFunc) error
}

type Collector interface {
	Collection(ctx context.Context, cat driver.Categorizer) *firestore.CollectionRef
	Collections(ctx context.Context) *firestore.CollectionIterator
}

type Handler interface {
	HandleObject(ctx context.Context, group driver.Grouping, fn functions.ObjectHandlerFunc) error
	HandleCollection(ctx context.Context, cat driver.Categorizer, fn functions.CollectionHandlerFunc) error
	HandleDocument(ctx context.Context, group driver.Grouping, fn functions.DocumentHandlerFunc) error
}

type Logger interface {
	Err(err error)
}
