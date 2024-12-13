package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Concat", "TestConcat3dAxis1", NewTestConcat3dAxis1)
}

// NewTestConcat3dAxis1 version: 3.
func NewTestConcat3dAxis1() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Concat",
		Title:  "TestConcat3dAxis1",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa0, 0x1, 0xa, 0x2d, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x6, 0x43, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x61, 0x74, 0x5f, 0x33, 0x64, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x31, 0x5a, 0x1c, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x1c, 0xa, 0x6, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1c, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"value0", "value1"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "Concat",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0003a0200)(name:"axis" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 2, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8}),
			),

			tensor.New(
				tensor.WithShape(2, 2, 2),
				tensor.WithBacking([]float32{9, 10, 11, 12, 13, 14, 15, 16}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 4, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4, 9, 10, 11, 12, 5, 6, 7, 8, 13, 14, 15, 16}),
			),
		},
	}
}