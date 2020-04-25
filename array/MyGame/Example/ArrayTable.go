// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayTableT struct {
	Bool bool
	BoolList []bool
	Color Color
	ColorList []Color
	I8 int8
	I8List []int8
	F32 float32
	F32List []float32
	String_ string
	StrList []string
	Struct_ *ItemStructT
	StructList []*ItemStructT
	Table_ *ItemTableT
	TableList []*ItemTableT
}

// ArrayTableT object pack function
func (t *ArrayTableT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	boolListOffset := flatbuffers.UOffsetT(0)
	if t.BoolList != nil {
		boolListLength := len(t.BoolList)
		ArrayTableStartBoolListVector(builder, boolListLength)
		for j := boolListLength - 1; j >= 0; j-- {
			builder.PrependBool(t.BoolList[j])
		}
		boolListOffset = ArrayTableEndBoolListVector(builder, boolListLength)
	}
	colorListOffset := flatbuffers.UOffsetT(0)
	if t.ColorList != nil {
		colorListLength := len(t.ColorList)
		ArrayTableStartColorListVector(builder, colorListLength)
		for j := colorListLength - 1; j >= 0; j-- {
			builder.PrependInt8(int8(t.ColorList[j]))
		}
		colorListOffset = ArrayTableEndColorListVector(builder, colorListLength)
	}
	i8ListOffset := flatbuffers.UOffsetT(0)
	if t.I8List != nil {
		i8ListLength := len(t.I8List)
		ArrayTableStartI8ListVector(builder, i8ListLength)
		for j := i8ListLength - 1; j >= 0; j-- {
			builder.PrependInt8(t.I8List[j])
		}
		i8ListOffset = ArrayTableEndI8ListVector(builder, i8ListLength)
	}
	f32ListOffset := flatbuffers.UOffsetT(0)
	if t.F32List != nil {
		f32ListLength := len(t.F32List)
		ArrayTableStartF32ListVector(builder, f32ListLength)
		for j := f32ListLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.F32List[j])
		}
		f32ListOffset = ArrayTableEndF32ListVector(builder, f32ListLength)
	}
	String_Offset := flatbuffers.UOffsetT(0)
	if len(t.String_) > 0 {
		String_Offset = builder.CreateString(t.String_)
	}
	strListOffset := flatbuffers.UOffsetT(0)
	if t.StrList != nil {
		strListOffset = builder.StringsVector(t.StrList...)
	}
	structListOffset := flatbuffers.UOffsetT(0)
	if t.StructList != nil {
		structListLength := len(t.StructList)
		ArrayTableStartStructListVector(builder, structListLength)
		for j := structListLength - 1; j >= 0; j-- {
			t.StructList[j].Pack(builder)
		}
		structListOffset = ArrayTableEndStructListVector(builder, structListLength)
	}
	Table_Offset := t.Table_.Pack(builder)
	// vector of tables 
	tableListOffset := flatbuffers.UOffsetT(0)
	if t.TableList != nil {
		tableListLength := len(t.TableList)
		tableListOffsets := make([]flatbuffers.UOffsetT, tableListLength)
		for j := tableListLength - 1; j >= 0; j-- {
			tableListOffsets[j] = t.TableList[j].Pack(builder)
		}
		ArrayTableStartTableListVector(builder, tableListLength)
		for j := tableListLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(tableListOffsets[j])
		}
		tableListOffset = ArrayTableEndTableListVector(builder, tableListLength)
	}

	// pack process all field

	ArrayTableStart(builder)
	ArrayTableAddBool(builder, t.Bool)
	ArrayTableAddBoolList(builder, boolListOffset)
	ArrayTableAddColor(builder, t.Color)
	ArrayTableAddColorList(builder, colorListOffset)
	ArrayTableAddI8(builder, t.I8)
	ArrayTableAddI8List(builder, i8ListOffset)
	ArrayTableAddF32(builder, t.F32)
	ArrayTableAddF32List(builder, f32ListOffset)
	ArrayTableAddString_(builder, String_Offset)
	ArrayTableAddStrList(builder, strListOffset)
	Struct_Offset := t.Struct_.Pack(builder)
	ArrayTableAddStruct_(builder, Struct_Offset)
	ArrayTableAddStructList(builder, structListOffset)
	ArrayTableAddTable_(builder, Table_Offset)
	ArrayTableAddTableList(builder, tableListOffset)
	return ArrayTableEnd(builder)
}

