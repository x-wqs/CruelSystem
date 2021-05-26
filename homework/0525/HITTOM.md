## source 
> 6.824 Lecture 4: Primary-Backup Replication

> https://www.youtube.com/watch?v=gXiDmq1zDq4

## plan
> failures

> challenge

> 2 approaches: state transfer, replicated state machine

> case study: VM FT system

## failures:
> solve fail-stop failures

> not solve logic bugs, configuration error, malicious error

## challenge
> has primary failed?(split-brain system)

> keep primary and backup in sync(apply changes in order, avoid non-determinism

> failover

## Two approaches
### state transfer
> send state

### replicate state machine 
> send operations(less cost)

## level of operations to replicate
### application-level operations
> gfs append write

### machine level
> or called instruction level

> today's discuss

> make replication transparent to application


## VMFT
> exploit virtualization

> transparent replication

> appears to clients that server is a single machine

> VMWare product

## Overview
> virtual machine monitor

> flag in storage server arbitrites who is primary after failover, set-and-get operation

## students' questions
### what happen if logging channel broke down
> system down, because nothing can be done

### explain set flag again
> storage server has two roles: basic storage and flag

## goal: behave like a single machine

## divergence source
> input packets

> time interrupts

> multicore --> not allow

## non-deterministic
> only non-determinsitic operations convert by logging channel

## output rule
> before primary response, make sure recieve backup's ack

## my questionss
> Q1: is it really only non-determinsitic operations converted by logging channel? how aboud determinsitic operations? compressed binay packet?

> A1: 

> Q2: how is the flag set? when there is one primary, the flag is id? when primary fail, the flag is 0? when another primary want to set flag when getting non-zero flag, it means already has a primary?

> A2:

> Q3: why are input packets non-determined?

> A3: 

## terms
> malicious 恶意的

> failover 故障转移

> emulate 模拟
