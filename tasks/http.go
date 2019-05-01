package tasks

import (
	"cloud.google.com/go/cloudtasks/apiv2beta3"
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/autom8ter/fire/api"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/api/option"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

type Client struct {
	Tasks     *cloudtasks.Client
	Publisher *pubsub.Client
	Project   string
}

func NewClient(ctx context.Context, project string, opts ...option.ClientOption) (*Client, error) {
	tasks, err := cloudtasks.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	pub, err := pubsub.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Tasks:     tasks,
		Publisher: pub,
		Project:   project,
	}, nil
}

func (c *Client) CreateHTTPTask(ctx context.Context, h api.JSONTask) (*taskspb.Task, error) {
	var meth taskspb.HttpMethod
	switch h.Method() {
	case "POST":
		meth = taskspb.HttpMethod_POST
	case "DELETE":
		meth = taskspb.HttpMethod_DELETE
	case "PATCH":
		meth = taskspb.HttpMethod_PATCH
	case "PUT":
		meth = taskspb.HttpMethod_PUT
	default:
		meth = taskspb.HttpMethod_GET
	}
	req := &taskspb.CreateTaskRequest{
		Parent: c.QueuePath(h.GetQueLocation(), h.GetQueID()),
		Task: &taskspb.Task{
			Name: h.Identifier(),
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2beta3#HttpRequest
			PayloadType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					Url:        h.URL(),
					HttpMethod: meth,
					Headers:    h.Headers(),
					Body:       []byte(h.JSONString()),
				},
			},
			ScheduleTime: &timestamp.Timestamp{
				Seconds: h.ExecuteAtUnix(),
			},
		},
	}
	createdTask, err := c.Tasks.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (c *Client) QueuePath(location, que string) string {
	return fmt.Sprintf("projects/%s/locations/%s/queues/%s", c.Project, location, que)
}
