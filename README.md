# fire
--
    import "github.com/autom8ter/fire"


## Usage

#### type Client

```go
type Client struct {
	DB    *db.Client      `validate:"required"`
	Tasks *publish.Client `validate:"required"`
}
```


#### func  NewClient

```go
func NewClient(ctx context.Context, project string, opts ...option.ClientOption) (*Client, error)
```

#### func (*Client) HandleFunc

```go
func (c *Client) HandleFunc(fns ...HandlerFunc) error
```

#### func (*Client) Validate

```go
func (c *Client) Validate() error
```

#### type HandlerFunc

```go
type HandlerFunc func(c *Client) error
```
