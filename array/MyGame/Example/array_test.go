package Example

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"testing"
)

func TestCreateItemStruct(t *testing.T) {

	fb := flatbuffers.NewBuilder(0)

	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	fb.Finish(CreateItemStruct(fb, ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte))
	buf := fb.FinishedBytes()
	fmt.Printf("%0b\n", buf[:])
	fmt.Println(" size : ", len(buf))
	fmt.Println(" head: ", fb.Head())
}

func TestCreateItemTable(t *testing.T) {
		as := assert.New(t )
		fb := flatbuffers.NewBuilder(0)
	ist := &ItemTableT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	fb.Finish(ist.Pack (fb))
	buf := fb.FinishedBytes()
	istt := GetRootAsItemTable( buf , 0 )
	// fmt.Printf("%0b\n", buf[:])
	as.Equal(istt.Color(), ColorBlue)
	as.Equal(istt.F32(), ist.F32)
	as.Equal(istt.I8(), ist.I8)
}

func TestFCreateItemStruct(t *testing.T) {
	as := assert.New(t )
	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	fb := FCreateItemStruct(ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte)
	buf := fb.FinishedBytes()
	// fmt.Printf("%0b\n", buf)
	// fmt.Println(" size : ", len(buf))
	// fmt.Println(" length : ", fb.ByteSize())
	// fmt.Println(" type : ", fb.Type())

	tb := &flatbuffers.Table{
		Bytes: buf ,
		Pos: 0,
	}

	istt := GetStructAsItemStruct( tb )
	as.Equal(istt.Bool(), ist.Bool)
	as.Equal(istt.F32(), ist.F32)
	as.Equal(istt.I8(), ist.I8)

}

func FCreateItemStruct(bool bool, u64 uint64, color Color, i8 int8, f32 float32, ubyte byte) flatbuffers.VField {
	builder := flatbuffers.NewStruct()
	builder.Prep(8, 32)
	builder.Pad(7, 25)
	builder.Byte(ubyte, 24)
	builder.Float32(f32, 20)
	builder.Pad(2, 18)
	builder.Int8(i8, 17)
	builder.Int8(int8(color), 16)
	builder.Uint64(u64, 8)
	builder.Pad(7, 1 )
	builder.Bool(bool, 0)
	return builder.StructEnd()
}
