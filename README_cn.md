# 清晰胜过聪明: 改进 flatbuffers-go




## 0. 起因

使用 flatbuffers 已经有相当长的一段时间了. 

在几个商用项目中, flatbuffers 也因快速的反序列化而带来性能上的不少提升. 

flatbuffers 尤其适合传输小块数据, 一次序列化, 多个地方进行反序列化.

但 go 的 flatbuffers 有一些小遗憾:

1. go flatbuffers 功能支持, 滞后于 c++ 版, 代码库也很久没有更新了.
   相比 c ++ ,  go 版本缺少一些功能. 如 vector of unions , 在 unions 中包含 struct / strings . ( 注: go 版本的 flatbuffers 在 unions 中只能包含 table )
2. 缺少 verifier 验证器 ( 这是我需要的)
3. go flatbuffers 序列化的速度, 慢于 gogo protobuf.  flatbuffers 序列化消耗的时间, 大约是 gogo protobuf 的两倍.
4. go flatbuffers 不支持 go module. 尤其是自动生成的 go 代码存在相互引用时的 import 并不友好.
5. go flatbuffers 的序列化代码不太优雅, 不太符合 go 的习惯风格



在这样情况下, 我起了改进 go flatbuffers 的念头.

flatbuffers 的编译器, 是 c++ 写的. 我已经很多年没有用过 c++ 开发了. 对我来说, 这可能是一次有趣的探险历程.


## 1. 我对 go flatbuffers 的折腾

刚开始, 我写一个 flatbuffers verifier , 本地验证通过后, 我向 google flatbuffers 发了一个 PR. 结果被建议我重读一下 flatbuffers 的设计规范文档. 嗯哼, 这就开始有趣了.

在接下来的两周左右, 我边读 flatbuffers 的关键规范文档 ( 见附录参考列表) , 边写了一个全新的序列化生成器 ( flatbuffers builder ) .

 我拆分了flatbuffers 的 memory block , 采用 goroutine 并发处理各个独立的 memory block 转化为二进制序列数据, 最后进行合并/排序/优化.  当这个手写序列化器看起来可以工作时, 我发现, 需要把这些手写代码嵌入 flatbuffers 编译器中, 支持自动代码生成, 我遇到了一个小难题.  我几乎忘记如何写 C++ 了. 



为此, 我重读了 **Effective C++** 这样的几本册子, 随书写几行代码跑跑.  一周之后, 重新熟悉 C++ , 意外收获是对 go 的内存管理有了进一步的认识.

如何让 go flatbuffers 序列化更快, 我还在尝试中. 

而熟悉了 C++ 后, 我先让 go flatbuffers API 变得清晰简单, 易用一些.

## 2. 移植 C++ 有用功能, 支持 vector of unions.    

union 是 flatbuffers 中很有趣也很有用的一个功能, 当然, struct 也很有用.
go flatbuffers 中, union 只支持 table , 并且不支持 union array  ( 被称为 vector of unions ) , 先加上这个

IDL
```
union Character {
  MuLan: Attacker,  // table, 相当于 protobuf 中的 message
  Rapunzel,         // struct , 与 c++ 的 struct 相当
  Belle: BookReader,
  BookFan: BookReader,
  Other: string,   // string 
  Unused: string
}

table Movie {
  main_character: Character;      // 单一 union 字段
  characters: [Character];            // vector of unions 
}
```

##  3. 支持 go module via Attribute ( 在 IDL 定义中 ). 

 每一个 fbs IDL 定义文件都支持各自的 module , 格式像这样:  "go_module:github.com/tsingson/flatbuffers-sample/go-example/";

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


生成的 go 代码 
```
package Example

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"

	weapons "github.com/tsingson/flatbuffers-sample/samplesNew/weapons"    /// 嗯哼!  
)

type Equipment byte

..........

```

## 4. 增加一些清晰易用的 API /生成代码.

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

---------

或许, 稍后更多, 让  Go flatbuffers ...... 更好用.


## 5.   happy hacking....... 折腾继续中

 本文持续有更新...........


祝安康愉快!



tsingson ( 三明智 ) 于深圳南山. 小罗号口琴音乐中心
2020/04/09






