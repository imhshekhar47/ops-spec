package serde

import (
	"encoding/json"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

type Serde interface {
	Serialize(proto.Message) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

type BinarySerde struct {
}

func (s *BinarySerde) Serialize(message proto.Message) ([]byte, error) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s *BinarySerde) Deserialize(bytes []byte, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarhsal to proto.Message, message is %T", v)
	}

	return proto.Unmarshal(bytes, p)
}

type JsonSerde struct {
}

func (s *JsonSerde) Serialize(message proto.Message) ([]byte, error) {
	bytes, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (s *JsonSerde) Deserialize(bytes []byte, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarhsal to proto.Message, message is %T", v)
	}

	return json.Unmarshal(bytes, p)
}
