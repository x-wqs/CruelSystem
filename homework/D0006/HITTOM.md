## methods
```
no class in go, methods are functions bind to type
```
```
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type Myfloat float64
func (f Myfloat) Abs() float64 {
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    v.Scale(10)
    fmt.Println(v.Abs())
    
    var a Abser
    f := Myfloat(2.0)
    a = f
    a = &v
    fmt.Println(a.Abs())
}
```
- interface
```
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}
type Myfloat float64
func (f Myfloat) M() {
	fmt.Println(float64(f))
}
func main() {
	var i I = T{"hello"}
	i.M()
	
	// type assertion, no panic
	t, ok := i.(T)
	fmt.Println(t, ok)
	f, ok := i.(Myfloat)
	fmt.Println(f, ok)
}
```

