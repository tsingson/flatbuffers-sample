package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tsingson/flatbuffers-sample/array/MyGame/Example"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"testing"
)

func TestUChar_InventoryBytes(t *testing.T) {
	as := assert.New(t)
	uct := &Example.UCharT{
		Inventory:  []byte("12345"),
		Inventory1: []int8{0x1, 0x2, 0x3, 0x4, 0x5},
	}
	builder := flatbuffers.NewBuilder(128)

	builder.Finish(uct.Pack(builder))

	buf := builder.FinishedBytes()

	uctt := Example.GetRootAsUChar(buf, 0)

	fmt.Println("byte: ", uctt.InventoryBytes())
	fmt.Println("byte len: ", uctt.InventoryLength())
	as.Equal(uctt.InventoryLength(), len(uct.Inventory))
	as.Equal(uctt.InventoryBytes(), uct.Inventory)
	fmt.Printf("%v \n", uct.Inventory)
	fmt.Printf("%v\n", uctt.InventoryBytes())

	fmt.Println("byte len: ", uctt.Inventory1Length())
	int8Array := make([]int8, uctt.Inventory1Length())
	for i := 0; i < uctt.Inventory1Length(); i++ {
		int8Array[i] = uctt.Inventory1(i)
	}

	as.Equal(uct.Inventory1, int8Array)

	fmt.Printf("%0b\n", int8Array)
	fmt.Printf("%0b\n", uct.Inventory1)
}
