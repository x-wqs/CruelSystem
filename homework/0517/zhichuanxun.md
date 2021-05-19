# lecture 2

## why go?

go 对于进程，rpc有比较好的支持，同时是gc，类型安全以及编译型语言（可以更多地关注分布式本身而不是语言）

## why threads？

去体会一致性。（确实，感觉high level的分布式和low level的操作系统，组成原理啥的有很多共通之处）

## threads challenge:

- race condition
  - avoid sharing
  - use locks
- coordination
  - channels or condition variables
- dead lock

## go的并发编程

- channels(通信)
- locks + condition variables(共享内存)

go run --race . 可以看race的情况race detector

```go
package main

import (
	"math/rand"
	// "sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	// var mu sync.Mutex
	// cond := sync.NewCond(&mu)
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			// vote := requestVote()
			// mu.Lock()
			// defer mu.Unlock()
			// if vote {
			// 	count++
			// }
			// finished++
			// cond.Broadcast()
			ch <- requestVote()
		}()
	}

	for count < 5 && finished < 10 {
		v := <-ch
		if v {
			count += 1
		}
		finished += 1
	}

	// for {
	// 	mu.Lock()
	// 	if count >= 5 || finished == 10 {
	// 		break
	// 	}
	// 	mu.Unlock()
	// 	time.Sleep(50 * time.Microsecond)
	// }

	// for count < 5 && finished != 10 {
	// 	// wait
	// }
	// mu.Lock()
	// for count < 5 && finished != 10 {
	// 	cond.Wait()
	// }
	if count >= 5 {
		println("received 5+ votes")
	} else {
		println("lost")
	}
	//mu.Unlock()
}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	return rand.Int() % 2 == 0
}

```

[TODO]之前go的课还没补完，这边看得不是很懂，等把day5和6补完了再回头看一遍

## rpc 

remote procedure call

stub: marshal unmarshal

RPC semantics under failures

- at least once: 至少执行一次，rpc server 挂了的话，会一直retry直到有回复（感觉会直接阻塞然后雪崩啊）
- at most once: 需要处理重复 （tcp来确保？）
- exactly once: 需要维护状态