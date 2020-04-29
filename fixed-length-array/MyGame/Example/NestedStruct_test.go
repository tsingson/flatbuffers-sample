package Example

import (
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"testing"
)

func TestCreateNestedStruct(t *testing.T) {
	builder := flatbuffers.NewBuilder(64)
	off :=  CreateNestedStruct(builder, [2]int32{ 1, 2 },TestEnumB, [2]TestEnum{TestEnumB, TestEnumC}, [2]int64{3,4}, int8(1),
		[2]int8{1,2 }, int8(1), true )

	builder.Finish(off )
	buf := builder.FinishedBytes()
	fmt.Printf("%0b\n", buf )

}
