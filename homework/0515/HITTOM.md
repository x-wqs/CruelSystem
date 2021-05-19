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
- error
```
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```
- reader
```
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```
## concurrency
```
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
	
	ch_buffer := make(chan int, 2)
	ch_buffer <- 1
	ch_buffer <- 2
	fmt.Println(<-ch_buffer)
	fmt.Println(<-ch_buffer)
}
```
```
Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic
```
- select
```
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

``` 
- Exercise: equipment binary tree
```
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if nil == t {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		c1 := <-ch1
		c2 := <-ch2
		if c1 != c2 {
			return false;
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
```
# TODO
- concurrency 10 Exercise: Web Crawler
```

```
- methods 20 Exercise: Errors
```

```
- methods 22 Exercise: Readers
```

```
- methods 23 Exercise: rot13Reader
```

```
- methods 25 Exercise: Images
```

```

