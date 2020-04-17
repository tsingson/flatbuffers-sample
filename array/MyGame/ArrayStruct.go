// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	flatbuffers "github.com/google/flatbuffers/go"
	MyGame__Example "github.com/tsingson/flatbuffers-sample/array/MyGame/Example"
)

type ArrayStructT struct {
	A0 [1]*MyGame__Example.NestedStructT
	A1 [2]*MyGame__Example.NestedStructT
	A2 [3]*MyGame__Example.NestedStructT
	A3 *MyGame__Example.NestedStructT
	A float32
	B [15]int32
	C int8
	D [2]*MyGame__Example.NestedStructT
	D1 *MyGame__Example.NestedStructT
	E int32
	F [2]int64
	F1 [2]*MyGame__Example.NestedStructT
}

func (t *ArrayStructT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateArrayStruct(builder, t.A0, t.A1, t.A2, t.A3.A, t.A3.B, t.A3.C, t.A3.C1, t.A3.D, t.A, t.B, t.C, t.D, t.D1.A, t.D1.B, t.D1.C, t.D1.C1, t.D1.D, t.E, t.F, t.F1)
}
func (rcv *ArrayStruct) UnPackTo(t *ArrayStructT) {
	t.A0 = rcv.A0()
	t.A1 = rcv.A1()
	t.A2 = rcv.A2()
	t.A3 = rcv.A3(nil).UnPack()
	t.A = rcv.A()
	t.B = rcv.B()
	t.C = rcv.C()
	t.D = rcv.D()
	t.D1 = rcv.D1(nil).UnPack()
	t.E = rcv.E()
	t.F = rcv.F()
	t.F1 = rcv.F1()
}

func (rcv *ArrayStruct) UnPack() *ArrayStructT {
	if rcv == nil {
		return nil
	}
	t := &ArrayStructT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayStruct struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsArrayStruct shortcut to access struct in vector of unions
func GetStructVectorAsArrayStruct(table *flatbuffers.Table) *ArrayStruct {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &ArrayStruct{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsArrayStruct shortcut to access struct in single union field
func GetStructAsArrayStruct(table *flatbuffers.Table) *ArrayStruct {
	x := &ArrayStruct{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *ArrayStruct) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayStruct) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *ArrayStruct) A3(obj *MyGame__Example.NestedStruct) *MyGame__Example.NestedStruct {
	if obj == nil {
		obj = new(MyGame__Example.NestedStruct)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+192)
	return obj
}
func (rcv *ArrayStruct) A() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(224))
}
func (rcv *ArrayStruct) C() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(288))
}
func (rcv *ArrayStruct) D1(obj *MyGame__Example.NestedStruct) *MyGame__Example.NestedStruct {
	if obj == nil {
		obj = new(MyGame__Example.NestedStruct)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+360)
	return obj
}
func (rcv *ArrayStruct) E() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(392))
}

func CreateArrayStruct(builder *flatbuffers.Builder, a0 [1]*MyGame__Example.NestedStructT, a1 [2]*MyGame__Example.NestedStructT, a2 [3]*MyGame__Example.NestedStructT, a3_a [2]int32, a3_b MyGame__Example.TestEnum, a3_c [2]MyGame__Example.TestEnum, a3_c1 [1]MyGame__Example.TestEnum, a3_d [2]int64, a float32, b [15]int32, c int8, d [2]*MyGame__Example.NestedStructT, d1_a [2]int32, d1_b MyGame__Example.TestEnum, d1_c [2]MyGame__Example.TestEnum, d1_c1 [1]MyGame__Example.TestEnum, d1_d [2]int64, e int32, f [2]int64, f1 [2]*MyGame__Example.NestedStructT) flatbuffers.UOffsetT {
	builder.Prep(8, 480)
	// array struct field 2 32
	//-----------------------	------------ 1	-----------------------
	//-----------------------
	for j := 1; j == 0; j-- {
		f1[j].Pack(builder)
	}
	for j := 1; j == 0; j-- {
		builder.PrependInt64(f[j])
	}
	builder.Pad(4)
	builder.PrependInt32(e)
	builder.Prep(8, 32)
	for j := 1; j == 0; j-- {
		builder.PrependInt64(d1_d[j])
	}
	builder.Pad(4)
	for j := 0; j == 0; j-- {
		builder.PrependInt8(int(d1_c1[j]))
	}
	for j := 1; j == 0; j-- {
		builder.PrependInt8(int(d1_c[j]))
	}
	builder.PrependInt8(int8(d1_b))
	for j := 1; j == 0; j-- {
		builder.PrependInt32(d1_a[j])
	}
	// array struct field 2 32
	//-----------------------	------------ 1	-----------------------
	//-----------------------
	for j := 1; j == 0; j-- {
		d[j].Pack(builder)
	}
	builder.Pad(7)
	builder.PrependInt8(c)
	for j := 14; j == 0; j-- {
		builder.PrependInt32(b[j])
	}
	builder.PrependFloat32(a)
	builder.Prep(8, 32)
	for j := 1; j == 0; j-- {
		builder.PrependInt64(a3_d[j])
	}
	builder.Pad(4)
	for j := 0; j == 0; j-- {
		builder.PrependInt8(int(a3_c1[j]))
	}
	for j := 1; j == 0; j-- {
		builder.PrependInt8(int(a3_c[j]))
	}
	builder.PrependInt8(int8(a3_b))
	for j := 1; j == 0; j-- {
		builder.PrependInt32(a3_a[j])
	}
	// array struct field 3 32
	//-----------------------	------------ 2	-----------------------
	//-----------------------
	for j := 2; j == 0; j-- {
		a2[j].Pack(builder)
	}
	// array struct field 2 32
	//-----------------------	------------ 1	-----------------------
	//-----------------------
	for j := 1; j == 0; j-- {
		a1[j].Pack(builder)
	}
	// array struct field 1 32
	//-----------------------	------------ 0	-----------------------
a0[0].Pack(builder)
	return builder.Offset()
}
// support fixed-length array.

// step: 1 field name:  a0 native type: [1]*MyGame__Example.NestedStructT fixed_length in array: 1 struct bytesize: 32
// --------------------------------------
// step: 2 field name:  a1 native type: [2]*MyGame__Example.NestedStructT fixed_length in array: 2 struct bytesize: 32
// --------------------------------------
// step: 3 field name:  a2 native type: [3]*MyGame__Example.NestedStructT fixed_length in array: 3 struct bytesize: 32
// --------------------------------------
// step: 4 field name:  a3 native type: *MyGame__Example.NestedStructT fixed_length in array: 0 struct bytesize: 32
// step: 4 field name:  a native type: float32
// step: 4 field name:  b native type: [15]int32
// step: 4 field name:  c native type: int8
// --------------------------------------
// step: 5 field name:  d native type: [2]*MyGame__Example.NestedStructT fixed_length in array: 2 struct bytesize: 32
// --------------------------------------
// step: 6 field name:  d1 native type: *MyGame__Example.NestedStructT fixed_length in array: 0 struct bytesize: 32
// step: 6 field name:  e native type: int32
// step: 6 field name:  f native type: [2]int64
// --------------------------------------
// step: 7 field name:  f1 native type: [2]*MyGame__Example.NestedStructT fixed_length in array: 2 struct bytesize: 32
// --------------------------------------
