package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Gemm", "TestGemmNobroadcast", NewTestGemmNobroadcast)
}

// NewTestGemmNobroadcast version: 3.
func NewTestGemmNobroadcast() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Gemm",
		Title:  "TestGemmNobroadcast",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa0, 0x1, 0xa, 0x33, 0xa, 0x1, 0x61, 0xa, 0x1, 0x62, 0xa, 0x1, 0x63, 0x12, 0x1, 0x79, 0x22, 0x4, 0x47, 0x65, 0x6d, 0x6d, 0x2a, 0xf, 0xa, 0x5, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x15, 0x0, 0x0, 0x0, 0x3f, 0xa0, 0x1, 0x1, 0x2a, 0xe, 0xa, 0x4, 0x62, 0x65, 0x74, 0x61, 0x15, 0x0, 0x0, 0x0, 0x3f, 0xa0, 0x1, 0x1, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x6d, 0x6d, 0x5f, 0x6e, 0x6f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5a, 0x13, 0xa, 0x1, 0x61, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x6, 0x5a, 0x13, 0xa, 0x1, 0x62, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x13, 0xa, 0x1, 0x63, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"a", "b", "c"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Gemm",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc000132a00)(name:"alpha" type:FLOAT f:0.5 ),
		    (*ir.AttributeProto)(0xc000132b00)(name:"beta" type:FLOAT f:0.5 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 6),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985}),
			),

			tensor.New(
				tensor.WithShape(6, 4),
				tensor.WithBacking([]float32{0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337, 0.45615032, 0.56843394, 0.0187898, 0.6176355, 0.6120957, 0.616934, 0.94374806, 0.6818203, 0.3595079, 0.43703195}),
			),

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{0.6976312, 0.06022547, 0.6667667, 0.67063785, 0.21038257, 0.12892629, 0.31542835, 0.36371076, 0.57019675, 0.43860152, 0.9883738, 0.10204481}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4),
				tensor.WithBacking([]float32{1.1514105, 1.3946263, 1.1715276, 1.335104, 0.85801727, 1.6312557, 1.1008396, 1.3106091, 1.1293826, 1.1850055, 1.0211351, 0.8018247}),
			),
		},
	}
}
