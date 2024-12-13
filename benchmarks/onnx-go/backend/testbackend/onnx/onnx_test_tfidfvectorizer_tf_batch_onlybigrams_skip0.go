package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("TfIdfVectorizer", "TestTfidfvectorizerTfBatchOnlybigramsSkip0", NewTestTfidfvectorizerTfBatchOnlybigramsSkip0)
}

// NewTestTfidfvectorizerTfBatchOnlybigramsSkip0 version: 3.
func NewTestTfidfvectorizerTfBatchOnlybigramsSkip0() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "TfIdfVectorizer",
		Title:  "TestTfidfvectorizerTfBatchOnlybigramsSkip0",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xaa, 0x2, 0xa, 0xcc, 0x1, 0xa, 0x1, 0x58, 0x12, 0x1, 0x59, 0x22, 0xf, 0x54, 0x66, 0x49, 0x64, 0x66, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x2a, 0x16, 0xa, 0xf, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x2, 0xa0, 0x1, 0x2, 0x2a, 0x15, 0xa, 0xe, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x2a, 0x16, 0xa, 0xf, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x2, 0xa0, 0x1, 0x2, 0x2a, 0xd, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x2, 0x54, 0x46, 0xa0, 0x1, 0x3, 0x2a, 0x15, 0xa, 0xc, 0x6e, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x40, 0x0, 0x40, 0x4, 0xa0, 0x1, 0x7, 0x2a, 0x20, 0xa, 0xd, 0x6e, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x40, 0x0, 0x40, 0x1, 0x40, 0x2, 0x40, 0x3, 0x40, 0x4, 0x40, 0x5, 0x40, 0x6, 0xa0, 0x1, 0x7, 0x2a, 0x24, 0xa, 0xb, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x73, 0x40, 0x2, 0x40, 0x3, 0x40, 0x5, 0x40, 0x4, 0x40, 0x5, 0x40, 0x6, 0x40, 0x7, 0x40, 0x8, 0x40, 0x6, 0x40, 0x7, 0xa0, 0x1, 0x7, 0x12, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x66, 0x69, 0x64, 0x66, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x74, 0x66, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x62, 0x69, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x30, 0x5a, 0x13, 0xa, 0x1, 0x58, 0x12, 0xe, 0xa, 0xc, 0x8, 0x6, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x6, 0x62, 0x13, 0xa, 0x1, 0x59, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x7, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "TfIdfVectorizer",
		     Attributes: ([]*ir.AttributeProto) (len=7 cap=8) {
		    (*ir.AttributeProto)(0xc000133a00)(name:"max_gram_length" type:INT i:2 ),
		    (*ir.AttributeProto)(0xc000133b00)(name:"max_skip_count" type:INT ),
		    (*ir.AttributeProto)(0xc000133c00)(name:"min_gram_length" type:INT i:2 ),
		    (*ir.AttributeProto)(0xc000133d00)(name:"mode" type:STRING s:"TF" ),
		    (*ir.AttributeProto)(0xc000133e00)(name:"ngram_counts" type:INTS ints:0 ints:4 ),
		    (*ir.AttributeProto)(0xc000133f00)(name:"ngram_indexes" type:INTS ints:0 ints:1 ints:2 ints:3 ints:4 ints:5 ints:6 ),
		    (*ir.AttributeProto)(0xc0001fc000)(name:"pool_int64s" type:INTS ints:2 ints:3 ints:5 ints:4 ints:5 ints:6 ints:7 ints:8 ints:6 ints:7 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 6),
				tensor.WithBacking([]int32{1, 1, 3, 3, 3, 7, 8, 6, 7, 5, 6, 8}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 7),
				tensor.WithBacking([]float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}),
			),
		},
	}
}