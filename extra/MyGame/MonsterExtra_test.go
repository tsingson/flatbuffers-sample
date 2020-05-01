package MyGame

import (
	"fmt"
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"testing"
)

func TestGetRootAsMonsterExtra(t *testing.T) {
	builder := flatbuffers.NewBuilder(256)

	dvecOffset := flatbuffers.UOffsetT(0)

	Dvec := []float64{1.1, 2.2, 3.3, 4.4}
	Fvec := []float32{1.1, 2.2, 3.3, 4.4}

	if Dvec != nil {
		dvecLength := len(Dvec)
		MonsterExtraStartDvecVector(builder, dvecLength)
		for j := dvecLength - 1; j >= 0; j-- {
			builder.PrependFloat64(Dvec[j])
		}
		dvecOffset = MonsterExtraEndDvecVector(builder, dvecLength)
	}
	fvecOffset := flatbuffers.UOffsetT(0)
	if Fvec != nil {
		fvecLength := len(Fvec)
		MonsterExtraStartFvecVector(builder, fvecLength)
		for j := fvecLength - 1; j >= 0; j-- {
			builder.PrependFloat32(Fvec[j])
		}
		fvecOffset = MonsterExtraEndFvecVector(builder, fvecLength)
	}

	// pack process all field

	MonsterExtraStart(builder)
	MonsterExtraAddD0(builder, 1.1)
	MonsterExtraAddD1(builder, 2.2)
	MonsterExtraAddD2(builder, 3.2)
	MonsterExtraAddD3(builder, 4.4)
	MonsterExtraAddF0(builder, 5.5)
	MonsterExtraAddF1(builder, 6.6)
	MonsterExtraAddF2(builder, 7.7)
	MonsterExtraAddF3(builder, 8.8)
	MonsterExtraAddDvec(builder, dvecOffset)
	MonsterExtraAddFvec(builder, fvecOffset)
	builder.Finish(MonsterExtraEnd(builder))
	buf := builder.FinishedBytes()

	fmt.Printf("%0b\n", buf)

}
