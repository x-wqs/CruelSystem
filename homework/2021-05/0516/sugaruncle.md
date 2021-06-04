# Lecture 2 notes

## Why go?

* good support for threads and RPC
* Garbage collector.
* Simple
* Compiled

## Thread of exection

Go has a runtime, a main thread.

And main thread can build more than one sub thread.

Each thread has its own PC/stack/register.

## Why threads？

express concurrency

* I/O concurrency
    * No time wasting on waiting for network response
* muti-core parallelism
* convenience
    * we can launch a thread for routine which 

## Thread challenges

- Race conditions 
    * Both thread read the same value at the same time, and increment it by one.
    * solution: avoid sharing, use locks
- Coordination
    * Channels，Condition variables，Wait Group
- Deadlock
    * t1 waits for t2 and t2 waits for t1

## go and challenges

go solve these problem with two plans.
* Channels
* locks + condition variables

## example 
* [mutex](https://cyent.github.io/golang/goroutine/sync_mutex/)

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, id int) {
    c.mux.Lock()
    fmt.Printf("%d. Inc lock.\n", id)
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key]++
    c.mux.Unlock()
    fmt.Printf("%d. Inc unlock.\n", id)
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    fmt.Println("Value lock.")
    // Lock so only one goroutine at a time can access the map c.v.
    defer fmt.Println("Value unlock.")
    defer c.mux.Unlock()
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 10; i++ {
        go c.Inc("somekey", i)
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
```
## Crawler

### goals
* I/O concurrency,
* Fetch urls once
* Exploit parallelism


