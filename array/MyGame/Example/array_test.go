package Example

import (
	"github.com/stretchr/testify/assert"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"runtime"
	"testing"
)

/**
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
	// fmt.Println(" size : ", len(buf))
	// fmt.Println(" head: ", fb.Head())
}


*/
func TestCreateItemTable(t *testing.T) {
	as := assert.New(t)
	builder := flatbuffers.NewBuilder(56)
	ist := &ItemTableT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	ItemTableStart(builder)
	ItemTableAddBool(builder, ist.Bool)
	ItemTableAddU64(builder, ist.U64)
	ItemTableAddColor(builder, ist.Color)
	ItemTableAddI8(builder, ist.I8)
	ItemTableAddF32(builder, ist.F32)
	ItemTableAddUbyte(builder, ist.Ubyte)
	builder.Finish(ItemTableEnd(builder))

	buf := builder.FinishedBytes()
	istt := GetRootAsItemTable(buf, 0)
	// fmt.Printf("%0b\n", buf[:])
	// fmt.Println("len: ", len(buf))
	as.Equal(istt.Color(), ColorBlue)
	as.Equal(istt.F32(), ist.F32)
	as.Equal(istt.I8(), ist.I8)
}

func BenchmarkCreateItemTable(b *testing.B) {

	ist := &ItemTableT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
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
			ItemTableStart(builder)
			ItemTableAddBool(builder, ist.Bool)
			ItemTableAddU64(builder, ist.U64)
			ItemTableAddColor(builder, ist.Color)
			ItemTableAddI8(builder, ist.I8)
			ItemTableAddF32(builder, ist.F32)
			ItemTableAddUbyte(builder, ist.Ubyte)
			builder.Finish(ItemTableEnd(builder))
			builder.FinishedBytes()
		}
	})

}

func TestFCreateItemStruct(t *testing.T) {
	as := assert.New(t)
	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}
	fsb := flatbuffers.NewStruct()
	fb := FCreateItemStruct(fsb, ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte)
	buf := fb.FinishedBytes()
	// fmt.Printf("%0b\n", buf)

	tb := &flatbuffers.Table{
		Bytes: buf,
		Pos:   0,
	}

	istt := GetStructAsItemStruct(tb)
	as.Equal(istt.Bool(), ist.Bool)
	as.Equal(istt.F32(), ist.F32)
	as.Equal(istt.I8(), ist.I8)

}

func BenchmarkFCreateItemStruct1(b *testing.B) {
	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}

	fsb := flatbuffers.NewStruct()
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FCreateItemStruct(fsb, ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte).FinishedBytes()
		}
	})
}
func BenchmarkFCreateItemStruct2(b *testing.B) {
	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}

	fsb := flatbuffers.NewFixedStruct(32)
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FCreateItemStruct(fsb, ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte).FinishedBytes()
		}
	})
}
func BenchmarkFCreateItemStruct3(b *testing.B) {
	ist := &ItemStructT{
		Bool:  true,
		U64:   100,
		Color: ColorBlue,
		I8:    7,
		F32:   32,
		Ubyte: uint8(0x3),
	}

	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// fsb := flatbuffers.NewStruct()
			fsb := flatbuffers.NewFixedStruct(32)
			FCreateItemStruct(fsb, ist.Bool, ist.U64, ist.Color, ist.I8, ist.F32, ist.Ubyte).FinishedBytes()
		}
	})
}

func BenchmarkFCreateItemStruct4(b *testing.B) {
	fsb := flatbuffers.NewFixedStruct(2)
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// fsb := flatbuffers.NewStruct()

			FCreateItemStruct1(fsb, true, ColorBlue).FinishedBytes()
		}
	})
}
func BenchmarkFCreateItemStruct5(b *testing.B) {
	b.SetParallelism(runtime.NumCPU() * 4)
	b.StartTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// fsb := flatbuffers.NewStruct()
			fsb := flatbuffers.NewFixedStruct(2)
			FCreateItemStruct1(fsb, true, ColorBlue).FinishedBytes()
		}
	})
}

func FCreateItemStruct(builder *flatbuffers.StructBuffers, bool bool, u64 uint64, color Color, i8 int8, f32 float32, ubyte byte) flatbuffers.VField {

	builder.Prep(8, 32)
	builder.Pad(7, 25)
	builder.Byte(ubyte, 24)
	builder.Float32(f32, 20)
	builder.Pad(2, 18)
	builder.Int8(i8, 17)
	builder.Int8(int8(color), 16)
	builder.Uint64(u64, 8)
	builder.Pad(7, 1)
	builder.Bool(bool, 0)
	return builder.StructEnd()
}

func FCreateItemStruct1(builder *flatbuffers.StructBuffers, Bool bool, Color Color) flatbuffers.VField {
	builder.Prep(1, 2)
	builder.Int8(int8(Color), 1)
	builder.Bool(Bool, 0)
	return builder.StructEnd()
}
