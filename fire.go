package fire

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Document interface {
	DocCategory() string
	DocName() string
}

type Client struct {
	Fire *firestore.Client
}

func NewClient(ctx context.Context, projectID string) (*Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Client{
		Fire: client,
	}, nil
}

func (c *Client) Set(ctx context.Context, d Document, data map[string]interface{}, merge bool) error {
	if merge {
		_, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Set(ctx, data, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Set(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) Snapshot(ctx context.Context, d Document) (*firestore.DocumentSnapshot, error) {
	return c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Get(ctx)
}

func (c *Client) Ref(ctx context.Context, d Document) *firestore.DocumentRef {
	return c.Fire.Collection(d.DocCategory()).Doc(d.DocName())
}

func (c *Client) MarshalTo(ctx context.Context, d Document, obj interface{}) error {
	snap, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Get(ctx)
	if err != nil {
		return err
	}
	return snap.DataTo(obj)
}

func (c *Client) DataAt(ctx context.Context, d Document, key string) (interface{}, error) {
	snap, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.DataAt(key)
}

func (c *Client) UpdateField(ctx context.Context, d Document, key string, value string) error {
	_, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Update(ctx, []firestore.Update{
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

func (c *Client) Data(ctx context.Context, d Document) (map[string]interface{}, error) {
	snap, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.Data(), nil
}

func (c *Client) Create(ctx context.Context, d Document, data map[string]interface{}) error {
	_, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Delete(ctx context.Context, d Document) error {
	_, err := c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
