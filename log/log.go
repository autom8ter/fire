package log

import (
	"cloud.google.com/go/logging"
	"context"
	"google.golang.org/api/option"
)

type Logger struct {
	log *logging.Client
}

func NewLogger(proj string, opts ...option.ClientOption) (*Logger, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, proj, opts...)
	if err != nil {
		return nil, err
	}
	return &Logger{
		log: client,
	}, nil
}

func (l *Logger) Err(err error) {
	l.log.OnError(err)
}
