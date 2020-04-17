// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
	MyGame__Sample "github.com/tsingson/flatbuffers-sample/namespace/MyGame/Sample"
	weapons "github.com/tsingson/flatbuffers-sample/namespace/weapons"
)

type Character byte

const (
	CharacterNONE    Character = 0
	Characterweapon  Character = 1
	CharacterMonster Character = 2

	CharacterVerifyValueMin Character = 0
	CharacterVerifyValueMax Character = 2
)

var EnumNamesCharacter = map[Character]string{
	CharacterNONE:    "NONE",
	Characterweapon:  "weapon",
	CharacterMonster: "Monster",
}

var EnumValuesCharacter = map[string]Character{
	"NONE":    CharacterNONE,
	"weapon":  Characterweapon,
	"Monster": CharacterMonster,
}

func (v Character) String() string {
	if s, ok := EnumNamesCharacter[v]; ok {
		return s
	}
	return "Character(" + strconv.FormatInt(int64(v), 10) + ")"
}

type CharacterT struct {
	Type  Character
	Value interface{}
}

func (t *CharacterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case Characterweapon:
		return t.Value.(*weapons.GunT).Pack(builder)
	case CharacterMonster:
		return t.Value.(*MyGame__Sample.MonsterT).Pack(builder)
	}
	return 0
}

// UnPack use for single union field
func (rcv Character) UnPack(table flatbuffers.Table) *CharacterT {
	switch rcv {
	case Characterweapon:
		x := weapons.GetTableAsGun(&table)
		return &CharacterT{Type: Characterweapon, Value: x.UnPack()}
	case CharacterMonster:
		x := MyGame__Sample.GetTableAsMonster(&table)
		return &CharacterT{Type: CharacterMonster, Value: x.UnPack()}
	}
	return nil
}

// UnPackVector use for vector of unions
func (rcv Character) UnPackVector(table flatbuffers.Table) *CharacterT {
	switch rcv {
	case Characterweapon:
		x := weapons.GetTableVectorAsGun(&table)
		return &CharacterT{Type: Characterweapon, Value: x.UnPack()}
	case CharacterMonster:
		x := MyGame__Sample.GetTableVectorAsMonster(&table)
		return &CharacterT{Type: CharacterMonster, Value: x.UnPack()}
	}
	return nil
}
