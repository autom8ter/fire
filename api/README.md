# api
--
    import "github.com/autom8ter/fire/api"


## Usage

#### type API

```go
type API interface {
	Publisher
	Documentor
	Objector
	Bucketor
	Collector
	Handler
	Logger
}
```


#### type Bucketor

```go
type Bucketor interface {
	Bucket(ctx context.Context, cat driver.Categorizer) *storage.BucketHandle
	CreateBucket(ctx context.Context, cat driver.Categorizer) error
	DeleteBucket(ctx context.Context, cat driver.Categorizer) error
	UpdateBucket(ctx context.Context, cat driver.Categorizer, attr storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
	BucketObject(ctx context.Context, cat driver.Categorizer) *storage.ObjectIterator
	Buckets(ctx context.Context) *storage.BucketIterator
	HandleBucket(ctx context.Context, cat driver.Categorizer, fn functions.BucketHandlerFunc) error
}
```


#### type Collector

```go
type Collector interface {
	Collection(ctx context.Context, cat driver.Categorizer) *firestore.CollectionRef
	Collections(ctx context.Context) *firestore.CollectionIterator
}
```


#### type Documentor

```go
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
```


#### type Handler

```go
type Handler interface {
	HandleObject(ctx context.Context, group driver.Grouping, fn functions.ObjectHandlerFunc) error
	HandleCollection(ctx context.Context, cat driver.Categorizer, fn functions.CollectionHandlerFunc) error
	HandleDocument(ctx context.Context, group driver.Grouping, fn functions.DocumentHandlerFunc) error
}
```


#### type Logger

```go
type Logger interface {
	Err(err error)
}
```


#### type Objector

```go
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
```


#### type Publisher

```go
type Publisher interface {
	GetTopic(ctx context.Context, cat driver.Categorizer) (*pubsub.Topic, error)
	Publish(ctx context.Context, message driver.Message) (string, error)
}
```
