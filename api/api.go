package api

import "github.com/golang/protobuf/proto"

type Grouping interface {
	Categorizer
	Identifier
}

type Identifier interface {
	Identifier() string
}

type Categorizer interface {
	Category() string
}

type JSON interface {
	JSONString() string
}

type Queue interface {
	GetQueLocation() string
	GetQueID() string
	ExecuteAtUnix() int64
}

type JSONTask interface {
	JSON
	Identifier
	URL() string
	Method() string
	Headers() map[string]string
	Queue
}

type Metadata interface {
	Meta()map[string]string
}

type JSONMessage interface {
	Grouping
	Metadata
	JSON
}

type Message interface {
	Grouping
	Metadata
	String() string
}

type ProtoMessage interface {
	Grouping
	Metadata
	proto.Message
}