package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Mod", "TestModMixedSignInt64", NewTestModMixedSignInt64)
}

// NewTestModMixedSignInt64 version: 5.
func NewTestModMixedSignInt64() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Mod",
		Title:  "TestModMixedSignInt64",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x5e, 0xa, 0xe, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x3, 0x4d, 0x6f, 0x64, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x5f, 0x6d, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x5a, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0xf, 0xa, 0x1, 0x7a, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"z"},
		     Name:      "",
		     OpType:    "Mod",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]int64{-4, 7, 5, 4, -7, 8}),
			),

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]int64{2, -3, 8, -2, 3, 5}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]int64{0, -2, 5, 0, 2, 3}),
			),
		},
	}
}