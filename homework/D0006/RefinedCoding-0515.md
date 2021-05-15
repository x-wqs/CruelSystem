# Golang Part 2

## Day 06, 0514
- https://tour.go-zh.org/methods/15

## Method
- func (v *Vertex) Scale(f float64) {

## Function
- func Scale(v *Vertex, f float64) {
```go
# work
func (v *Vertex) Scale(f float64)
var v Vertex
v.Scale(10)
p := & v
p.Scale(10)

func Scale(v *Vertex, f float64)
Scale(&v, 10)

# no work
Scale(v, 10)
```

## Interface
- type I interface
- type T struct

## Printf
- https://golang.org/pkg/fmt/

## Type
switch v := i.(type) {
```go
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
```

# Errors
- and ? 	if err := run(); err != nil {
```go
// https://tour.go-zh.org/methods/20
type ErrNegativeSqrt struct {
	x float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.x)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, &ErrNegativeSqrt { x }
	}
	return math.Sqrt(x), nil
}
```

## IO
```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

/*
func (r MyReader) Read() []byte {
	return 'a'
}

func (r MyReader) Read([]byte) (int, error) {
	b[0] = 'A'
	return 1 nil
}

go: finding module for package golang.org/x/tour/reader
go: downloading golang.org/x/tour v0.0.0-20210512164546-a278aee398d5
go: found golang.org/x/tour/reader in golang.org/x/tour v0.0.0-20210512164546-a278aee398d5
./prog.go:17:11: syntax error: unexpected nil at end of statement
*/

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
```

## Notes
- 抓狂了，从来没有哪个语言学起来这么痛苦
- 这样可以么？照葫芦画瓢写了一个，还是不懂这个-2怎么来的



