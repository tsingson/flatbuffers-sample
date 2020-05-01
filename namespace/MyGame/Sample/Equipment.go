// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Sample

import (
	"strconv"

	weapons "github.com/tsingson/flatbuffers-sample/namespace/weapons"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type Equipment byte

const (
	EquipmentNONE   Equipment = 0
	EquipmentMuLan  Equipment = 1
	EquipmentWeapon Equipment = 2
	EquipmentGun    Equipment = 3

	EquipmentVerifyValueMin Equipment = 0
	EquipmentVerifyValueMax Equipment = 3
)

var EnumNamesEquipment = map[Equipment]string{
	EquipmentNONE:   "NONE",
	EquipmentMuLan:  "MuLan",
	EquipmentWeapon: "Weapon",
	EquipmentGun:    "Gun",
}

var EnumValuesEquipment = map[string]Equipment{
	"NONE":   EquipmentNONE,
	"MuLan":  EquipmentMuLan,
	"Weapon": EquipmentWeapon,
	"Gun":    EquipmentGun,
}

func (v Equipment) String() string {
	if s, ok := EnumNamesEquipment[v]; ok {
		return s
	}
	return "Equipment(" + strconv.FormatInt(int64(v), 10) + ")"
}

type EquipmentT struct {
	Type  Equipment
	Value interface{}
}

func (t *EquipmentT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case EquipmentMuLan:
		return t.Value.(*WeaponT).Pack(builder)
	case EquipmentWeapon:
		return t.Value.(*WeaponT).Pack(builder)
	case EquipmentGun:
		return t.Value.(*weapons.GunT).Pack(builder)
	}
	return 0
}

// UnPack use for single union field
func (rcv Equipment) UnPack(table flatbuffers.Table) *EquipmentT {
	switch rcv {
	case EquipmentMuLan:
		x := GetTableAsWeapon(&table)
		return &EquipmentT{Type: EquipmentMuLan, Value: x.UnPack()}
	case EquipmentWeapon:
		x := GetTableAsWeapon(&table)
		return &EquipmentT{Type: EquipmentWeapon, Value: x.UnPack()}
	case EquipmentGun:
		x := weapons.GetTableAsGun(&table)
		return &EquipmentT{Type: EquipmentGun, Value: x.UnPack()}
	}
	return nil
}

// UnPackVector use for vector of unions
func (rcv Equipment) UnPackVector(table flatbuffers.Table) *EquipmentT {
	switch rcv {
	case EquipmentMuLan:
		x := GetTableVectorAsWeapon(&table)
		return &EquipmentT{Type: EquipmentMuLan, Value: x.UnPack()}
	case EquipmentWeapon:
		x := GetTableVectorAsWeapon(&table)
		return &EquipmentT{Type: EquipmentWeapon, Value: x.UnPack()}
	case EquipmentGun:
		x := weapons.GetTableVectorAsGun(&table)
		return &EquipmentT{Type: EquipmentGun, Value: x.UnPack()}
	}
	return nil
}
