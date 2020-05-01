// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Weapon struct {
	_tab flatbuffers.Table
}

// GetRootAsWeapon shortcut to access root table
func GetRootAsWeapon(buf []byte, offset flatbuffers.UOffsetT) *Weapon {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Weapon{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsWeapon shortcut to access table in vector of  unions
func GetTableVectorAsWeapon(table *flatbuffers.Table) *Weapon {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Weapon{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsWeapon shortcut to access table in single union field
func GetTableAsWeapon(table *flatbuffers.Table) *Weapon {
	x := &Weapon{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Weapon) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Weapon) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Weapon) Size(obj *Vec3) *Vec3 {
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

func (rcv *Weapon) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Weapon) Power() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func WeaponStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}

func WeaponAddSize(builder *flatbuffers.Builder, size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(size), 0)
}

func WeaponAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(1, int8(color), 0)
}

func WeaponAddPower(builder *flatbuffers.Builder, power int32) {
	builder.PrependInt32Slot(2, power, 0)
}

func WeaponAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}

func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
