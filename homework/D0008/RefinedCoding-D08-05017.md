# RefinedCoding-D08-0517
- Lecture 2: RPC and Threads
- https://www.youtube.com/watch?v=oZR76REwSyA

## Advantage of Go
- threads/RPC/GC
- type safe

## Go Routine & Channel
- 内建的race检测器，可以检查出死锁？
- To use it, add the -race flag to the go command: go run -race mysrc.go
- 冲突检测标记位？
- 轻松检测Golang并发的数据竞争
- https://segmentfault.com/a/1190000017028859

## Threads
- alternative, event driven
- eliminates thread costs
- race condition
- lock, sync.Mutex
- chanels vs sync.Cond & WaitGroup
- detach dead lock, -race ?

## Crawler
- 1242. Web Crawler Multithreaded
- https://leetcode.com/problems/web-crawler-multithreaded/
- serial DFS
- concurrent BFS
- ConcurrentMutex crawler
- ConcurrentChannel crawler
- 本门课实验，我推荐用共享加锁来维护状态
- waiting/notification： 
- channels, sync.Cond, time.Sleep()

## IPC & Concurrency

## RPC
- 至多一次的简单形式，a simple form of "at-most-once"
- never resend ?
- Go语言从不重发？
- Better RPC behavior: "at most once"
- Go语言如何实现至多一次

## RPC Example - KV
- func (server *Server) Register(rcvr interface{}) error {
- rpcs.Register(kv)
- err := client.Call("KV.Put", &args, &reply)
- func (kv *KV) Put(args *PutArgs, reply *PutReply) error {

## Idempotent Request
- init connection faild
- server broken at middle
- client didn't see response
- check UUID in request
- repeat retry
- 三种错误，两端解决
- Designing robust and predictable APIs with idempotency
- https://stripe.com/blog/idempotency
- Avoiding Double Payments in a Distributed Payments System
- https://medium.com/airbnb-engineering/avoiding-double-payments-in-a-distributed-payments-system-2981f6b070bb

## Questions
- C++和Go可以传入参数引用，来修改参数的值，相当于返回多个参数
- err := client.Call("KV.Put", &args, &reply)
- 能不能问一下Java里面怎么才能修改参数的值，或者返回多个值啊？谢谢了！
- https://pdos.csail.mit.edu/6.824/notes/kv.go

## Teminology      
goroutines,例行程序  
interleaves,交错  
machinery,机械  
      
      
      
      
      
      
      
      
      
      
      