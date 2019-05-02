# publish
--
    import "github.com/autom8ter/fire/publish"


## Usage

#### type Client

```go
type Client struct {
}
```


#### func  NewClient

```go
func NewClient(ctx context.Context, project string, opts ...option.ClientOption) (*Client, error)
```

#### func (*Client) GetTopic

```go
func (c *Client) GetTopic(ctx context.Context, cat driver.Categorizer) (*pubsub.Topic, error)
```

#### func (*Client) Publish

```go
func (c *Client) Publish(ctx context.Context, message driver.Message) (string, error)
```
