package gojopb

import (
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Options for marshal and unmarshal operations.
// A copy will be made before operation as the content of option will be
// modified (nil Resolver will be set) during operation.
var (
	DefaultMarshalOptions   protojson.MarshalOptions
	DefaultUnmarshalOptions proto.UnmarshalOptions
)

// ErrNilMessageReference indicate given message reference is nil.
var ErrNilMessageReference = errors.New("nil message reference")

// ProtobufMessage wrap protocol buffers message make it able to marshal and
// unmarshal with encoding/json package.
type ProtobufMessage struct {
	Message proto.Message
}

func (m *ProtobufMessage) MarshalJSON() ([]byte, error) {
	if m.Message == nil {
		return []byte("null"), nil
	}
	o := DefaultMarshalOptions
	return o.Marshal(m.Message)
}

func (m *ProtobufMessage) UnmarshalJSON(buf []byte) error {
	if m.Message == nil {
		return ErrNilMessageReference
	}
	o := DefaultUnmarshalOptions
	return o.Unmarshal(buf, m.Message)
}
