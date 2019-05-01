package config

import (
	"github.com/autom8ter/fire/util"
	"google.golang.org/api/option"
)

type Config struct {
	ProjectID     string                `validate:"required"`
	TasksLocation string                `validate:"required"`
	TasksQueID    string                `validate:"required"`
	Opts          []option.ClientOption `validate:"required"`
}

func NewConfig(projectID string, tasksLocation string, tasksQueID string, opts []option.ClientOption) *Config {
	return &Config{ProjectID: projectID, TasksLocation: tasksLocation, TasksQueID: tasksQueID, Opts: opts}
}

func (c *Config) Validate() error {
	return util.Util.Validate(c)
}
