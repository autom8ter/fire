# tasks
--
    import "github.com/autom8ter/fire/tasks"


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

#### func (*Client) CreateHTTPTask

```go
func (c *Client) CreateHTTPTask(ctx context.Context, h driver.JSONTask) (*taskspb.Task, error)
```

#### func (*Client) GetTopic

```go
func (c *Client) GetTopic(ctx context.Context, cat driver.Categorizer) (*pubsub.Topic, error)
```

#### func (*Client) HandleTask

```go
func (c *Client) HandleTask(tsk *taskspb.Task, fn TaskHandleFunc) error
```

#### func (*Client) Publish

```go
func (c *Client) Publish(ctx context.Context, message driver.Message) (string, error)
```

#### func (*Client) PublishJSON

```go
func (c *Client) PublishJSON(ctx context.Context, message driver.JSONMessage) (string, error)
```

#### func (*Client) PublishProto

```go
func (c *Client) PublishProto(ctx context.Context, message driver.ProtoMessage) (string, error)
```

#### func (*Client) QueuePath

```go
func (c *Client) QueuePath(location, que string) string
```

#### type TaskHandleFunc

```go
type TaskHandleFunc func(c *Client) func(fn *taskspb.Task) error
```
