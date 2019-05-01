package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/autom8ter/fire/api"
	"google.golang.org/api/option"
)

type Client struct {
	Store *firestore.Client
}

func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Store: client,
	}, nil
}

func (c *Client) Set(ctx context.Context, group api.Grouping, data map[string]interface{}, merge bool) error {
	if merge {
		_, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Set(ctx, data, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Set(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) DocSnapshot(ctx context.Context, group api.Grouping) (*firestore.DocumentSnapshot, error) {
	return c.Store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
}

func (c *Client) DocRef(ctx context.Context, group api.Grouping) *firestore.DocumentRef {
	return c.Store.Collection(group.Category()).Doc(group.Identifier())
}

func (c *Client) MarshalDocTo(ctx context.Context, group api.Grouping, obj interface{}) error {
	snap, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return err
	}
	return snap.DataTo(obj)
}

func (c *Client) DocDataAt(ctx context.Context, group api.Grouping, key string) (interface{}, error) {
	snap, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.DataAt(key)
}

func (c *Client) UpdateDocField(ctx context.Context, group api.Grouping, key string, value string) error {
	_, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Update(ctx, []firestore.Update{
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

func (c *Client) DocData(ctx context.Context, group api.Grouping) (map[string]interface{}, error) {
	snap, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.Data(), nil
}

func (c *Client) CreateDoc(ctx context.Context, group api.Grouping, data map[string]interface{}) error {
	_, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteDoc(ctx context.Context, group api.Grouping) error {
	_, err := c.Store.Collection(group.Category()).Doc(group.Identifier()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
