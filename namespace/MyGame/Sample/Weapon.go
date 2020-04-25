// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Sample

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type WeaponT struct {
	Name   string
	Damage int16
}

// WeaponT object pack function
func (t *WeaponT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if len(t.Name) > 0 {
		nameOffset = builder.CreateString(t.Name)
	}

	// pack process all field

	WeaponStart(builder)
	WeaponAddName(builder, nameOffset)
	WeaponAddDamage(builder, t.Damage)
	return WeaponEnd(builder)
}

// WeaponT object unpack function
func (rcv *Weapon) UnPackTo(t *WeaponT) {
	t.Name = string(rcv.Name())
	t.Damage = rcv.Damage()
}

func (rcv *Weapon) UnPack() *WeaponT {
	if rcv == nil {
		return nil
	}
	t := &WeaponT{}
	rcv.UnPackTo(t)
	return t
}

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

func (rcv *Weapon) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Weapon) Damage() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func WeaponStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}

func WeaponAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}

func WeaponAddDamage(builder *flatbuffers.Builder, damage int16) {
	builder.PrependInt16Slot(1, damage, 0)
}

func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
