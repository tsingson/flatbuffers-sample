// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec3 struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsVec3 shortcut to access struct in vector of unions
func GetStructVectorAsVec3(table *flatbuffers.Table) *Vec3 {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Vec3{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsVec3 shortcut to access struct in single union field
func GetStructAsVec3(table *flatbuffers.Table) *Vec3 {
	x := &Vec3{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Vec3) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec3) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec3) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Vec3) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Vec3) Z() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}

func CreateVec3(builder *flatbuffers.Builder, x float32, y float32, z float32) flatbuffers.UOffsetT {
	builder.Prep(4, 12)
	builder.PrependFloat32(z)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
