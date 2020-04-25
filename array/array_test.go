package array

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"github.com/tsingson/flatbuffers-sample/array/MyGame/Example"
	"runtime"
	"testing"
)

func TestCreateItemTable(t *testing.T) {
	as := assert.New(t)
	builder := flatbuffers.NewBuilder(56)
	ist := &Example.ItemTableT{
		Bool:  true,
		U64:   100,
		Color: Example.ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	Example.ItemTableStart(builder)
	Example.ItemTableAddBool(builder, ist.Bool)
	Example.ItemTableAddU64(builder, ist.U64)
	Example.ItemTableAddColor(builder, ist.Color)
	Example.ItemTableAddI8(builder, ist.I8)
	Example.ItemTableAddF32(builder, ist.F32)
	Example.ItemTableAddUbyte(builder, ist.Ubyte)
	builder.Finish(Example.ItemTableEnd(builder))

	buf := builder.FinishedBytes()
	istt := Example.GetRootAsItemTable(buf, 0)
	// fmt.Printf("%0b\n", buf[:])
	// fmt.Println("len: ", len(buf))
	as.Equal(istt.Color(), Example.ColorBlue)
	as.Equal(istt.F32(), ist.F32)
	as.Equal(istt.I8(), ist.I8)
}

func BenchmarkCreateItemTable(b *testing.B) {

	ist := &Example.ItemTableT{
		Bool:  true,
		U64:   100,
		Color: Example.ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			builder := flatbuffers.NewBuilder(128)
			Example.ItemTableStart(builder)
			Example.ItemTableAddBool(builder, ist.Bool)
			Example.ItemTableAddU64(builder, ist.U64)
			Example.ItemTableAddColor(builder, ist.Color)
			Example.ItemTableAddI8(builder, ist.I8)
			Example.ItemTableAddF32(builder, ist.F32)
			Example.ItemTableAddUbyte(builder, ist.Ubyte)
			builder.Finish(Example.ItemTableEnd(builder))
			builder.FinishedBytes()
		}
	})

}