// ArrayTableT object unpack function
func (rcv *ArrayTable) UnPackTo(t *ArrayTableT) {
	t.Bool = rcv.Bool()
	boolListLength := rcv.BoolListLength()
	t.BoolList = make([]bool, boolListLength)
	for j := 0; j < boolListLength; j++ {
		t.BoolList[j] = rcv.BoolList(j)	}
	t.Color = rcv.Color()
	colorListLength := rcv.ColorListLength()
	t.ColorList = make([]Color, colorListLength)
	for j := 0; j < colorListLength; j++ {
		t.ColorList[j] = rcv.ColorList(j)	}
	t.I8 = rcv.I8()
	i8ListLength := rcv.I8ListLength()
	t.I8List = make([]int8, i8ListLength)
	for j := 0; j < i8ListLength; j++ {
		t.I8List[j] = rcv.I8List(j)	}
	t.F32 = rcv.F32()
	f32ListLength := rcv.F32ListLength()
	t.F32List = make([]float32, f32ListLength)
	for j := 0; j < f32ListLength; j++ {
		t.F32List[j] = rcv.F32List(j)	}
	t.String_ = string(rcv.String_())
	strListLength := rcv.StrListLength()
	t.StrList = make([]string, strListLength)
	for j := 0; j < strListLength; j++ {
		t.StrList[j] = string(rcv.StrList(j))
	}
	t.Struct_ = rcv.Struct_(nil).UnPack()
	structListLength := rcv.StructListLength()
	t.StructList = make([]*ItemStructT, structListLength)
	for j := 0; j < structListLength; j++ {
		x := ItemStruct{}
		rcv.StructList(&x, j)
		t.StructList[j] = x.UnPack()
	}
	t.Table_ = rcv.Table_(nil).UnPack()
	tableListLength := rcv.TableListLength()
	t.TableList = make([]*ItemTableT, tableListLength)
	for j := 0; j < tableListLength; j++ {
		x := ItemTable{}
		rcv.TableList(&x, j)
		t.TableList[j] = x.UnPack()
	}
}

func (rcv *ArrayTable) UnPack() *ArrayTableT {
	if rcv == nil {
		return nil
	}
	t := &ArrayTableT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayTable struct {
	_tab flatbuffers.Table
}

// GetRootAsArrayTable shortcut to access root table
func GetRootAsArrayTable(buf []byte, offset flatbuffers.UOffsetT) *ArrayTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayTable{}
	x.Init(buf, n+offset)
	return x
}

// GetTableVectorAsArrayTable shortcut to access table in vector of  unions
func GetTableVectorAsArrayTable(table *flatbuffers.Table) *ArrayTable {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &ArrayTable{}
	x.Init(table.Bytes, n+table.Pos)
	return x
}

// GetTableAsArrayTable shortcut to access table in single union field
func GetTableAsArrayTable(table *flatbuffers.Table) *ArrayTable {
	x := &ArrayTable{}
	x.Init(table.Bytes, table.Pos)
	return x
}

func (rcv *ArrayTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayTable) Bool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ArrayTable) BoolListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) BoolList(j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetBool(a + flatbuffers.UOffsetT(j*1))
	}
	return false
}

func (rcv *ArrayTable) Color() Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return Color(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ArrayTable) ColorListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) ColorList(j int) Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return Color(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *ArrayTable) I8() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ArrayTable) I8ListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) I8List(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ArrayTable) F32() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ArrayTable) F32ListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) F32List(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ArrayTable) String_() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ArrayTable) StrListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) StrList(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ArrayTable) Struct_(obj *ItemStruct) *ItemStruct {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ItemStruct)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ArrayTable) StructListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) StructList(obj *ItemStruct, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 32
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ArrayTable) Table_(obj *ItemTable) *ItemTable {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ItemTable)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ArrayTable) TableListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayTable) TableList(obj *ItemTable, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func ArrayTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(14)
}

func ArrayTableAddBool(builder *flatbuffers.Builder, Bool bool) {
	builder.PrependBoolSlot(0, Bool, false)
}

func ArrayTableStartBoolListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(1, numElems, 1)
}

func ArrayTableEndBoolListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddBoolList(builder *flatbuffers.Builder, BoolList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(BoolList), 0)
}

func ArrayTableAddColor(builder *flatbuffers.Builder, Color Color) {
	builder.PrependInt8Slot(2, int8(Color), 0)
}

func ArrayTableStartColorListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(1, numElems, 1)
}

func ArrayTableEndColorListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddColorList(builder *flatbuffers.Builder, ColorList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(ColorList), 0)
}

func ArrayTableAddI8(builder *flatbuffers.Builder, I8 int8) {
	builder.PrependInt8Slot(4, I8, 0)
}

func ArrayTableStartI8ListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(1, numElems, 1)
}

func ArrayTableEndI8ListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddI8List(builder *flatbuffers.Builder, I8List flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(I8List), 0)
}

func ArrayTableAddF32(builder *flatbuffers.Builder, F32 float32) {
	builder.PrependFloat32Slot(6, F32, 0.0)
}

func ArrayTableStartF32ListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(4, numElems, 4)
}

func ArrayTableEndF32ListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddF32List(builder *flatbuffers.Builder, F32List flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(F32List), 0)
}

func ArrayTableAddString_(builder *flatbuffers.Builder, String_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(String_), 0)
}

func ArrayTableStartStrListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(4, numElems, 4)
}

func ArrayTableEndStrListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddStrList(builder *flatbuffers.Builder, StrList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(StrList), 0)
}

func ArrayTableAddStruct_(builder *flatbuffers.Builder, Struct_ flatbuffers.UOffsetT) {
	builder.PrependStructSlot(10, flatbuffers.UOffsetT(Struct_), 0)
}

func ArrayTableStartStructListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(32, numElems, 8)
}

func ArrayTableEndStructListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddStructList(builder *flatbuffers.Builder, StructList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(StructList), 0)
}

func ArrayTableAddTable_(builder *flatbuffers.Builder, Table_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(Table_), 0)
}

func ArrayTableStartTableListVector(builder *flatbuffers.Builder, numElems int) {
	builder.StartVector(4, numElems, 4)
}

func ArrayTableEndTableListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.EndVector(numElems)
}

func ArrayTableAddTableList(builder *flatbuffers.Builder, TableList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(TableList), 0)
}

func ArrayTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
