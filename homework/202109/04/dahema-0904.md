# GO

## 入门
- package main： 应用程序的入口

## go 的工作区
- 工作空间由环境变量「GOPATH」定义， export GOPATH=~/workspace

## 变量
- 定义： var a int， int 默认初始化为0
- 初始化：var a = 1， message := "hello world"， var b, c int = 2, 3

## 数据类型

### number: 
int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
var b int = 1
var d float32 = 1.222
var x complex128 = cmplx.Sqrt(-5 + 12i)

### string
string keyword
var c string = 'hello world'

### boolean 
bool keyword.
var a bool = true

## 数组
- 数组在声明中定义要指定长度，因此不能进行扩展
- 当数组的值在运行时不能进行更改。 数组也不提供获取子数组的能力
- var a [5]int， var multiD [2][3]int
- 切片：
  - 切片包含三个组件：容量，长度和指向底层数组的指针
  - numbers := make([]int,5,10)
  - append 或 copy 函数可以增加切片的容量

## map
-  Key-Value 类型的数据结构
-  m := make(map[string]int)
-  m["clearity"] = 2

## 类型转化
a := 1.1
b := int(a)

## 流程控制
- if else, same as c++ without bracket
- switch case, same as c++

## 循环
similar as C++
for i := 0; i < 10; i++ {
} 

## 指针
same as c++
var ap *int
a := 12
ap = &a

## 函数
func func_name(args...)(return_args...){
body
}
单个函数返回多个返回值，将返回值与逗号分隔开

## 结构体
type sutrcut_name sturct{
fileds...
}

type person struct {
  name string
  age int
  gender string
}
p := person{name: "Bob", age: 42, gender: "Male"}
person{"Bob", 42, "Male"}
pp = &person{name: "Bob", age: 42, gender: "Male"}
pp.name
指针访问属性是.operator

## 方法
方法是一个特殊类型的带有返回值的函数。返回值既可以是值，也可以是指针
func (p *person) setAge(age int) {
  p.age = age
}

## 接口
将不同类型的结构体组合在一起，对应的结构体必须拥有接口定义的函数

## 包
main 包， fmt包
Go 的包主要是用来进行大规模编程，并且可以将大型项目分成更小的部分

TODO
