// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type SpaceShipT struct {
	Size  *Vec3T
	Power int32
	Name  string
}

// SpaceShipT object pack function
func (t *SpaceShipT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if len(t.Name) > 0 {
		nameOffset = builder.CreateString(t.Name)
	}

	// pack process all field

	SpaceShipStart(builder)
	sizeOffset := t.Size.Pack(builder)
	SpaceShipAddSize(builder, sizeOffset)
	SpaceShipAddPower(builder, t.Power)
	SpaceShipAddName(builder, nameOffset)
	return SpaceShipEnd(builder)
}

// SpaceShipT object unpack function
func (rcv *SpaceShip) UnPackTo(t *SpaceShipT) {
	t.Size = rcv.Size(nil).UnPack()
	t.Power = rcv.Power()
	t.Name = string(rcv.Name())
}

func (rcv *SpaceShip) UnPack() *SpaceShipT {
	if rcv == nil {
		return nil
	}
	t := &SpaceShipT{}
	rcv.UnPackTo(t)
	return t
}

type SpaceShip struct {
	_tab flatbuffers.Table
}

// GetRootAsSpaceShip shortcut to access root table
func GetRootAsSpaceShip(buf []byte, offset flatbuffers.UOffsetT) *SpaceShip {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SpaceShip{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsSpaceShip shortcut to access table in vector of  unions
func GetTableVectorAsSpaceShip(table *flatbuffers.Table) *SpaceShip {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &SpaceShip{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsSpaceShip shortcut to access table in single union field
func GetTableAsSpaceShip(table *flatbuffers.Table) *SpaceShip {
	x := &SpaceShip{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *SpaceShip) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SpaceShip) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SpaceShip) Size(obj *Vec3) *Vec3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SpaceShip) Power() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SpaceShip) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SpaceShipStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}

func SpaceShipAddSize(builder *flatbuffers.Builder, Size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(Size), 0)
}

func SpaceShipAddPower(builder *flatbuffers.Builder, Power int32) {
	builder.PrependInt32Slot(1, Power, 0)
}

func SpaceShipAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Name), 0)
}

func SpaceShipEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
