//go:generate godocdown -o README.md

package publish

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/autom8ter/api/driver"
	"github.com/autom8ter/fire/util"
	"google.golang.org/api/option"
)

type Client struct {
	pub  *pubsub.Client
	proj string
}

func NewClient(ctx context.Context, project string, opts ...option.ClientOption) (*Client, error) {
	pub, err := pubsub.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{

		pub:  pub,
		proj: project,
	}, nil
}

func (c *Client) Publish(ctx context.Context, message driver.Message) (string, error) {
	t, err := c.GetTopic(ctx, message)
	if err != nil {
		return "", err
	}
	r := t.Publish(ctx, &pubsub.Message{
		ID:         message.Identifier(),
		Data:       []byte(message.String()),
		Attributes: message.Meta(),
	})

	return r.Get(ctx)
}

func (c *Client) PublishJSON(ctx context.Context, message driver.JSONMessage) (string, error) {
	t, err := c.GetTopic(ctx, message)
	if err != nil {
		return "", err
	}
	r := t.Publish(ctx, &pubsub.Message{
		ID:         message.Identifier(),
		Data:       []byte(message.JSONString()),
		Attributes: message.Meta(),
	})

	return r.Get(ctx)
}

func (c *Client) PublishProto(ctx context.Context, message driver.ProtoMessage) (string, error) {
	t, err := c.GetTopic(ctx, message)
	if err != nil {
		return "", err
	}

	r := t.Publish(ctx, &pubsub.Message{
		ID:         message.Identifier(),
		Data:       util.Util.MarshalProto(message),
		Attributes: message.Meta(),
	})

	return r.Get(ctx)
}

func (c *Client) GetTopic(ctx context.Context, cat driver.Categorizer) (*pubsub.Topic, error) {
	t := c.pub.Topic(cat.Category())
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ok {
		t, err = c.pub.CreateTopic(ctx, cat.Category())
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
