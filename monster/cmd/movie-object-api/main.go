package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/tsingson/flatbuffers-sample/union-vector/Movie"
)

func main() {
	str := "hello world"
	c1 := &Movie.CharacterT{
		Type:  Movie.CharacterOther,
		Value: str,
	}
	r := &Movie.RapunzelT{
		HairLength: 100,
	}
	c2 := &Movie.CharacterT{
		Type:  Movie.CharacterRapunzel,
		Value: r,
	}

	r1 := &Movie.BookReaderT{
		BooksRead: 100,
	}
	c3 := &Movie.CharacterT{
		Type:  Movie.CharacterBookFan,
		Value: r1,
	}
	ak := &Movie.AttackerT{
		SwordAttackDamage: 100,
	}
	c4 := &Movie.CharacterT{
		Type:  Movie.CharacterMuLan,
		Value: ak,
	}

	m := &Movie.MovieT{
		Single:   c1,
		Multiple: []*Movie.CharacterT{c1, c2, c1, c3, c4},
	}
	builder := flatbuffers.NewBuilder(0)
	builder.Finish(m.Pack(builder))
	buf := builder.FinishedBytes()

	mt := Movie.GetRootAsMovie(buf, 0)

	mvt := mt.UnPack()
	if mvt.Single.Type == Movie.CharacterOther {
		fmt.Println(mvt.Single.Value.(string))
	}

	l := len(mvt.Multiple)
	for i := 0; i < l; i++ {
		ct := mvt.Multiple[i]
		switch ct.Type {
		case Movie.CharacterMuLan:

		// case Movie.CharacterAttacker:
		// 	if ct.Value.(*Movie.AttackerT).SwordAttackDamage == int32(100) {
		// 		fmt.Println(ct.Type)
		// 		fmt.Println(ct.Value.(*Movie.AttackerT).SwordAttackDamage)
		// 	}
		case Movie.CharacterRapunzel:
			fmt.Println(ct.Type)
			fmt.Println(ct.Value.(*Movie.RapunzelT).HairLength, " is ", int32(100))
		case Movie.CharacterBelle:

		case Movie.CharacterBookFan:
			fmt.Println(ct.Type)
			fmt.Println(ct.Value.(*Movie.BookReaderT).BooksRead, int32(100))
		case Movie.CharacterOther:
			fmt.Println(ct.Type)
			fmt.Println(ct.Value.(string), str)
		case Movie.CharacterUnused:
		}
	}
}
