package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/autom8ter/fire/api"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"github.com/go-redis/redis"
	"cloud.google.com/go/storage"
)

type Client struct {
	store *firestore.Client
	blob *storage.Client
}

func NewClient(ctx context.Context, addr string, password string, projectID string, opts ...option.ClientOption) (*Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	strg, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		store: client,
		blob: strg,
	}, nil
}

func (c *Client) SetDocData(ctx context.Context, group api.Grouping, data map[string]interface{}, merge bool) error {
	if merge {
		_, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Set(ctx, data, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Set(ctx, data)
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) DocSnapshot(ctx context.Context, group api.Grouping) (*firestore.DocumentSnapshot, error) {
	return c.store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
}

func (c *Client) DocRef(ctx context.Context, group api.Grouping) *firestore.DocumentRef {
	return c.store.Collection(group.Category()).Doc(group.Identifier())
}

func (c *Client) MarshalDocTo(ctx context.Context, group api.Grouping, obj interface{}) error {
	snap, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return err
	}
	return snap.DataTo(obj)
}

func (c *Client) DocDataAt(ctx context.Context, group api.Grouping, key string) (interface{}, error) {
	snap, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.DataAt(key)
}

func (c *Client) UpdateDocField(ctx context.Context, group api.Grouping, key string, value string) error {
	_, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Update(ctx, []firestore.Update{
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
	snap, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Get(ctx)
	if err != nil {
		return nil, err
	}
	return snap.Data(), nil
}

func (c *Client) CreateDoc(ctx context.Context, group api.Grouping, data map[string]interface{}) error {
	_, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteDoc(ctx context.Context, group api.Grouping) error {
	_, err := c.store.Collection(group.Category()).Doc(group.Identifier()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
