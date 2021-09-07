#### 定义变量

```
message := "hello jingwen"
```



#### 创建数组需要声明长度

```
var b [6]int
var matrix [2][3]int
```



#### 切片存储

类似于python的list，可以不用声明长度

用make来创建

用append增加容量



#### map

类似于dict

用make来声明



#### 类型转换

类似于c#和java里的cast

并不是所有类型都可以转换为所有类型的



### 流程控制

用花括号但是不需要把条件写在小括号里

if

switch case

for

好像没有while循环



#### 指针

go支持指针，目前还不知道是干什么的



#### 函数

可以定义输入输出类型

main是主程序入口



#### 结构体 struct

有点像class，还不确定是用来做什么的

可能就是用来代替class的



#### interface接口

类似于java里面的接口概念



#### 部分有用的内置包

fmt：格式化的I/O功能，标准输出

json：解码编码json格式的数据

