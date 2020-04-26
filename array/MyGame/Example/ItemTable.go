// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

// ItemTableT native go object
type ItemTableT struct {
	Bool  bool
	U64   uint64
	Color Color
	I8    int8
	F32   float32
	Ubyte byte
}

// ItemTableT object pack function
func (t *ItemTableT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}

	// pack process all field

	ItemTableStart(builder)
	ItemTableAddBool(builder, t.Bool)
	ItemTableAddU64(builder, t.U64)
	ItemTableAddColor(builder, t.Color)
	ItemTableAddI8(builder, t.I8)
	ItemTableAddF32(builder, t.F32)
	ItemTableAddUbyte(builder, t.Ubyte)
	return ItemTableEnd(builder)
}

// ItemTableT object unpack function
func (rcv *ItemTable) UnPackTo(t *ItemTableT) {
	t.Bool = rcv.Bool()
	t.U64 = rcv.U64()
	t.Color = rcv.Color()
	t.I8 = rcv.I8()
	t.F32 = rcv.F32()
	t.Ubyte = rcv.Ubyte()
}

func (rcv *ItemTable) UnPack() *ItemTableT {
	if rcv == nil {
		return nil
	}
	t := &ItemTableT{}
	rcv.UnPackTo(t)
	return t
}

type ItemTable struct {
	_tab flatbuffers.Table
}

// GetRootAsItemTable shortcut to access root table
func GetRootAsItemTable(buf []byte, offset flatbuffers.UOffsetT) *ItemTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ItemTable{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsItemTable shortcut to access table in vector of  unions
func GetTableVectorAsItemTable(table *flatbuffers.Table) *ItemTable {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &ItemTable{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsItemTable shortcut to access table in single union field
func GetTableAsItemTable(table *flatbuffers.Table) *ItemTable {
	x := &ItemTable{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *ItemTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ItemTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ItemTable) Bool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ItemTable) MutateBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *ItemTable) U64() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ItemTable) MutateU64(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *ItemTable) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ItemTable) MutateColor(n Color) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func (rcv *ItemTable) I8() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ItemTable) MutateI8(n int8) bool {
	return rcv._tab.MutateInt8Slot(10, n)
}

func (rcv *ItemTable) F32() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ItemTable) MutateF32(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func (rcv *ItemTable) Ubyte() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ItemTable) MutateUbyte(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
}

func ItemTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}

func ItemTableAddBool(builder *flatbuffers.Builder, bool bool) {
	builder.PrependBoolSlot(0, bool, false)
}

func ItemTableAddU64(builder *flatbuffers.Builder, u64 uint64) {
	builder.PrependUint64Slot(1, u64, 0)
}

func ItemTableAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(2, int8(color), 0)
}

func ItemTableAddI8(builder *flatbuffers.Builder, i8 int8) {
	builder.PrependInt8Slot(3, i8, 0)
}

func ItemTableAddF32(builder *flatbuffers.Builder, f32 float32) {
	builder.PrependFloat32Slot(4, f32, 0.0)
}

func ItemTableAddUbyte(builder *flatbuffers.Builder, ubyte byte) {
	builder.PrependByteSlot(5, ubyte, 0)
}

func ItemTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
