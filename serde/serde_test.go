package serde

import (
	"testing"

	pb "github.com/imhshekhar47/ops-spec/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var mockAgent = pb.Agent{
	Meta: &pb.Metadata{
		Timestamp: timestamppb.Now(),
		Version:   "0.0.1",
	},
	Uuid:    "mock",
	Address: "0.0.0.0:8085",
}

func TestBinarySerde(t *testing.T) {
	serde := &BinarySerde{}
	bytes, err := serde.Serialize(&mockAgent)

	if err != nil {
		t.Logf("serialization failed")
	}

	t.Logf("serialized to \n%s", string(bytes))

	copy := &pb.Agent{}
	err = serde.Deserialize(bytes, copy)

	if err != nil {
		t.Logf("deserialization failed")
	}

	if mockAgent.Meta.Timestamp.AsTime() != copy.Meta.Timestamp.AsTime() {
		t.Logf("deserialization failed, mismatch timestamp expected %v but found %v", mockAgent.Meta.Timestamp, copy.Meta.Timestamp)
	}
}

func TestJsonSerde(t *testing.T) {
	serde := &JsonSerde{}
	bytes, err := serde.Serialize(&mockAgent)

	if err != nil {
		t.Logf("serialization failed")
	}

	t.Logf("serialized to \n%s", string(bytes))

	copy := &pb.Agent{}
	err = serde.Deserialize(bytes, copy)

	if err != nil {
		t.Logf("deserialization failed")
	}

	if mockAgent.Meta.Timestamp.AsTime() != copy.Meta.Timestamp.AsTime() {
		t.Logf("deserialization failed, mismatch timestamp expected %v but found %v", mockAgent.Meta.Timestamp, copy.Meta.Timestamp)
	}
}
