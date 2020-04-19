# fbs support multiple namespace / go module

same as https://github.com/google/flatbuffers/samples but support multiple namespace and attribute for go module 


## 1. setup go repo ( local go module )

initial a new go project call gomodule

```
cd ~
mkdir gomodule
cd gomodule
go init gomodule
```

success init a go project and support go module call "gomodule" 


## 2. fbs IDL file 

IDL file name monster_namespace_gomodule.fbs

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

table Weapon {
  name:string;
  damage:short;
}

table SpaceShip {
 size:Vec3;
 power:int32;
 name:string;
}

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

there 2 namespace
```
namespace Mygame.Example;

...

struct Vec3 {
  x:float;
  y:float;
  z:float;
}
...

namespace Mygame;

...

table Monster {
  pos:Mygame.Example.Vec3; // reference to pre-define namespace 
  
  ....
  
  
```


attribute for go module setting

```
attribute "go_module:gomodule/";
```


## 3. compile fbs

```
flatc --go --gen-object-api --gen-mutable ./monster_namespace_gomodule.fbs 
```


## 4. run example go code

