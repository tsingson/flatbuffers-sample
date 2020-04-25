// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package monster

import (
	"strconv"

	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type Color int8

const (
	ColorRed   Color = 0
	ColorGreen Color = 1
	ColorBlue  Color = 2

	ColorVerifyValueMin Color = 0
	ColorVerifyValueMax Color = 2
)

var EnumNamesColor = map[Color]string{
	ColorRed:   "Red",
	ColorGreen: "Green",
	ColorBlue:  "Blue",
}

var EnumValuesColor = map[string]Color{
	"Red":   ColorRed,
	"Green": ColorGreen,
	"Blue":  ColorBlue,
}

func (v Color) String() string {
	if s, ok := EnumNamesColor[v]; ok {
		return s
	}
	return "Color(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Equipment byte

const (
	EquipmentNONE      Equipment = 0
	EquipmentMuLan     Equipment = 1
	EquipmentWeapon    Equipment = 2
	EquipmentSpaceShip Equipment = 3
	EquipmentOther     Equipment = 4

	EquipmentVerifyValueMin Equipment = 0
	EquipmentVerifyValueMax Equipment = 4
)

var EnumNamesEquipment = map[Equipment]string{
	EquipmentNONE:      "NONE",
	EquipmentMuLan:     "MuLan",
	EquipmentWeapon:    "Weapon",
	EquipmentSpaceShip: "SpaceShip",
	EquipmentOther:     "Other",
}

var EnumValuesEquipment = map[string]Equipment{
	"NONE":      EquipmentNONE,
	"MuLan":     EquipmentMuLan,
	"Weapon":    EquipmentWeapon,
	"SpaceShip": EquipmentSpaceShip,
	"Other":     EquipmentOther,
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
	case EquipmentSpaceShip:
		return t.Value.(*SpaceShipT).Pack(builder)
	case EquipmentOther:
		return builder.CreateString(t.Value.(string))
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
	case EquipmentSpaceShip:
		x := GetTableAsSpaceShip(&table)
		return &EquipmentT{Type: EquipmentSpaceShip, Value: x.UnPack()}
	case EquipmentOther:
		x := string(table.StringsVector(table.Pos))
		return &EquipmentT{Type: EquipmentOther, Value: x}
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
	case EquipmentSpaceShip:
		x := GetTableVectorAsSpaceShip(&table)
		return &EquipmentT{Type: EquipmentSpaceShip, Value: x.UnPack()}
	case EquipmentOther:
		x := ""
		b := table.ByteVector(table.Pos)
		if b != nil {
			x = string(b)
		}
		return &EquipmentT{Type: EquipmentOther, Value: x}
	}
	return nil
}

type Vec3T struct {
	X float32
	Y float32
	Z float32
}

func (t *Vec3T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	return CreateVec3(builder, t.X, t.Y, t.Z)
}
func (rcv *Vec3) UnPackTo(t *Vec3T) {
	t.X = rcv.X()
	t.Y = rcv.Y()
	t.Z = rcv.Z()
}

func (rcv *Vec3) UnPack() *Vec3T {
	if rcv == nil {
		return nil
	}
	t := &Vec3T{}
	rcv.UnPackTo(t)
	return t
}

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

type MonsterT struct {
	Pos            *Vec3T
	Mana           int16
	Hp             int16
	Name           string
	Names          []string
	Inventory      []byte
	Color          Color
	Weapons        []*WeaponT
	Equipped       *EquipmentT
	VectorOfUnions []*EquipmentT
	Path           []*Vec3T
}

// MonsterT object pack function
func (t *MonsterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
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
	inventoryOffset := flatbuffers.UOffsetT(0)
	if t.Inventory != nil {
		inventoryOffset = builder.CreateByteString(t.Inventory)
	}
	weaponsOffset := flatbuffers.UOffsetT(0)
	if t.Weapons != nil {
		weaponsLength := len(t.Weapons)
		weaponsOffsets := make([]flatbuffers.UOffsetT, weaponsLength)
		for j := weaponsLength - 1; j >= 0; j-- {
			weaponsOffsets[j] = t.Weapons[j].Pack(builder)
		}
		MonsterStartWeaponsVector(builder, weaponsLength)
		for j := weaponsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(weaponsOffsets[j])
		}
		weaponsOffset = MonsterEndWeaponsVector(builder, weaponsLength)
	}
	equippedOffset := t.Equipped.Pack(builder)

	// vector of unions
	vectorOfUnionsOffset := flatbuffers.UOffsetT(0)
	vectorOfUnionsTypeOffset := flatbuffers.UOffsetT(0)
	if t.VectorOfUnions != nil {
		vectorOfUnionsLength := len(t.VectorOfUnions)
		vectorOfUnionsOffsets := make([]flatbuffers.UOffsetT, vectorOfUnionsLength)
		for j := vectorOfUnionsLength - 1; j >= 0; j-- {
			vectorOfUnionsOffsets[j] = t.VectorOfUnions[j].Pack(builder)
		}
		MonsterStartVectorOfUnionsTypeVector(builder, vectorOfUnionsLength)
		for j := vectorOfUnionsLength - 1; j >= 0; j-- {
			builder.PrependByte(byte(t.VectorOfUnions[j].Type))
		}
		vectorOfUnionsTypeOffset = MonsterEndVectorOfUnionsTypeVector(builder, vectorOfUnionsLength)
		MonsterStartVectorOfUnionsVector(builder, vectorOfUnionsLength)
		for j := vectorOfUnionsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(vectorOfUnionsOffsets[j])
		}
		vectorOfUnionsOffset = MonsterEndVectorOfUnionsVector(builder, vectorOfUnionsLength)
	}
	pathOffset := flatbuffers.UOffsetT(0)
	if t.Path != nil {
		pathLength := len(t.Path)
		MonsterStartPathVector(builder, pathLength)
		for j := pathLength - 1; j >= 0; j-- {
			t.Path[j].Pack(builder)
		}
		pathOffset = MonsterEndPathVector(builder, pathLength)
	}

	// pack process all field

	MonsterStart(builder)
	posOffset := t.Pos.Pack(builder)
	MonsterAddPos(builder, posOffset)
	MonsterAddMana(builder, t.Mana)
	MonsterAddHp(builder, t.Hp)
	MonsterAddName(builder, nameOffset)
	MonsterAddNames(builder, namesOffset)
	MonsterAddInventory(builder, inventoryOffset)
	MonsterAddColor(builder, t.Color)
	MonsterAddWeapons(builder, weaponsOffset)
	if t.Equipped != nil {
		MonsterAddEquippedType(builder, t.Equipped.Type)
	}
	MonsterAddEquipped(builder, equippedOffset)
	MonsterAddVectorOfUnionsType(builder, vectorOfUnionsTypeOffset)
	MonsterAddVectorOfUnions(builder, vectorOfUnionsOffset)
	MonsterAddPath(builder, pathOffset)
	return MonsterEnd(builder)
}

