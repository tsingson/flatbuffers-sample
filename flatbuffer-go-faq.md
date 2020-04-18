# FAQ for Use FlatBuffers in Go

## 1. namespace in flatbuffers IDL ,  C++ / Python / Go

### 1.1 namespace in IDL

In the fbs IDL file, define namespace like:

```
namespace Mygame.Example;
```
### 1.2 namespace in C++



flatc --cpp will generate the following cpp header file:

```
./monster_generated.h
```

in monster_generated.h

```
#ifndef FLATBUFFERS_GENERATED_MONSTER_MYGAME_EXAMPLE_H_
#define FLATBUFFERS_GENERATED_MONSTER_MYGAME_EXAMPLE_H_

#include "flatbuffers/flatbuffers.h"

#include "weapons_generated.h"

namespace Mygame {
namespace Example {

struct Vec3;

struct Monster;
struct MonsterBuilder;
struct MonsterT;

...

```

namespace here

```
namespace Mygame {
namespace Example {
....
```

use in cpp

```
#include "monster_generated.h"  //  

using namespace MyGame::Example;
```



### 1.3 namespace in Go

For the go language, the namespace in fbs is called go package.

 flatc --go will generated the following go package file:

````.
./Mygame/Example/Monster.go
````
in Monster.go 
```
package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Monster struct {
	_tab flatbuffers.Table
}

// GetRootAsMonster shortcut to access root table
func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n+offset)
	return x
}

....

```



in go v1.11+ ( support go module) , for example , go project repo store in https://github.com/tsingson/flatbuffers-sample

it’s go module define by go compiler ( go tools ) 

```
go mod init github.com/tsingson/flatbuffers-sample
```

and generated a go.mod file  like this:

```
module github.com/tsingson/flatbuffers-sample

go 1.11

require (
	github.com/google/flatbuffers v1.12.0
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)
```



go  repo directory like this:

```
.
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
├── monster
│   ├── Mygame
│   │   └── Example
│   │       ├── Color.go
│   │       ├── Color.java
│   │       ├── Equipment.go
│   │       ├── Equipment.java
│   │       ├── Monster.go
│   │       ├── Monster.java
│   │       ├── SpaceShip.go
│   │       ├── SpaceShip.java
│   │       ├── Vec3.go
│   │       ├── Vec3.java
│   │       ├── Weapon.go
│   │       └── Weapon.java
│   ├── monster.fbs
```



use in go , example in main package ( main.go)

```
package main

import (
  sample "github.com/tsingson/flatbuffers-sample/monster/Mygame/Example"
)

func main() {

....

monster := sample.GetRootAsMonster(buf, 0)

....

}

```



### 1.4 namespace in Java



flatc --java  will generated the following java package file:

```
./Mygame/Example/Monster.java
```

in Monster.java

```
package Mygame.Example;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")
public final class Monster extends Table {
  public static void ValidateVersion() { Constants.FLATBUFFERS_1_12_0(); }
  public static Monster getRootAsMonster(ByteBuffer _bb) { return getRootAsMonster(_bb, new Monster()); }
  
  ....
  
```

namespace here

```
package Mygame.Example;
....
```

use in java:

```
import Mygame.Example;
```




### 1.5 namespace in Python











## 2. multiple namespace and go module support

## 2. --gen-onefile / --go-namespace parameter and namespace

