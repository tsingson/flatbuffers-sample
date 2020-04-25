# flatc compiler parameter FAQ ( for Golang )

## 1. scalar field encode to little-endian byte

> scalar field include:
>
> * int8 / int16 / int32 / int64
>
> * uint8 / uint16 / uint32 / uint64
>
> * ubyte(  int8  in go ) / byte
>
> * float32 / float64
>
> * bool ( encode as byte 0x0 as false, 0x1 as true )

float32 ( 1.1 ) be serialize to little-endian byte is 

```
11001101 11001100 10001100 111111
```

## 2. struct

> all field inside struct is scalar only

struct IDL

```
enum Color:byte { Red = 0, Green, Blue = 2 }

struct ItemStruct{
 bool:bool=false;
 u64: uint64;
 color:Color=Red;
 i8: int8;
 f32: float32;
 ubyte:ubyte;
}
```
struct ItemStruct  as go native object is:

```
ItemStructT{
		Bool:  true,                  // field 1,  byte,    1 byte
		U64:   1,                     // field 2,  uint64,  4 byte
		Color: Example.ColorGreen,    // field 3,  ubyte,   1 byte
		I8:    1,                     // field 4,  int8,    1 byte
		F32:   1.1,                   // field 5,  float32, 4 byte
		Ubyte: uint8(0x1),            // field 6,  uint8,   1 byte
}
```

serialized to binary slice as 32 byte:
```
1 0 0 0 0 0 0 0 
1 0 0 0 0 0 0 0 
1 1 0 0 
11001101 11001100 10001100 111111 
1 0 0 0 0 0 0 0
```

binary format

```
// +-----------------------------------+
// | 1                                 |
// +-----------------------------------+ <- bool = true, 1 byte
// | 0 0 0 0 0 0 0                     |
// +-----------------------------------+ <- pad, 7 byte
// | 1 0 0 0 0 0 0 0                   | 
// +-----------------------------------+ <- uint64 = 1, 8 byte
// | 1                                 |
// +-----------------------------------+ <- enum Color = ColorGreen ( 1 ), 1 byte
// | 1                                 |
// +-----------------------------------+ <- int8 = 1, 1 byte
// | 0 0                               |
// +-----------------------------------+ <- pad, 2 byte
// | 11001101 11001100 10001100 111111 |
// +-----------------------------------+ <- float32 = 1, 4 byte
// | 1                                 |
// +-----------------------------------+ <- ubyte = 0x1, 1 byte
// | 0 0 0 0 0 0 0                     |
// +-----------------------------------+ <- pad, 7 byte
```

## 3. table 

> example is all scalar field only, but table supoort all data type

table IDL :

```
enum Color:byte { Red = 0, Green, Blue = 2 }

table ItemTable{
 bool:bool=false;
 u64: uint64;
 color:Color=Red;
 i8: int8;
 f32: float32;
 ubyte:ubyte;
}
```
table  ItemTable as go native object is:

```
ItemTableT{
		Bool:  true,                  // field 1,  byte,    1 byte
		U64:   1,                     // field 2,  uint64,  4 byte
		Color: Example.ColorGreen,    // field 3,  ubyte,   1 byte
		I8:    1,                     // field 4,  int8,    1 byte
		F32:   1.1,                   // field 5,  float32, 4 byte
		Ubyte: uint8(0x1),            // field 6,  uint8,   1 byte
}
```

serial to binary slice ( 56 byte ) 
```
11000 0 0 0 
0 0 0 0 
10000 0 100000 0 
11111 0 10000 0 
1111 0 1110 0 
1000 0 111 0 
10000 0 0 0 
0 0 0 1 
11001101 11001100 10001100 111111 
0 0 1 1 
1 0 0 0 
0 0 0 0 
0 0 0 0 
0 0 0 1
```

binary format
```
// root header start here ( 8 byte )
// +-----------------------------------+
// | 11000 0 0 0                       |
// +-----------------------------------+ <- point to root table, SOFFSET = 24
// | 0 0 0 0                           |
// +-----------------------------------+


// vtable start here ( 16 byte )
// +-----------------------------------+ <- pad, 4 byte
// | 10000 0 100000 0                  | 
// +-----------------------------------+ <- vtable start
//                                       2 byte as vtable size ( 16 byte) 
//                                       2 byte as table size ( 32 byte)
// | 11111 0 10000 0                   |
// +-----------------------------------+ <- field 1 VOffset is 31 / field 2 VOffset is 16
// | 1111 0 1110 0                     |
// +-----------------------------------+ <- field 3 VOffset is 15 / field 4 VOffset is 14
// | 1000 0 111 0                      |
// +-----------------------------------+ <- field 5 VOffset is 8 / field 6 VOffset is 7


// table start here ( 32 byte )
// +-----------------------------------+
// | 10000 0 0 0                       | 
// +-----------------------------------+ <- table start, 4 byte point to vtable start
// | 0 0 0 1                           | 
// +-----------------------------------+ <- field 6 in 7, uint8 = 0x1, 1 byte , pad 3 byte
// | 11001101 11001100 10001100 111111 |
// +-----------------------------------+ <- field 5 in 8, float32 = 1.1, 4 byte
// | 0 0 1 1                           |
// +-----------------------------------+ <- field 4 in 14, int8 = 1, 1 byte 
                                         / field 3 in 15, byte = 1, 1 byte  
                                         / pad 2 byte
// | 1 0 0 0  0 0 0 0                  |
// +-----------------------------------+ <- field 2 in 16, uint64=1, 8 byte
// | 0 0 0 0  0 0 0 1                  |
// +-----------------------------------+ <- field 1 in 31, bool=true, 1 byte ( pad 7 byte)
```



## 4. flatc compiler parameter list in go

```
--go               -g    Generate Go files for tables/structs.
--go-namespace         Generate the overrided namespace in Golang.
--go-import            Generate the overrided import for flatbuffers in Golang
                         (default is "github.com/google/flatbuffers/go").
```



## 5. parameter --go-import

default import flatbuffers in go is  "github.com/google/flatbuffers/go”

when you need a customized go flatbuffers, use “--go-import” parametor to override default go import.

example:

fbs in google/flatbuffers/sample/monster.fbs

```
flatc --go --gen-object-api --gen-mutable  ./*.fbs
```

by default , the generated go code is :

```
// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MonsterT struct {
	Pos *Vec3T
	Mana int16
	Hp int16
	Name string
	Names []string
	Inventory []byte
	Color Color
	Weapons []*WeaponT
	Equipped *EquipmentT
	VectorOfUnions []*EquipmentT
	Path []*Vec3T
}

 ...
```



when i develop a experimental / customized go flatbuffers  in here( to support fixed-length array / structbuffers / unions vector ) :

[github.com/tsingson/goflatbuffers](https://github.com/tsingson/goflatbuffers)

use this --go-import to testing it:

```
flatc --go --gen-object-api  --gen-mutable --go-import github.com/tsingson/goflatbuffers/go  ./*.fbs
```
the generated go code like this:
```
// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/tsingson/goflatbuffers/go"
)

type MonsterT struct {
	Pos *Vec3T
	Mana int16
	Hp int16
	Name string
	Names []string
	Inventory []byte
	Color Color
	Weapons []*WeaponT
	Equipped *EquipmentT
	VectorOfUnions []*EquipmentT
	Path []*Vec3T
}

 ...
```

>Notice:
>
>It is recommended not to use the --go-import parameter in a production, unless you know the worth of what you are doing

