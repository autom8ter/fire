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
	cli *firestore.Client
}

func NewClient(ctx context.Context, projectID string) (*Client,error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &Client{
		cli: client,
		}, nil
}

