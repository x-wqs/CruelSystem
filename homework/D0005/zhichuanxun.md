# go lesson1

## basic

```go
package main

import (
  "fmt"
  "math/rand"
  "math"
)

func main() {
  fmt.Println("Welcome to the playground!")
  fmt.Println("The time is", time.Now())
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println(math.Pi)
}
```

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z	complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

## control

```go
package main

import (
	"fmt"
	"runtime"
)

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

	if x := -1; x < 0 {
		fmt.Println("<0")
	} else {
		fmt.Println(">=0")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}
```

## defer

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
  
  fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

## Pointer

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

## Structs

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	q  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 2
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Print(v)
	fmt.Println(v1, q, v2, v3)
}
```

## Arrays

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	// slice like references
	s[0] = 1000
	fmt.Println(primes)
	// slice 的第一个参数会影响cap
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// make(arr, len, cap)
	b := make([]int, 5)
	fmt.Println(b)
	b = append(b, 1)
	fmt.Println(b)
	b = append(b, primes[:]...)
	fmt.Println(b)

	for i,v := range b {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

## Maps

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m2 = map[string]Vertex {
	"Bell Labs": {1, 2},
	"Google": {2, 3},
}

func main() {
	fmt.Println(m)
	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["1"] = 1
	m3["2"] = 2
	delete(m3, "2")
	fmt.Println(m3)
	//v, ok := m3["1"]
	v, ok := m3["2"]
	fmt.Println("The value:", v, "Present?", ok)
}

```

## Function

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int
    return func() int {
        a, b = b, a + b
        if b == 0 {
            b = 1
        }
        return a
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

```