// MonsterT object unpack function
func (rcv *Monster) UnPackTo(t *MonsterT) {
	t.Pos = rcv.Pos(nil).UnPack()
	t.Mana = rcv.Mana()
	t.Hp = rcv.Hp()
	t.Name = string(rcv.Name())
	namesLength := rcv.NamesLength()
	t.Names = make([]string, namesLength)
	for j := 0; j < namesLength; j++ {
		t.Names[j] = string(rcv.Names(j))
	}
	t.Inventory = rcv.InventoryBytes()
	t.Color = rcv.Color()
	weaponsLength := rcv.WeaponsLength()
	t.Weapons = make([]*WeaponT, weaponsLength)
	for j := 0; j < weaponsLength; j++ {
		x := Weapon{}
		rcv.Weapons(&x, j)
		t.Weapons[j] = x.UnPack()
	}
	equippedTable := flatbuffers.Table{}
	if rcv.Equipped(&equippedTable) {
		t.Equipped = rcv.EquippedType().UnPack(equippedTable)
	}
	vectorOfUnionsLength := rcv.VectorOfUnionsLength()
	t.VectorOfUnions = make([]*EquipmentT, vectorOfUnionsLength)
	for j := 0; j < vectorOfUnionsLength; j++ {
		// vector of unions table unpack
		VectorOfUnionsType := rcv.VectorOfUnionsType(j)
		VectorOfUnionsTable := flatbuffers.Table{}
		if rcv.VectorOfUnions(j, &VectorOfUnionsTable) {
			t.VectorOfUnions[j] = VectorOfUnionsType.UnPackVector(VectorOfUnionsTable)
		}
	}
	pathLength := rcv.PathLength()
	t.Path = make([]*Vec3T, pathLength)
	for j := 0; j < pathLength; j++ {
		x := Vec3{}
		rcv.Path(&x, j)
		t.Path[j] = x.UnPack()
	}
}

