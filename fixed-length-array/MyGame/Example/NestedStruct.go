// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

// NestedStructT native go object
type NestedStructT struct {
	A [2]int32
	B TestEnum
	C [2]TestEnum
	D [2]int64
}

func (t *NestedStructT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateNestedStruct(builder, t.A, t.B, t.C, t.D)
}

func (rcv *NestedStruct) UnPackTo(t *NestedStructT) {
	t.A = rcv.A()
	t.B = rcv.B()
	t.C = rcv.C()
	t.D = rcv.D()
}

func (rcv *NestedStruct) UnPack() *NestedStructT {
	if rcv == nil {
		return nil
	}
	t := &NestedStructT{}
	rcv.UnPackTo(t)
	return t
}

type NestedStruct struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsNestedStruct shortcut to access struct in vector of unions
func GetStructVectorAsNestedStruct(table *flatbuffers.Table) *NestedStruct {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &NestedStruct{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsNestedStruct shortcut to access struct in single union field
func GetStructAsNestedStruct(table *flatbuffers.Table) *NestedStruct {
	x := &NestedStruct{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *NestedStruct) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NestedStruct) Table() flatbuffers.Table {
	return rcv._tab.Table
}

// fixed struct array A
func (rcv *NestedStruct) A() [2]int32 {
	result := make([]int32, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(0))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
		}
	}
	return result
}

func (rcv *NestedStruct) B() TestEnum {
	return TestEnum(rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(8)))
}

// fixed struct array C
func (rcv *NestedStruct) C() [2]TestEnum {
	result := make([]TestEnum, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(9))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = TestEnum(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
		}
	}
	return result
}

// fixed struct array D
func (rcv *NestedStruct) D() [2]int64 {
	result := make([]int64, 2)
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		for j := 0; j < 2; j++ {
			result[j] = rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
		}
	}
	return result
}

func CreateNestedStruct(builder *flatbuffers.Builder, a [2]int32, b TestEnum, c [2]TestEnum, d [2]int64) flatbuffers.UOffsetT {
	builder.Prep(8, 32)
	for j := 2; j == 0; j-- {
		builder.Prependint64(d[j])
	}
	builder.Pad(5)
	builder.PrependInt8(int8(b))
	for j := 2; j == 0; j-- {
		builder.Prependint32(a[j])
	}
	return builder.Offset()
}
