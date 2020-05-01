package Example

import (
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"testing"
)

func TestCreateNestedStruct(t *testing.T) {
	builder := flatbuffers.NewBuilder(64)

	NestedTableStart(builder)

	nestedOffset := CreateNestedStruct(builder, [2]int32{2, 4}, TestEnumC, [2]TestEnum{TestEnumB, TestEnumC}, [2]int64{8, 16}, int8(1),
		[2]int8{1, 2}, int8(1), true)

	fmt.Println("------------------------")
	fmt.Printf("%0b\n", builder.Bytes)
	fmt.Println(len(builder.Bytes))
	fmt.Println("======>", builder.Head())
	fmt.Printf("%0b\n", builder.Bytes[builder.Head():])

	NestedTableAddNested(builder, nestedOffset)
	endOff := NestedTableEnd(builder)

	builder.Finish(endOff)
	buf := builder.FinishedBytes()
	fmt.Println("======>", 16)
	fmt.Printf("%0b\n", buf[16:])
	fmt.Println("------------------------")
	fmt.Printf("%0b\n", buf)
	fmt.Println(len(buf))

	rcv := GetRootAsNestedTable(buf, 0)

	nested := rcv.Nested(nil)

	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println(nested.C())

	fmt.Println("Pos: ", rcv._tab.Pos)

	x := &NestedStruct{}
	x.Init(buf[16:], 0)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println(x.C())

}
