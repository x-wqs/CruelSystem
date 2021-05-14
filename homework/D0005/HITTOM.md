## welcome
```
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

## basic
```
package main

import (
    "fmt"
    "math"
    "math/rand"
)

func add(x int, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

// return value can be named
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var var_package_level bool

func main() {
    fmt.Println(math.Pi)
    fmt.Println(rand.Intn(10))
    fmt.Println(math.Sqrt(10))
    fmt.Println(add(23, 45))
    
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    
    fmt.Println(split(17))
    
    var var_function_level int
    fmt.Println(var_function_level, var_package_level)
    
    var i, j = true, "no!"
    k := 3
    fmt.Println(i, j, k)
}
```

#### basic types
```
bool
string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
```
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```
```
Variables declared without an explicit initial value are given their zero value.
```
```
Numeric constants are high-precision values.
```
## loop
```
package main

import (
    "fmt"
    "math"
    "time"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	const eps = 1e-6	
	for delta := (z * z - x) / (2 * z); delta < -eps || delta > eps; delta = (z * z - x) / (2 * z) {
		z -= delta
		fmt.Println(z)
	}
	return z
}

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
    
    sum = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
    
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
    
    fmt.Println(Sqrt(2))
    
    t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	
	defer fmt.Println("world")
	fmt.Println("hello")
	
	for i := 0; i < 10; i++ {
	    defer fmt.Println(i)
	}
}
```
## more type
```
Unlike C, Go has no pointer arithmetic.
```
- slice1.go
```
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}
func main() {
    fmt.Println(Vertex{1, 2})
    
    primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	var sub_primes []int = primes[1:4]
	fmt.Println(sub_primes)
	
	// equals to a := [6]struct { i int b bool }{......}, then create s that references it
	s := []struct {
	    i int
	    b bool
	}{
	    {2, true},
	    {3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	//s[0:6] = s[:6] = s[0:] = s[:]
	
	// Slice the slice to give it zero length.
	// length 0, capacity 6
	sub_primes = primes[:0]
	// Extend its length.
	// length 4, capacity 6
	sub_primes = sub_primes[:4]
	// Drop its first two values.
	// length 2, capacity 4
	sub_primes = sub_primes[2:]
	
	// The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5) // len(a) = 5
	fmt.Println(a)
	// To specify a capacity, pass a third argument to make
	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	b = b[:cap(b)] // len(b) = cap(b) = 5
	b = b[1:] // len(b) = cap(b) = 4
}
```
- slice2.go
```
package main

import "fmt"

func main() {
	var s []int
	// len = 0, cap = 0, []
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	// len = 5, cap = 6 [0 1 2 3 4]
	// ??? cap = 6
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
- iterate slice

```
package main

import (
    "fmt"
)

func main() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}
```
- slice to paint
```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8
	for i := 0; i < dy; i++ {
		var slice_dx []uint8
		for j := 0; j < dx; j++ {
			slice_dx = append(slice_dx, uint8(j * i))
		}
		res = append(res, slice_dx)
	}
	return res;
}

func main() {
	pic.Show(Pic)
}
```
- map
```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	
	var m2 = map[string]Vertex{
    	"Bell Labs": Vertex{
    		40.68433, -74.39967,
    	},
    	"Google": Vertex{
    		37.42202, -122.08408,
    	},
    }
	fmt.Println(m2)
	
	// v = 0, ok = false
	v, ok := m2["tencent"]
}

```
- function values
```
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// 13
	fmt.Println(hypot(5, 12))

    // 5
	fmt.Println(compute(hypot))
	// 81
	fmt.Println(compute(math.Pow))
}
```
- closure
```
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
```
```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		res := x
		x, y = y, x + y
		return res
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // 0 1 3 6 10 ...
			neg(-2*i),
		)
	}
	
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
