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

```