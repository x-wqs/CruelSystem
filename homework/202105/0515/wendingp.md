# Go notes

## 包
每个 Go 程序都是由包构成的。按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。

导入: 用括号分组导入更好
```go
import (
	"fmt"
	"math"
)
```

导出名用大写:

`package.Name`的形式

## 函数
函数变量类型在变量名之后 连续两个或以上形参类型相同可以只写最后一个

```go
func add(x int, y int) int {
	return x + y
}
```

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

## 变量
若初始化值存在 则可省略类型
```go
var c, python int = 1, 2
var c, python = true, "no!"
```

**函数内**也可以用赋值语句代替var声明
```go
func main() {
    c, python := true, "no!"
}
```

## 基本类型

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

没有明确初始值的变量声明会被赋予它们的 零值。
    数值类型为 0，
    布尔类型为 false，
    字符串为 ""（空字符串）。

## 类型转换

必须显式转换

## const
常量用const 不能用:=赋值
```go
    const Pi = 3.14
```
数值常量: 一个值 由上下文确定类型(本身不确定类型)


