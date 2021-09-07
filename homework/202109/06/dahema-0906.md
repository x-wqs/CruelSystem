# GO second part

## 包
- main 包， fmt包
- Go 的包主要是用来进行大规模编程，并且可以将大型项目分成更小的部分
- 包保存在环境变量 env 的 GOPATH 目录下
- 创建自定义包, 创建一个和包名一样的文件夹, 在文件中创建一个相关文件
- 安装包： 
- 包文档： godoc package_name Description
- 常见包： fmt， json， etc

## 错误处理
- 从函数返回自定义错误， error keyword
- Go 中内置的包或我们使用的外部的包都有一个错误处理机制
- 不能忽略错误

## panic
- 在程序执行期间突然遇到未经处理的异常
- panic 不是处理程序中异常的理想方式， 应使用error 对象
- 发生 panic 时，程序执行停止
- panic 之后要继续执行的程序就使用 defer

## defer
- 将使程序在程序执行结束时执行该行， 例如关闭文件

## 并发
-  Go routine： 需在函数前面添加关键字 Go， 简单非常轻量级
-  通道： 在两个 Go routine 之间传递数据， 必须指定通道接收的数据类型
-  单向通道： 通过通道接收数据但不发送数据， 或者反之
-  使用 select 为 Go routine 处理多个通道
-  缓冲通道： 在缓冲区满之前接受方不会收到任何消息

```
// Select example
package main

import (
 "fmt"
 "time"
)

func main() {
 c1 := make(chan string)
 c2 := make(chan string)
 go speed1(c1)
 go speed2(c2)
 fmt.Println("The first to arrive is:")
 select {
 case s1 := <-c1:
  fmt.Println(s1)
 case s2 := <-c2:
  fmt.Println(s2)
 }
}

func speed1(ch chan string) {
 time.Sleep(2 * time.Second)
 ch <- "speed 1"
}

func speed2(ch chan string) {
 time.Sleep(1 * time.Second)
 ch <- "speed 2"
}
```

## 总结
- 变量，数据类型
- 数组 切片和映射
- 函数
- 流程控制语句
- 指针
- 包
- 方法，结构体和接口
- 错误处理
- 并发 — Go 的线程和通道
