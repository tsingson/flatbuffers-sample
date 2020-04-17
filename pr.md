
 ## 1. support vector of unions.    
union support table / struct / strings .

IDL
```
union Character {
  MuLan: Attacker,  // table
  Rapunzel,         // struct
  Belle: BookReader,
  BookFan: BookReader,
  Other: string,   // string 
  Unused: string
}

table Movie {
  main_character: Character;      // union field
  characters: [Character];            // vector of unions 
}
```
 ## 2. support go module via Attribute ( in IDL declaration ). 
 Each file supports   separate configuration like this : attribute "go_module:github.com/tsingson/flatbuffers-sample/go-example/";

weapons.fbs
```
namespace weapons;

attribute "go_module:github.com/tsingson/flatbuffers-sample/samplesNew/";

table Gun {
  damage:short;
  bool:bool;
  name:string;
  names:[string];
}
```

monster.fbs
```
include "../weapons.fbs";

namespace Mygame.Example;

attribute "go_module:github.com/tsingson/flatbuffers-sample/go-example/";

enum Color:byte { Red = 0, Green, Blue = 2 }
union Equipment {   MuLan: Weapon, Weapon, Gun:weapons.Gun, SpaceShip,   Other: string } // Optionally add more tables.

......
```


the output generated go code
```
package Example

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"

	weapons "github.com/tsingson/flatbuffers-sample/samplesNew/weapons"    /// it work!!!  
)

type Equipment byte

..........

```

## 3. add some shortcut func or generated code to make api more clear.

```
	weaponsOffset := flatbuffers.UOffsetT(0)
	if t.Weapons != nil {
		weaponsLength := len(t.Weapons)
		weaponsOffsets := make([]flatbuffers.UOffsetT, weaponsLength)
		for j := weaponsLength - 1; j >= 0; j-- {
			weaponsOffsets[j] = t.Weapons[j].Pack(builder)
		}
		MonsterStartWeaponsVector(builder, weaponsLength)            //////// start
		for j := weaponsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(weaponsOffsets[j])
		}
		weaponsOffset = MonsterEndWeaponsVector(builder, weaponsLength)   /////// end 
	}
```

shortcut for []strings vector 

```
// native object 

	Names []string


// builder

namesOffset := builder.StringsVector( t.Names...)


```

getter for vector of unions
```

func (rcv *Movie) Characters(j int, obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		obj.Pos = a + flatbuffers.UOffsetT(j*4)
		obj.Bytes = rcv._tab.Bytes
		return true
	}
	return false
}

```
so  get  struct or table 

```
// GetStructVectorAsBookReader shortcut to access struct in vector of unions
func GetStructVectorAsBookReader(table *flatbuffers.Table) *BookReader {
	n := flatbuffers.GetUOffsetT(table.Bytes[table.Pos:])
	x := &BookReader{}
	x.Init(table.Bytes, n+ table.Pos)
	return x
}

// GetStructAsBookReader shortcut to access struct in single union field
func GetStructAsBookReader(table *flatbuffers.Table) *BookReader {
	x := &BookReader{}
	x.Init(table.Bytes, table.Pos)
	return x
}

```

for object-api , comments in generated code to make it clear
```

// UnPack use for single union field
 func (rcv Character) UnPack(table flatbuffers.Table) *CharacterT {
	switch rcv {
	case CharacterMuLan:
		x := GetTableAsAttacker(&table)
		return &CharacterT{ Type: CharacterMuLan, Value: x.UnPack() }
 .............

// UnPackVector use for vector of unions 
func (rcv Character) UnPackVector(table flatbuffers.Table) *CharacterT {
	switch rcv {
	case CharacterMuLan:
		x := GetTableVectorAsAttacker(&table)
		return &CharacterT{ Type: CharacterMuLan, Value: x.UnPack() }
	case CharacterRapunzel:
.........
```

## 4. fix namespace is missing, go package name be the blank identifier issue

https://github.com/google/flatbuffers/issues/5832#issuecomment-611320918_


---------

maybe more for Go flatbuffers ...... :)





