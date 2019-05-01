package fire

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Document interface {
	DocCategory() string
	DocName() string
	DocData() map[string]interface{}
}

type Client struct {
	Fire *firestore.Client
}

func NewClient(ctx context.Context, projectID string) (*Client,error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Client{
		Fire: client,
		}, nil
}

func (c *Client) Set(ctx context.Context, d Document, merge bool) error {
	if merge {
		_, err :=c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Set(ctx, d.DocData(), firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err :=c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Set(ctx, d.DocData())
		if err != nil {
			return err
		}
		return nil
	}
}

func (c *Client) Snapshot(ctx context.Context, d Document) (*firestore.DocumentSnapshot, error) {
	return  c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Get(ctx)
}
func (c *Client) Ref(ctx context.Context, d Document) (*firestore.DocumentRef) {
	return  c.Fire.Collection(d.DocCategory()).Doc(d.DocName())
}
func (c *Client) Create(ctx context.Context, d Document) (error) {
	_, err :=  c.Fire.Collection(d.DocCategory()).Doc(d.DocName()).Create(ctx, d.DocData())
	if err != nil {
		return err
	}
	return nil
}