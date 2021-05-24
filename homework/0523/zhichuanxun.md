# FT VM

## Introduction

FT system：

primary/back up(备份机随时可以替代主机器)，backup server必须随时和主机保持链接，完全备份所有数据（cpu，disk。。。）

State-machine system：

将server模型化成确定状态机，然后保证initial state以及同样顺序的operation，但是对于物理主机，coordination很难实现。

但如果是vm的话，所有的operation也都可以被虚拟化

VM FT基于deterministic replay, 但增加了必要的protocol和functionality

## Basic FT Design

virtual lock-step(lock-step 查了下是一个同步技术[TODO])
所有的输入都传给primary并通过网络传到backup vm
backup vm会drop掉所有的output

### Deterministic Replay Implementation

inputs: incoming network packets, disk, keyboard
non-deterministic: virtual interrupts, clock cycle counter

1. capture inputs and non-determinism
2. apply 1
3. doesn't degrade the performance

### FT Protocol

send execution via logging channel, replay in real time

output rule: primary VM may not send an output until backup have received the log

Some FT is guaranteed by network protocol(TCP)

### Detecting and responding to failure

use UDP heartbeat

split-brain problem(the down may caused by network):
[solve]if the primary or backup want to live, it will execute an atomic test-and set operation on the shard storage

## Practical implementation of FT

### starting and restarting FT VM

VMware VMotion: clone a vm rather than migrating

### Managing the logging channel

log -> buffer -> flushed -> ack

primary vm will be blocked if the logs haven't been flushed

a mechanism(dynamically allocate the resource) to keep the log execution speed in a phase(动态分配资源的话，会不会对现有的数据有影响呢)

### Operation on FT VMs

the only operation that can be done independently is VMotion(ensure that no double VM on a same server)

when the backup VM is at the final point, it will requests primary temporarily quiesce all IOs

### Implementation Issues for Disk IOs

Problem1: disk operations are non-blocking and so can execute in parallel
DMA can also lead to non-determinism

Solution: detect IO race and force backup do the same thing

Problem2: disk operation race with a memory access by an application

Solution: bounce buffer(the same size as the memory accessed by disk op), disk read -> read from bounce buffer

Problem3: new primary not sure about whether disk IO has been finished
Solution: return error to indicate the IO's finished state

### Implementation Issues for Network IO

Some performance optimizations for VM networking may lead to non-determinism

FT disable the asynchronous network optimization which will downgrade the network performance: no context switch while sending and receiving the log entries

## Design alternatives

### shared vs non-shared disk

in the non-shared disk design, disk will be a part of state machine

non-shard disk can not deal with split-brain problem which will need majority algorithm

### executing disk reads on the backup VM

disk read => input

can reduce the traffic on the logging channel
need extra work to deal with failed disk operations

## Performance evaluation

1git/s

## Related work

epoch

Java FT VM

## Conclusion and future work

deterministic replay only for uni-processor
