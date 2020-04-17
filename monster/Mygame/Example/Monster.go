// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

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

func (rcv *Monster) MutateMana(n int16) bool {
	return rcv._tab.MutateInt16Slot(6, n)
}

func (rcv *Monster) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 100
}

func (rcv *Monster) MutateHp(n int16) bool {
	return rcv._tab.MutateInt16Slot(8, n)
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

func (rcv *Monster) MutateInventory(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Monster) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

func (rcv *Monster) MutateColor(n Color) bool {
	return rcv._tab.MutateInt8Slot(18, int8(n))
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

func (rcv *Monster) MutateEquippedType(n Equipment) bool {
	return rcv._tab.MutateByteSlot(22, byte(n))
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

func (rcv *Monster) MutateVectorOfUnionsType(j int, n Equipment) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), byte(n))
	}
	return false
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
