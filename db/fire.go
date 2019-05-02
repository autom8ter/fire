//go:generate godocdown -o README.md

package db

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	"github.com/autom8ter/api/driver"
	"github.com/autom8ter/fire/functions"
	"google.golang.org/api/option"
	"io"
)

type Client struct {
	proj  string
	store *firestore.Client
	blob  *storage.Client
}

func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	strg, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		proj:  projectID,
		store: client,
		blob:  strg,
	}, nil
}

func (c *Client) Collection(ctx context.Context, cat driver.Categorizer) *firestore.CollectionRef {
	return c.store.Collection(cat.GetCategory())
}

func (c *Client) Collections(ctx context.Context) *firestore.CollectionIterator {
	return c.store.Collections(ctx)
}

func (c *Client) Documents(ctx context.Context, cat driver.Categorizer) *firestore.DocumentIterator {
	return c.Collection(ctx, cat).Documents(ctx)
}

func (c *Client) DocSnapshot(ctx context.Context, group driver.Grouping) (*firestore.DocumentSnapshot, error) {
	return c.Document(ctx, group).Get(ctx)
}

func (c *Client) Document(ctx context.Context, group driver.Grouping) *firestore.DocumentRef {
	return c.store.Collection(group.GetCategory()).Doc(group.GetIdentifier())
}

func (c *Client) MarshalDocTo(ctx context.Context, group driver.Grouping, obj interface{}) error {
	snap, err := c.Document(ctx, group).Get(ctx)
	if err != nil {
		return err
	}
	return snap.DataTo(obj)
}

func (c *Client) DocDataAt(ctx context.Context, group driver.Grouping, key string) (interface{}, error) {
	snap, err := c.Document(ctx, group).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.DataAt(key)
}

func (c *Client) DocData(ctx context.Context, group driver.Grouping) (map[string]interface{}, error) {
	snap, err := c.Document(ctx, group).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.Data(), nil
}

func (c *Client) UpdateDocField(ctx context.Context, group driver.Grouping, key string, value string) error {
	_, err := c.Document(ctx, group).Update(ctx, []firestore.Update{
		{
			Path:  key,
			Value: value,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateDoc(ctx context.Context, group driver.Grouping, data map[string]interface{}) error {
	_, err := c.Document(ctx, group).Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteDoc(ctx context.Context, group driver.Grouping) error {
	_, err := c.Document(ctx, group).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SetDocData(ctx context.Context, group driver.Grouping, data map[string]interface{}, merge bool) error {
	if merge {
		_, err := c.Document(ctx, group).Set(ctx, data, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := c.Document(ctx, group).Set(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) Object(ctx context.Context, group driver.Grouping) *storage.ObjectHandle {
	return c.blob.Bucket(group.GetCategory()).Object(group.GetIdentifier())
}

func (c *Client) CopyFromObject(ctx context.Context, from driver.Grouping, to driver.Grouping) *storage.Copier {
	return c.blob.Bucket(to.GetCategory()).Object(to.GetIdentifier()).CopierFrom(c.Object(ctx, from))
}

func (c *Client) DeleteObject(ctx context.Context, group driver.Grouping) error {
	return c.Object(ctx, group).Delete(ctx)
}

func (c *Client) UpdateObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error) {
	return c.Object(ctx, metagroup).Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: metagroup.GetMeta(),
	})
}

func (c *Client) ObjectAttributes(ctx context.Context, metagroup driver.MetaGrouping) (*storage.ObjectAttrs, error) {
	return c.Object(ctx, metagroup).Attrs(ctx)
}

func (c *Client) GetObjectMetadata(ctx context.Context, metagroup driver.MetaGrouping) (map[string]string, error) {
	attrs, err := c.Object(ctx, metagroup).Attrs(ctx)
	if err != nil {
		return nil, err
	}
	return attrs.Metadata, nil
}

func (c *Client) Bucket(ctx context.Context, cat driver.Categorizer) *storage.BucketHandle {
	return c.blob.Bucket(cat.GetCategory())
}

func (c *Client) CreateBucket(ctx context.Context, cat driver.Categorizer) error {
	return c.blob.Bucket(cat.GetCategory()).Create(ctx, c.proj, nil)
}

func (c *Client) ObjectsBucketName(ctx context.Context, grp driver.Grouping) string {
	return c.Object(ctx, grp).BucketName()
}

func (c *Client) ObjectWriter(ctx context.Context, grp driver.Grouping) *storage.Writer {
	return c.Object(ctx, grp).NewWriter(ctx)
}

func (c *Client) ObjectReader(ctx context.Context, grp driver.Grouping) (*storage.Reader, error) {
	return c.Object(ctx, grp).NewReader(ctx)
}

func (c *Client) CopyObjectTo(ctx context.Context, dst io.Writer, grp driver.Grouping) error {
	r, err := c.ObjectReader(ctx, grp)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CopyToObjectFrom(ctx context.Context, from io.Reader, grp driver.Grouping) error {
	_, err := io.Copy(c.ObjectWriter(ctx, grp), from)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteBucket(ctx context.Context, cat driver.Categorizer) error {
	return c.Bucket(ctx, cat).Delete(ctx)
}

func (c *Client) UpdateBucket(ctx context.Context, cat driver.Categorizer, attr storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	return c.Bucket(ctx, cat).Update(ctx, attr)
}

func (c *Client) BucketObject(ctx context.Context, cat driver.Categorizer) *storage.ObjectIterator {
	return c.Bucket(ctx, cat).Objects(ctx, nil)
}

func (c *Client) Buckets(ctx context.Context) *storage.BucketIterator {
	return c.blob.Buckets(ctx, c.proj)
}

func (c *Client) HandleBucket(ctx context.Context, cat driver.Categorizer, fn functions.BucketHandlerFunc) error {
	return fn(c.Bucket(ctx, cat))
}

func (c *Client) HandleObject(ctx context.Context, group driver.Grouping, fn functions.ObjectHandlerFunc) error {
	return fn(c.Object(ctx, group))
}

func (c *Client) HandleCollection(ctx context.Context, cat driver.Categorizer, fn functions.CollectionHandlerFunc) error {
	return fn(c.Collection(ctx, cat))
}

func (c *Client) HandleDocument(ctx context.Context, group driver.Grouping, fn functions.DocumentHandlerFunc) error {
	return fn(c.Document(ctx, group))
}
