### Methods

- A special type of function with a receiver. A receiver can be a value or a pointer

```golang
func (p *person) describe() {
  fmt.Printf("%v is %v years old.", p.name, p.age)
}
```

### Interfaces
- A collection of methods

### Packages
- main package is the entry point for the program execution
- installing a package `go get github.com/satori/go.uuid`

### documentation
- generate documentation for Description function inside package person `godoc person Description`
- see the documentation `godoc -http=":8080"`

### errors

```golang
return errors.New("custom error")
```

### panic

- unhandled, not ideal way to handle exceptions
- when panic occurs, program execution get's halted

### go routine
- go routines are the function that can run in parallel or concurrently with another function
- `go funcName()`

### channels
- We can pass data between two Go routines using channels
- receiver Channels wait until the sender sends data to the channel.
```golang
c := make(chan string)
```

```golang
package main

import "fmt"

func main() {
	c := make(chan string)
	go func() { c <- "hello"} ()
	msg := <-c
	fmt.Println(msg) //=> "hello"
}
```

- one way channel - only send or receive data but not both

```golang
package main

import (
 "fmt"
)

func main() {
 ch := make(chan string)
 
 go sc(ch)
 fmt.Println(<-ch)
}

func sc(ch chan<- string) {
 ch <- "hello"
}
```

- buffered channel - the messages sent to the channel will be blocked if the buffer is full
```golang
package main

import "fmt"

func main(){
  ch := make(chan string, 2)
  ch <- "hello"
  ch <- "world"
  ch <- "!" # extra message in buffer
  fmt.Println(<-ch)
}

# => fatal error: all goroutines are asleep - deadlock!
```
