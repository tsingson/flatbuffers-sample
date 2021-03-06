package test_test

import (
	"testing"

	"github.com/tsingson/flatbuffers-sample/union-vector/Movie"

	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

func TestMovieUnionField1(t *testing.T) {
	str := &Movie.AttackerT{
		SwordAttackDamage: 100,
	}
	c1 := &Movie.CharacterT{
		Type:  Movie.CharacterMuLan,
		Value: str,
	}
	m := &Movie.MovieT{
		Single: c1,
	}
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()
	mt := Movie.GetRootAsMovie(buf, 0)
	mct := mt.SingleType()

	if mct != Movie.CharacterMuLan {
		t.Errorf("error")
	}
	tb := &flatbuffers.Table{}
	if mt.Single(tb) {
		mct := mt.SingleType()
		ct := mct.UnPack(*tb)
		if ct.Type != mct {
			t.Errorf("error")
		}
		if ct.Value.(*Movie.AttackerT).SwordAttackDamage != int32(100) {
			t.Errorf("error ")
		}
	}
}

func TestMovieUnionField2(t *testing.T) {
	str := "hello world"
	c1 := &Movie.CharacterT{
		Type:  Movie.CharacterOther,
		Value: str,
	}
	m := &Movie.MovieT{
		Single: c1,
	}
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()
	mt := Movie.GetRootAsMovie(buf, 0)
	mct := mt.SingleType()
	if mct != Movie.CharacterOther {
		t.Errorf("error ")
	}
	tb := &flatbuffers.Table{}
	if mt.Single(tb) {
		mct := mt.SingleType()
		ct := mct.UnPack(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(string) != str {
			t.Errorf("error ")
		}
	}
}

func TestMovieUnionField3(t *testing.T) {
	r := &Movie.RapunzelT{
		HairLength: 1,
	}
	c2 := &Movie.CharacterT{
		Type:  Movie.CharacterRapunzel,
		Value: r,
	}
	m := &Movie.MovieT{
		Single: c2,
	}
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()
	mt := Movie.GetRootAsMovie(buf, 0)
	mct := mt.SingleType()
	if mct != Movie.CharacterRapunzel {
		t.Errorf("error ")
	}
	tb := &flatbuffers.Table{}
	if mt.Single(tb) {
		mct := mt.SingleType()
		ct := mct.UnPack(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(*Movie.RapunzelT).HairLength != int32(1) {
			t.Errorf("error ")
		}
	}
}

func TestMovieUnionField4(t *testing.T) {
	r := &Movie.BookReaderT{
		BooksRead: 100,
	}
	c2 := &Movie.CharacterT{
		Type:  Movie.CharacterBookFan,
		Value: r,
	}
	m := &Movie.MovieT{
		Single: c2,
	}
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()
	mt := Movie.GetRootAsMovie(buf, 0)
	mct := mt.SingleType()
	if mct != Movie.CharacterBookFan {
		t.Errorf("error ")
	}
	tb := &flatbuffers.Table{}
	if mt.Single(tb) {
		mct := mt.SingleType()
		ct := mct.UnPack(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(*Movie.BookReaderT).BooksRead != int32(100) {
			t.Errorf("error ")
		}
	}
}

func TestMovieVectorUnions1(t *testing.T) {
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
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()

	mt := Movie.GetRootAsMovie(buf, 0)

	tb := &flatbuffers.Table{}
	if mt.Single(tb) {
		mct := mt.SingleType()
		ct := mct.UnPack(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(string) != str {
			t.Errorf("error ")
		}
	}
	if mt.Multiple(0, tb) {
		mct := mt.MultipleType(0)
		ct := mct.UnPackVector(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(string) != str {
			t.Errorf("error ")
		}
	}
	if mt.Multiple(2, tb) {
		mct := mt.MultipleType(2)
		ct := mct.UnPackVector(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(string) != str {
			t.Errorf("error ")
		}
	}
	if mt.Multiple(1, tb) {
		mct := mt.MultipleType(1)
		ct := mct.UnPackVector(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(*Movie.RapunzelT).HairLength != int32(100) {
			t.Errorf("error ")
		}
	}
	if mt.Multiple(3, tb) {
		mct := mt.MultipleType(3)
		ct := mct.UnPackVector(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(*Movie.BookReaderT).BooksRead != int32(100) {
			t.Errorf("error ")
		}
	}
	if mt.Multiple(4, tb) {
		mct := mt.MultipleType(4)
		ct := mct.UnPackVector(*tb)
		if ct.Type != mct {
			t.Errorf("error ")
		}
		if ct.Value.(*Movie.AttackerT).SwordAttackDamage != int32(100) {
			t.Errorf("error ")
		}
	}
}

func TestMovieVectorUnions2(t *testing.T) {
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
	fb := flatbuffers.NewBuilder(0)
	fb.Finish(m.Pack(fb))
	buf := fb.FinishedBytes()
	mt := Movie.GetRootAsMovie(buf, 0)

	mvt := mt.UnPack()
	if mvt.Single.Type != Movie.CharacterOther {
		t.Errorf("error ")
	}
	if mvt.Single.Value.(string) != str {
		t.Errorf("error ")
	}

	l := len(mvt.Multiple)
	for i := 0; i < l; i++ {
		ct := mvt.Multiple[i]
		switch ct.Type {
		case Movie.CharacterMuLan:

			if ct.Value.(*Movie.AttackerT).SwordAttackDamage != int32(100) {
				t.Errorf("error ")
			}
		case Movie.CharacterRapunzel:
			if ct.Value.(*Movie.RapunzelT).HairLength != int32(100) {
				t.Errorf("error ")
			}
		case Movie.CharacterBelle:

		case Movie.CharacterBookFan:
			if ct.Value.(*Movie.BookReaderT).BooksRead != int32(100) {
				t.Errorf("error ")
			}
		case Movie.CharacterOther:
			if ct.Value.(string) != str {
				t.Errorf("error ")
			}
		case Movie.CharacterUnused:
		}
	}
}
