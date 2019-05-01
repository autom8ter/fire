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
	Meta() map[string]string
}

type JSONMessage interface {
	MetaGrouping
	JSON
}

type Message interface {
	MetaGrouping
	String() string
}

type ProtoMessage interface {
	MetaGrouping
	proto.Message
}

type MetaGrouping interface {
	Grouping
	Metadata
}
