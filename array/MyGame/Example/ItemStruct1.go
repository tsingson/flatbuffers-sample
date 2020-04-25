// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ItemStruct1T struct {
	Bool bool
	Color Color
}

func (t *ItemStruct1T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateItemStruct1(builder, t.Bool, t.Color)
}
func (rcv *ItemStruct1) UnPackTo(t *ItemStruct1T) {
	t.Bool = rcv.Bool()
	t.Color = rcv.Color()
}

func (rcv *ItemStruct1) UnPack() *ItemStruct1T {
	if rcv == nil {
		return nil
	}
	t := &ItemStruct1T{}
	rcv.UnPackTo(t)
	return t
}

type ItemStruct1 struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsItemStruct1 shortcut to access struct in vector of unions
func GetStructVectorAsItemStruct1(table *flatbuffers.Table) *ItemStruct1 {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &ItemStruct1{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsItemStruct1 shortcut to access struct in single union field
func GetStructAsItemStruct1(table *flatbuffers.Table) *ItemStruct1 {
	x := &ItemStruct1{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *ItemStruct1) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ItemStruct1) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *ItemStruct1) Bool() bool {
	return rcv._tab.GetBool(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *ItemStruct1) Color() Color {
	return Color(rcv._tab.GetInt8(rcv._tab.Pos + flatbuffers.UOffsetT(1)))
}

func CreateItemStruct1(builder *flatbuffers.Builder, Bool bool, Color Color) flatbuffers.UOffsetT {
	builder.Prep(1, 2)
	builder.PrependInt8(int8(Color))
	builder.PrependBool(Bool)
	return builder.Offset()
}
