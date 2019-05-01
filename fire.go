//go:generate godocdown -o README.md

package fire

import (
	"context"
	"github.com/autom8ter/fire/db"
	"github.com/autom8ter/fire/publish"
	"github.com/autom8ter/fire/util"
	"google.golang.org/api/option"
)

type Client struct {
	DB    *db.Client      `validate:"required"`
	Tasks *publish.Client `validate:"required"`
}

func (c *Client) Validate() error {
	return util.Util.Validate(c)
}

type HandlerFunc func(c *Client) error

func (c *Client) HandleFunc(fns ...HandlerFunc) error {
	if err := util.Util.Validate(c); err != nil {
		return err
	}
	for _, f := range fns {
		err := f(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewClient(ctx context.Context, project string, opts ...option.ClientOption) (*Client, error) {
	client, err := db.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	tsks, err := publish.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		DB:    client,
		Tasks: tsks,
	}, nil
}
