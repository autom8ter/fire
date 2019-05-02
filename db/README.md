# db
--
    import "github.com/autom8ter/fire/db"


## Usage

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error)
```

#### func (*Client) Bucket

```go
func (c *Client) Bucket(ctx context.Context, cat driver.Categorizer) *storage.BucketHandle
```

#### func (*Client) BucketObject

```go
func (c *Client) BucketObject(ctx context.Context, cat driver.Categorizer) *storage.ObjectIterator
```

#### func (*Client) Buckets

```go
func (c *Client) Buckets(ctx context.Context) *storage.BucketIterator
```

#### func (*Client) Collection

```go
func (c *Client) Collection(ctx context.Context, cat driver.Categorizer) *firestore.CollectionRef
```

#### func (*Client) Collections

```go
func (c *Client) Collections(ctx context.Context) *firestore.CollectionIterator
```

#### func (*Client) CopyFromObject

```go
func (c *Client) CopyFromObject(ctx context.Context, from driver.Grouping, to driver.Grouping) *storage.Copier
```

#### func (*Client) CopyObjectTo

```go
func (c *Client) CopyObjectTo(ctx context.Context, dst io.Writer, grp driver.Grouping) error
```

#### func (*Client) CopyToObjectFrom

```go
func (c *Client) CopyToObjectFrom(ctx context.Context, from io.Reader, grp driver.Grouping) error
```

#### func (*Client) CreateBucket

```go
func (c *Client) CreateBucket(ctx context.Context, cat driver.Categorizer) error
```

#### func (*Client) CreateDoc

```go
func (c *Client) CreateDoc(ctx context.Context, group driver.Grouping, data map[string]interface{}) error
```

#### func (*Client) DeleteBucket

```go
func (c *Client) DeleteBucket(ctx context.Context, cat driver.Categorizer) error
```

#### func (*Client) DeleteDoc

```go
func (c *Client) DeleteDoc(ctx context.Context, group driver.Grouping) error
```

#### func (*Client) DeleteObject

```go
func (c *Client) DeleteObject(ctx context.Context, group driver.Grouping) error
```

#### func (*Client) DocData

```go
func (c *Client) DocData(ctx context.Context, group driver.Grouping) (map[string]interface{}, error)
```

#### func (*Client) DocDataAt

```go
func (c *Client) DocDataAt(ctx context.Context, group driver.Grouping, key string) (interface{}, error)
```

#### func (*Client) DocSnapshot

```go
func (c *Client) DocSnapshot(ctx context.Context, group driver.Grouping) (*firestore.DocumentSnapshot, error)
```

#### func (*Client) Document

```go
func (c *Client) Document(ctx context.Context, group driver.Grouping) *firestore.DocumentRef
```

#### func (*Client) Documents

```go
func (c *Client) Documents(ctx context.Context, cat driver.Categorizer) *firestore.DocumentIterator
```

#### func (*Client) GetObjectMetadata

```go
func (c *Client) GetObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (map[string]string, error)
```

#### func (*Client) HandleBucket

```go
func (c *Client) HandleBucket(ctx context.Context, cat driver.Categorizer, fn functions.BucketHandlerFunc) error
```

#### func (*Client) HandleCollection

```go
func (c *Client) HandleCollection(ctx context.Context, cat driver.Categorizer, fn functions.CollectionHandlerFunc) error
```

#### func (*Client) HandleDocument

```go
func (c *Client) HandleDocument(ctx context.Context, group driver.Grouping, fn functions.DocumentHandlerFunc) error
```

#### func (*Client) HandleObject

```go
func (c *Client) HandleObject(ctx context.Context, group driver.Grouping, fn functions.ObjectHandlerFunc) error
```

#### func (*Client) MarshalDocTo

```go
func (c *Client) MarshalDocTo(ctx context.Context, group driver.Grouping, obj interface{}) error
```

#### func (*Client) Object

```go
func (c *Client) Object(ctx context.Context, group driver.Grouping) *storage.ObjectHandle
```

#### func (*Client) ObjectAttributes

```go
func (c *Client) ObjectAttributes(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error)
```

#### func (*Client) ObjectReader

```go
func (c *Client) ObjectReader(ctx context.Context, grp driver.Grouping) (*storage.Reader, error)
```

#### func (*Client) ObjectWriter

```go
func (c *Client) ObjectWriter(ctx context.Context, grp driver.Grouping) *storage.Writer
```

#### func (*Client) ObjectsBucketName

```go
func (c *Client) ObjectsBucketName(ctx context.Context, grp driver.Grouping) string
```

#### func (*Client) SetDocData

```go
func (c *Client) SetDocData(ctx context.Context, group driver.Grouping, data map[string]interface{}, merge bool) error
```

#### func (*Client) UpdateBucket

```go
func (c *Client) UpdateBucket(ctx context.Context, cat driver.Categorizer, attr storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error)
```

#### func (*Client) UpdateDocField

```go
func (c *Client) UpdateDocField(ctx context.Context, group driver.Grouping, key string, value string) error
```

#### func (*Client) UpdateObjectMetadata

```go
func (c *Client) UpdateObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error)
```
