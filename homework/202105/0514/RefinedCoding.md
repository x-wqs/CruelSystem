# Golang

## Install
- https://golang.org/doc/install
- https://golang.org/dl/go1.16.4.windows-amd64.msi
- go version
- go version go1.16 windows/amd64

## Tourial
- go get -u github.com/Go-zh/tour
- https://tour.go-zh.org/flowcontrol/8
- nums[1:4], exclusive last index, is a reference not copy
- array cap & slice length
- for i, v := range pow

## Slice
- https://tour.golang.org/moretypes/15
```go
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]

cap = 2 + 3 = 5
if (2 + 2 < 5) {
  new cap = 5 // end ?
} 
```
## Question
Q: int, uint uintptr is 32 bits on i686 and 64 bits on x86_64  
A: uintptr is used for variable address?  

Q: 	Big = 1 << 100  
	fmt.Println(needFloat(Big))  
1.2676506002282295e+29  
	fmt.Println(1 << 100)  
constant 1267650600228229401496703205376 overflows int  

## Solutions
Q: go get -u github.com/Go-zh/tour tour  
go: downloading github.com/Go-zh/tour v0.0.0-20190515134539-b61130663b4d  
go get: malformed module path "tour": missing dot in first path element  
A: go get -u github.com/Go-zh/tour  

Q: non-constant array bound dx  
A: can only use make to create dynamic array  

## Terminology
rune, 如尼字母（古代北欧文字）；神秘符号；芬兰民族史诗

## Examples 1
```go

import (
	"fmt"
)

// z² − x 是 z² 到它所要到达的值（即 x）的距离， 除以的 2z 为 z² 的导数，
// 我们通过 z² 的变化速度来改变 z 的调整量。 这种通用方法叫做牛顿法。
func Sqrt(x float64) float64 {
	var z = 1.0
	for i := 1; i <= 10; i ++ {
		z -= (z * z - x) / (2 * z)
	}
	return z;
}

func main() {
	fmt.Println(Sqrt(2))
}
```

## Example 2
```go

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)
	for i := 0; i < dy; i ++ {
		pic[i] = make([]uint8, dx)
	}
	for i := 0; i < dy && i < dx; i ++ {
		pic[i][i] = 255
	}
	return pic;
}

func main() {
	pic.Show(Pic)
}
```

## Example 3
```go

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, c + b
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```


