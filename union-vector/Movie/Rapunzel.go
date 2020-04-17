// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Movie

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RapunzelT struct {
	HairLength int32
}

func (t *RapunzelT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateRapunzel(builder, t.HairLength)
}
func (rcv *Rapunzel) UnPackTo(t *RapunzelT) {
	t.HairLength = rcv.HairLength()
}

func (rcv *Rapunzel) UnPack() *RapunzelT {
	if rcv == nil {
		return nil
	}
	t := &RapunzelT{}
	rcv.UnPackTo(t)
	return t
}

type Rapunzel struct {
	_tab flatbuffers.Struct
}

// GetStructVectorAsRapunzel shortcut to access struct in vector of unions
func GetStructVectorAsRapunzel(table *flatbuffers.Table) *Rapunzel {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Rapunzel{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetStructAsRapunzel shortcut to access struct in single union field
func GetStructAsRapunzel(table *flatbuffers.Table) *Rapunzel {
	x := &Rapunzel{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Rapunzel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Rapunzel) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Rapunzel) HairLength() int32 {
	return rcv._tab.GetInt32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Rapunzel) MutateHairLength(n int32) bool {
	return rcv._tab.MutateInt32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func CreateRapunzel(builder *flatbuffers.Builder, hairLength int32) flatbuffers.UOffsetT {
	builder.Prep(4, 4)
	builder.PrependInt32(hairLength)
	return builder.Offset()
}

// support fixed-length array.

func CreateRapunzel(builder *flatbuffers.Builder, hairLength int32) flatbuffers.UOffsetT {
	builder.Prep(4, 4)
	builder.PrependInt32(hairLength)
	return builder.Offset()
}
