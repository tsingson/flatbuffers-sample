package main

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/tsingson/flatbuffers-sample/union-vector/Movie"
)

func main() {
	fb := flatbuffers.NewBuilder(0)

	Movie.AttackerStart(fb)
	Movie.AttackerAddSwordAttackDamage(fb, int32(100))
	attOff := Movie.AttackerEnd(fb)

	attType := Movie.CharacterMuLan

	Movie.MovieStart(fb)
	Movie.MovieAddMainCharacterType(fb, attType)
	Movie.MovieAddMainCharacter(fb, attOff)

	Movie.MovieStartCharactersTypeVector(fb, 2)

}
