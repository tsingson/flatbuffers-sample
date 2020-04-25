// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	flatbuffers "github.com/google/flatbuffers/go"
	MyGame__Example "github.com/tsingson/flatbuffers-sample/fixed-length-array/MyGame/Example"
)

// ArrayStructT native go object
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

// fixed array A0
func (rcv *ArrayStruct) A0() [1]*MyGame__Example.NestedStructT {
	result := make([]*MyGame__Example.NestedStructT, 1)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(0))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 1; j++ {
			result[j] = rcv._tab.GetInt(a + flatbuffers.UOffsetT(j*32))
		}
	}
	return result
}

// fixed array A1
func (rcv *ArrayStruct) A1() [2]*MyGame__Example.NestedStructT {
	result := make([]*MyGame__Example.NestedStructT, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt(a + flatbuffers.UOffsetT(j*32))
		}
	}
	return result
}

// fixed array A2
func (rcv *ArrayStruct) A2() [3]*MyGame__Example.NestedStructT {
	result := make([]*MyGame__Example.NestedStructT, 3)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(96))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 3; j++ {
			result[j] = rcv._tab.GetInt(a + flatbuffers.UOffsetT(j*32))
		}
	}
	return result
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

// fixed array B
func (rcv *ArrayStruct) B() [15]int32 {
	result := make([]int32, 15)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(228))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 15; j++ {
			result[j] = rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
		}
	}
	return result
}

func (rcv *ArrayStruct) C() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(288))
}

// fixed array D
func (rcv *ArrayStruct) D() [2]*MyGame__Example.NestedStructT {
	result := make([]*MyGame__Example.NestedStructT, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(296))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt(a + flatbuffers.UOffsetT(j*32))
		}
	}
	return result
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

// fixed array F
func (rcv *ArrayStruct) F() [2]int64 {
	result := make([]int64, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(400))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
		}
	}
	return result
}

// fixed array F1
func (rcv *ArrayStruct) F1() [2]*MyGame__Example.NestedStructT {
	result := make([]*MyGame__Example.NestedStructT, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(416))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt(a + flatbuffers.UOffsetT(j*32))
		}
	}
	return result
}

func CreateArrayStruct(builder *flatbuffers.Builder,
	a0 [1]*MyGame__Example.NestedStructT,
	a1 [2]*MyGame__Example.NestedStructT,
	a2 [3]*MyGame__Example.NestedStructT,
	a3_a [2]int32, a3_b MyGame__Example.TestEnum,
	a3_c [2]MyGame__Example.TestEnum,
	a3_c1 [1]MyGame__Example.TestEnum,
	a3_d [2]int64, a float32,
	b [15]int32, c int8,
	d [2]*MyGame__Example.NestedStructT,
	d1_a [2]int32, d1_b MyGame__Example.TestEnum,
	d1_c [2]MyGame__Example.TestEnum,
	d1_c1 [1]MyGame__Example.TestEnum,
	d1_d [2]int64, e int32,
	f [2]int64,
	f1 [2]*MyGame__Example.NestedStructT) flatbuffers.UOffsetT {
	builder.Prep(8, 480)
	builder.PrependUOffsetT(f1)
	builder.PrependUOffsetT(f)
	builder.Pad(4)
	builder.PrependInt32(e)
	builder.Prep(8, 32)
	builder.PrependUOffsetT(d1_d)
	builder.Pad(4)
	builder.PrependUOffsetT(int(d1_c1))
	builder.PrependUOffsetT(int(d1_c))
	builder.PrependInt8(int8(d1_b))
	builder.PrependUOffsetT(d1_a)
	builder.PrependUOffsetT(d)
	builder.Pad(7)
	builder.PrependInt8(c)
	builder.PrependUOffsetT(b)
	builder.PrependFloat32(a)
	builder.Prep(8, 32)
	builder.PrependUOffsetT(a3_d)
	builder.Pad(4)
	builder.PrependUOffsetT(int(a3_c1))
	builder.PrependUOffsetT(int(a3_c))
	builder.PrependInt8(int8(a3_b))
	builder.PrependUOffsetT(a3_a)
	builder.PrependUOffsetT(a2)
	builder.PrependUOffsetT(a1)
	builder.PrependUOffsetT(a0)
	return builder.Offset()
}