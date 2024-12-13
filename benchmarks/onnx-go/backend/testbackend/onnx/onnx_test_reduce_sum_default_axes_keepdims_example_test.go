package onnxtest

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/owulveryck/onnx-go/internal/onnx/ir"
)

func TestNewTestReduceSumDefaultAxesKeepdimsExample(t *testing.T) {
	mytest := NewTestReduceSumDefaultAxesKeepdimsExample()
	var model ir.ModelProto
	err := proto.Unmarshal(mytest.ModelB, &model)
	if err != nil {
		t.Fatal(err)
	}
	if model.Graph == nil {
		t.Fatal("graph is nil")
	}
	if len(model.Graph.Input) != len(mytest.Input) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Input), len(mytest.Input))
	}
	if len(model.Graph.Output) != len(mytest.ExpectedOutput) {
		t.Fatalf("invalid test: model has %v input, but test only provide %v", len(model.Graph.Output), len(mytest.ExpectedOutput))
	}
}
