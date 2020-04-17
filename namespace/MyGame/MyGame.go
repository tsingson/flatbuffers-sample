// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MyGameT struct {
	Characters *CharacterT
}

// MyGameT object pack function
func (t *MyGameT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	CharactersOffset := t.Characters.Pack(builder)
	

	// pack process all field

	MyGameStart(builder)
	if t.Characters != nil {
		MyGameAddCharactersType(builder, t.Characters.Type)
	}
	MyGameAddCharacters(builder, CharactersOffset)
	return MyGameEnd(builder)
}

// MyGameT object unpack function
func (rcv *MyGame) UnPackTo(t *MyGameT) {
	CharactersTable := flatbuffers.Table{}
	if rcv.Characters(&CharactersTable) {
		t.Characters = rcv.CharactersType().UnPack(CharactersTable)
	}
}

func (rcv *MyGame) UnPack() *MyGameT {
	if rcv == nil {
		return nil
	}
	t := &MyGameT{}
	rcv.UnPackTo(t)
	return t
}

type MyGame struct {
	_tab flatbuffers.Table
}

// GetRootAsMyGame shortcut to access root table
func GetRootAsMyGame(buf []byte, offset flatbuffers.UOffsetT) *MyGame {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MyGame{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsMyGame shortcut to access table in vector of  unions
func GetTableVectorAsMyGame(table *flatbuffers.Table) *MyGame {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &MyGame{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsMyGame shortcut to access table in single union field
func GetTableAsMyGame(table *flatbuffers.Table) *MyGame {
	x := &MyGame{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *MyGame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MyGame) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MyGame) CharactersType() Character {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Character(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *MyGame) MutateCharactersType(n Character) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *MyGame) Characters(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MyGameStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}

func MyGameAddCharactersType(builder *flatbuffers.Builder, CharactersType Character) {
	builder.PrependByteSlot(0, byte(CharactersType), 0)
}

func MyGameAddCharacters(builder *flatbuffers.Builder, Characters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Characters), 0)
}

func MyGameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
