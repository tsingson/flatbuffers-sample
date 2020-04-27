// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

// ItemStructT native go object
type ItemStructT struct {
	Bool bool
	U64 uint64
	Color Color
	I8 int8
	F32 float32
	Ubyte byte
}

func (t *ItemStructT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateItemStruct(builder, t.Bool, t.U64, t.Color, t.I8, t.F32, t.Ubyte)
}

func (rcv *ItemStruct) UnPackTo(t *ItemStructT) {
	t.Bool = rcv.Bool()
	t.U64 = rcv.U64()
	t.Color = rcv.Color()
	t.I8 = rcv.I8()
	t.F32 = rcv.F32()
	t.Ubyte = rcv.Ubyte()
}

func (rcv *ItemStruct) UnPack() *ItemStructT {
	if rcv == nil {
		return nil
	}
	t := &ItemStructT{}
	rcv.UnPackTo(t)
	return t
}

type ItemStruct struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsItemStruct shortcut to access struct in vector of unions
func GetStructVectorAsItemStruct(table *flatbuffers.Table) *ItemStruct {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &ItemStruct{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsItemStruct shortcut to access struct in single union field
func GetStructAsItemStruct(table *flatbuffers.Table) *ItemStruct {
	x := &ItemStruct{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *ItemStruct) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ItemStruct) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *ItemStruct) Bool() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}

func (rcv *ItemStruct) U64() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}

func (rcv *ItemStruct) Color() Color {
	return Color(rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(16)))
}

func (rcv *ItemStruct) I8() int8 {
	return rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(17))
}

func (rcv *ItemStruct) F32() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(20))
}

func (rcv *ItemStruct) Ubyte() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(24))
}

func CreateItemStruct(builder *flatbuffers.Builder, 
	bool bool, 
	u64 uint64, 
	color Color, 
	i8 int8, 
	f32 float32, 
	ubyte byte) flatbuffers.UOffsetT {
	builder.Prep(8, 32)
	builder.Pad(7)
	builder.PrependFloat32(f32)
	builder.Pad(2)
	builder.PrependInt8(int8(color))
	builder.PrependUint64(u64)
	builder.Pad(7)
	return builder.Offset()
}
