// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package weapons

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type GunT struct {
	Damage int16
	Bool   bool
	Name   string
	Names  []string
}

// GunT object pack function
func (t *GunT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if len(t.Name) > 0 {
		nameOffset = builder.CreateString(t.Name)
	}
	namesOffset := flatbuffers.UOffsetT(0)
	if t.Names != nil {
		namesOffset = builder.StringsVector(t.Names...)
	}

	// pack process all field

	GunStart(builder)
	GunAddDamage(builder, t.Damage)
	GunAddBool(builder, t.Bool)
	GunAddName(builder, nameOffset)
	GunAddNames(builder, namesOffset)
	return GunEnd(builder)
}

// GunT object unpack function
func (rcv *Gun) UnPackTo(t *GunT) {
	t.Damage = rcv.Damage()
	t.Bool = rcv.Bool()
	t.Name = string(rcv.Name())
	namesLength := rcv.NamesLength()
	t.Names = make([]string, namesLength)
	for j := 0; j < namesLength; j++ {
		t.Names[j] = string(rcv.Names(j))
	}
}

func (rcv *Gun) UnPack() *GunT {
	if rcv == nil {
		return nil
	}
	t := &GunT{}
	rcv.UnPackTo(t)
	return t
}

type Gun struct {
	_tab flatbuffers.Table
}

// GetRootAsGun shortcut to access root table
func GetRootAsGun(buf []byte, offset flatbuffers.UOffsetT) *Gun {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Gun{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsGun shortcut to access table in vector of  unions
func GetTableVectorAsGun(table *flatbuffers.Table) *Gun {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Gun{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsGun shortcut to access table in single union field
func GetTableAsGun(table *flatbuffers.Table) *Gun {
	x := &Gun{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Gun) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Gun) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Gun) Damage() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Gun) MutateDamage(n int16) bool {
	return rcv._tab.MutateInt16Slot(4, n)
}

func (rcv *Gun) Bool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Gun) MutateBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Gun) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Gun) Names(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Gun) NamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func GunStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}

func GunAddDamage(builder *flatbuffers.Builder, damage int16) {
	builder.PrependInt16Slot(0, damage, 0)
}

func GunAddBool(builder *flatbuffers.Builder, bool bool) {
	builder.PrependBoolSlot(1, bool, false)
}

func GunAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}

func GunAddNames(builder *flatbuffers.Builder, names flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(names), 0)
}

func GunStartNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func GunEndNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func GunEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
