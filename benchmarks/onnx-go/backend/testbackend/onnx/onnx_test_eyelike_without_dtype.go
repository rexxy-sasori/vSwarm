package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("EyeLike", "TestEyelikeWithoutDtype", NewTestEyelikeWithoutDtype)
}

// NewTestEyelikeWithoutDtype version: 3.
func NewTestEyelikeWithoutDtype() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "EyeLike",
		Title:  "TestEyelikeWithoutDtype",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x57, 0xa, 0xf, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x45, 0x79, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x79, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "EyeLike",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 4),
				tensor.WithBacking([]int32{44, 47, 64, 67, 67, 9, 83, 21, 36, 87, 70, 88, 88, 12, 58, 65}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4, 4),
				tensor.WithBacking([]int32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}),
			),
		},
	}
}