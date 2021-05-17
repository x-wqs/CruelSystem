# a tour to go

# install go

安装地址：https://golang.org/doc/install
windows 下直接安装

# How to Write Go Code

## 模块的建立与使用

在 example.com/user 下运行以下shell

    $ mkdir hello # Alternatively, clone it if it already exists in version control.
    $ cd hello
    $ go mod init example.com/user/hello
    go: creating new go.mod: module example.com/user/hello
    $ cat go.mod
    module example.com/user/hello

    go 1.16
    $

在hello下新增 hello.go 文件

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

## go 的运行

### 模块的安装
```shell
go install example.com/user/hello
```
这个安装命令会将程序安装到系统环境中的GOBIN，
可以通过
```shell
go env -w GOBIN=/somewhere/else/bin
```
修改go的运行目录，只要GOBIN直接在环境变量中，编译的文件就能直接运行
### go 的运行
不过实际使用中，如果不需要模块特性的话，也能直接运行
```shell
go run main.go
```
上面这个命令会编译并运行

```shell
go build main.go
```

根据上面这两个

根据这两个命令可以配置vim 文件，用于写比较小的代码，vim 文件配置放在最后

## go 从网络上安装包
命令是 go get, 比如
```sh
go get github.com/google/go-cmp/cmp
```

```sh
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量，以下三选一

# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct
# 原文作者：Summer
# 转自链接：https://learnku.com/go/wikis/38122

```
## 需要在意的写法
### 省略类型
```go
package main
import "fmt"
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
```go
package main
import "fmt"
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
### 可以有多个返回值
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```
### Named return values

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
## Variables
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.
```go 
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```


### Variables with initializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

​
```go
package main
import "fmt"
var i, j int = 1, 2
func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
```

## Short variable declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
​


```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

= 与 := 的区别

    = 是赋值， := 是声明变量并赋值。

### Basic types
Go's basic types are

* bool
* string
* int  int8  int16  int32  int64
* uint uint8 uint16 uint32 uint64 uintptr
* byte // alias for uint8
* rune // alias for int32
     // represents a Unicode code point
* float32 float64
* complex64 complex128

没有初始化的数据默认为0

### Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this blog post.

```go
package main
​
import "fmt"
​
func main() {
    fmt.Println("counting")
​
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
​
    fmt.Println("done")
}
```




# vim 配置文件

### 分流配置
```sh
autocmd BufReadPre *.go exec ":source $VIMRUNTIME/VimSource/go_src.vim"
autocmd BufReadPre  *.go exec " !type NUL >> %<.goin "
autocmd BufReadPre  *.go exec " !type NUL >> %<.goout "
```


###  语法配置文件
```sh
"The TitleSet Part**************************

map <F2> : call SetTitle()<CR>
func! SetTitle()
	if &filetype=='go'
		call SetCppTitle()
	endif
endfunc

func! SetCppTitle()
	let l = 0 
	let l = l + 1 | call setline(l,"/*****************************************************************") 
	let l = l + 1 | call setline(l,"    > File Name: ".expand("%"))	 
	let l = l + 1 | call setline(l,"    > Author: Sevenger")
	let l = l + 1 | call setline(l,"    > Mail: Sevenger@gmail.com")								  
	let l = l + 1 | call setline(l,"    > Created Time: ".strftime("%c"))
	let l = l + 1 | call setline(l,"*****************************************************************/") 
	let l = l + 1 | call setline(l,"package main")
	let l = l + 1 | call setline(l,"")
	let l = l + 1 | call setline(l,"import \"fmt\"")
	let l = l + 1 | call setline(l,"")
	let l = l + 1 | call setline(l,"func main(){")
	let l = l + 1 | call setline(l,"\tfmt.Println(\"Hello, world.\")")  
	let l = l + 1 | call setline(l,"\t") | call cursor(l,0)
	let l = l + 1 | call setline(l,"}")
	let l = l + 1 | call setline(l,"")  
	let l = l + 1 | call setline(l,"")  
	let l = l + 1 | call setline(l,"")  
	let l = l + 1 | call setline(l,"")  
	let l = l + 1 | call setline(l,"")  
	let l = l + 1 | call setline(l,"")  
endfunc

"The Compile&Run Part***********************
map <C-x> :call CompileRunGpp()<CR>
imap<C-x> <ESC>:call CompileRunGpp()<CR>	
vmap<C-x> <ESC>:call CompileRunGpp()<CR>
func! CompileRunGpp()
	if &filetype=='go'
		exec "!go run %"
	endif
endfunc	

"Only Compile 
map<C-k> :call Compile()<CR>
imap<C-k> <ESC>:call Compile()<CR>
func! Compile()
	exec "w"
	if &filetype=='go'
		let startTime = reltime()
		silent exec " !go build %"
		let endTime = reltime()
		let result = reltimestr(reltime(startTime, endTime))
		echo result
	endif
endfunc
"Only Run
map<C-j> :call RunResult()<CR>
imap<C-j> <ESC>:call RunResult()<CR>
func! RunResult()
	if &filetype=='go'
		let startTime = reltime()
		silent exec "!%<.exe < %<.goin > %<.goout "
		let endTime = reltime()
		let result = reltimestr(reltime(startTime, endTime))
		let str = "Time: ".result."s."
		silent exec "!echo \t".str .">> %<.goout"
		" echo 'Time: '.string(result)
	endif
endfunc



"The IO Window***************************
nmap <C-n> :call OpenIOWIn()<CR>
func! OpenIOWIn()
	exec "vsplit  %<.goin "
	exec "wincmd L"
	exec "split   %<.goout "
	exec "resize 20"
	exec "vertical resize 40"
	exec "wincmd h"
endfunc

map <C-u> :call AddAnnot()<CR>
imap <C-u> <Esc> <C-u>i
func! AddAnnot()
	let str = getline(line("."))
	if -1 != match(str, '^\s*//')
		exec "norm ^xxx"
	else 
		exec "norm I// \<Esc>"
	endif
endfunc


"Make brace pair
vnoremap { dk$a{<ESC>o}<ESC>kpk$v%==k$
vnoremap [ di[<ESC>pi]<ESC>
vnoremap ( di(<ESC>pi)<ESC>
vnoremap < di<<ESC>pi><ESC>
"vnoremap " di"<ESC>pi"<ESC>
vnoremap ' di'<ESC>pi'<ESC>

```
