# flatbuffers go compiler change list

## 1. --gen-onefile parameter exception handling

more then one namespace in IDL:

```
// Example IDL file for our monster's schema.

namespace Mygame.Example;

attribute "go_module:gomodule/";

enum Color:byte { Red = 0, Green, Blue = 2 }
union Equipment {   MuLan: Weapon, Weapon,   SpaceShip,   Other: string } // Optionally add more tables.

struct Vec3 {
  x:float;
  y:float;
  z:float;
}

....

namespace Mygame;

attribute "go_module:gomodule/";

table Monster {
  pos:Mygame.Example.Vec3;
  mana:short = 150;
  hp:short = 100;
  name:string;
  names:[string];
  friendly:bool = false (deprecated);
  inventory:[ubyte];
  color:Mygame.Example.Color = Blue;
  weapons:[Mygame.Example.Weapon];
  equipped:Mygame.Example.Equipment;
  vectorOfUnions:[Mygame.Example.Equipment];  // vector of unions
  path:[Mygame.Example.Vec3];
}


root_type Monster;
file_identifier "MNST";

```

generated go code 

```
.
├── Mygame
│   ├── Example
│   │   ├── Color.go
│   │   ├── Equipment.go
│   │   ├── SpaceShip.go
│   │   ├── Vec3.go
│   │   └── Weapon.go
│   └── Monster.go
├── README.md
├── SampleBinary.java
├── go.mod
├── go.sum
├── monster_namespace.go
└── monster_namespace_gomodule.fbs
```

go code in ./Mygame/Example for namespace Mygame.Example,    go package name Example

```
├── Mygame
│   ├── Example
│   │   ├── Color.go
│   │   ├── Equipment.go
│   │   ├── SpaceShip.go
│   │   ├── Vec3.go
│   │   └── Weapon.go
```

file Color.go / Equipment.go / SapceShip.go / Vec3.go / Weapon.go  is one go package name “Example”.

go code in ./Mygame for namespace Mygame, go package name Mygame

for --gen-onefile parameter,  more then one namespace should be generated into onefile. 

for go, more then one package  can’t be stored in one file.

so , more then one namespace with --gen-onefile , flatc compiler return error msg:

```
~/flatbuffers-sample/gomodule      flatc --go --gen-object-api --gen-onefile ./*.fbs 
flatc: error: more the one namespace defined, can't use --gen-onefile parameter in Go
```

