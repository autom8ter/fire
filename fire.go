//go:generate godocdown -o README.md

package fire

import (
	"context"
	"github.com/autom8ter/api/common"
	"github.com/autom8ter/fire/db"
	"github.com/autom8ter/fire/publish"
	"github.com/autom8ter/gosaas/util"
	"google.golang.org/api/option"
)

type Client struct {
	db    *db.Client      `validate:"required"`
	pub *publish.Client `validate:"required"`
}

func (c *Client) Validate() error {
	return common.Util.Validate(c)
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
	p, err := publish.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		db:    client,
		pub: p,
	}, nil
}
