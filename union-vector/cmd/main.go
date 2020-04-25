package main

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
	"github.com/tsingson/flatbuffers-sample/union-vector/Movie"
)

func main() {
	fb := flatbuffers.NewBuilder(0)

	Movie.AttackerStart(fb)
	Movie.AttackerAddSwordAttackDamage(fb, int32(100))
	attOff := Movie.AttackerEnd(fb)

	attType := Movie.CharacterMuLan

	Movie.MovieStart(fb)
	Movie.MovieAddSingleType(fb, attType)
	Movie.MovieAddSingle(fb, attOff)

	Movie.MovieAddMultiple(fb, 2)

}
