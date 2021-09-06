## 6.824 lecture 2
https://www.youtube.com/watch?v=oZR76REwSyA

## Why go

> good support for threads and rpc

> convenient gc

> type safe

> simple

> compiled

## thread of execution

## why threads

> I/O concurrency

> multicore parallelism

> convenience when need background function

## thread challenges

### race conditions

> example: both t1 and t2 operate n = n + 1

> avoid sharing: 1. use locks 2. race detect

### coordination

> channels or condition variables

### dead lock

## go and challenges

### two plans

> channels

> locks and condition variables

## crawlers
### goals

> I/O concurrency

> fetch each url once

> exploit parallelism

## RPC

### goal

RPC approxiamatly equel PC

## RPC semantics under failure

> at-least-once: request

> at-most-once: no duplication

> exactly-once

## terms

coordinate 协调

stub 树桩

marshal 打包

unmarshal 解包



