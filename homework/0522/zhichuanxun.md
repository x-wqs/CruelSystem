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