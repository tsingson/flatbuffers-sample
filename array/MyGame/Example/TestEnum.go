// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import "strconv"

type TestEnum int8

const (
	TestEnumA TestEnum = 0
	TestEnumB TestEnum = 1
	TestEnumC TestEnum = 2

	TestEnumVerifyValueMin TestEnum = 0
	TestEnumVerifyValueMax TestEnum = 2
)

var EnumNamesTestEnum = map[TestEnum]string{
	TestEnumA: "A",
	TestEnumB: "B",
	TestEnumC: "C",
}

var EnumValuesTestEnum = map[string]TestEnum{
	"A": TestEnumA,
	"B": TestEnumB,
	"C": TestEnumC,
}

func (v TestEnum) String() string {
	if s, ok := EnumNamesTestEnum[v]; ok {
		return s
	}
	return "TestEnum(" + strconv.FormatInt(int64(v), 10) + ")"
}