func (rcv *Monster) UnPack() *MonsterT {
	if rcv == nil {
		return nil
	}
	t := &MonsterT{}
	rcv.UnPackTo(t)
	return t
}

type Monster struct {
	_tab flatbuffers.Table
}

// GetRootAsMonster shortcut to access root table
func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsMonster shortcut to access table in vector of  unions
func GetTableVectorAsMonster(table *flatbuffers.Table) *Monster {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &Monster{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsMonster shortcut to access table in single union field
func GetTableAsMonster(table *flatbuffers.Table) *Monster {
	x := &Monster{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Monster) Pos(obj *Vec3) *Vec3 {
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

func (rcv *Monster) Mana() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 150
}

func (rcv *Monster) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 100
}

func (rcv *Monster) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Names(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Monster) NamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Inventory(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Monster) InventoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) InventoryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

func (rcv *Monster) Weapons(obj *Weapon, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) WeaponsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) EquippedType() Equipment {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return Equipment(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Monster) Equipped(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *Monster) VectorOfUnionsType(j int) Equipment {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return Equipment(rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *Monster) VectorOfUnionsTypeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) VectorOfUnions(j int, obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		a := rcv._tab.Vector(o)
		obj.Pos = a + flatbuffers.UOffsetT(j*4)
		obj.Bytes = rcv._tab.Bytes
		return true
	}
	return false
}

func (rcv *Monster) VectorOfUnionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Path(obj *Vec3, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 12
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) PathLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MonsterStart(builder *flatbuffers.Builder) {
	builder.StartObject(14)
}

func MonsterAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0)
}

func MonsterAddMana(builder *flatbuffers.Builder, mana int16) {
	builder.PrependInt16Slot(1, mana, 150)
}

func MonsterAddHp(builder *flatbuffers.Builder, hp int16) {
	builder.PrependInt16Slot(2, hp, 100)
}

func MonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}

func MonsterAddNames(builder *flatbuffers.Builder, names flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(names), 0)
}

func MonsterStartNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func MonsterEndNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterAddInventory(builder *flatbuffers.Builder, inventory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(inventory), 0)
}

func MonsterStartInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}

func MonsterEndInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(7, int8(color), 2)
}

func MonsterAddWeapons(builder *flatbuffers.Builder, weapons flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(weapons), 0)
}

func MonsterStartWeaponsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func MonsterEndWeaponsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterAddEquippedType(builder *flatbuffers.Builder, equippedType Equipment) {
	builder.PrependByteSlot(9, byte(equippedType), 0)
}

func MonsterAddEquipped(builder *flatbuffers.Builder, equipped flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(equipped), 0)
}

func MonsterAddVectorOfUnionsType(builder *flatbuffers.Builder, vectorOfUnionsType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(vectorOfUnionsType), 0)
}

func MonsterStartVectorOfUnionsTypeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}

func MonsterEndVectorOfUnionsTypeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterAddVectorOfUnions(builder *flatbuffers.Builder, vectorOfUnions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(vectorOfUnions), 0)
}

func MonsterStartVectorOfUnionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}

func MonsterEndVectorOfUnionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(path), 0)
}

func MonsterStartPathVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(12, numElems, 4)
}

func MonsterEndPathVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type WeaponT struct {
	Color Color
	Power int32
	Name  string
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
	WeaponAddColor(builder, t.Color)
	WeaponAddPower(builder, t.Power)
	WeaponAddName(builder, nameOffset)
	return WeaponEnd(builder)
}

// WeaponT object unpack function
func (rcv *Weapon) UnPackTo(t *WeaponT) {
	t.Color = rcv.Color()
	t.Power = rcv.Power()
	t.Name = string(rcv.Name())
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

func (rcv *Weapon) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Weapon) Power() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Weapon) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func WeaponStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}

func WeaponAddColor(builder *flatbuffers.Builder, color Color) {
	builder.PrependInt8Slot(0, int8(color), 0)
}

func WeaponAddPower(builder *flatbuffers.Builder, power int32) {
	builder.PrependInt32Slot(1, power, 0)
}

func WeaponAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}

func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

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

func SpaceShipAddSize(builder *flatbuffers.Builder, size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(size), 0)
}

func SpaceShipAddPower(builder *flatbuffers.Builder, power int32) {
	builder.PrependInt32Slot(1, power, 0)
}

func SpaceShipAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0)
}

func SpaceShipEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
