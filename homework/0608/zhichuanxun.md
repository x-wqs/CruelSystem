# Lesson5 Fault Tolerance raft 1

## Pattern

Single point of failure(Coordinator, Master, Storage server)
avoid brain-split problem

Idea: test-and-set server replication

## Network partition

majority rule

majority > n / 2

## Protocols

Paxos, View-stamped replication

Raft(2014)

term plays a logical clock in Raft they allow servers to detect obsolete information

if candidate or leader discover that its term number is smaller, it will become follower

only two rpc is required:

- RequestVote(election)
- AppendEntries(heartbeat)

## Lab

clients --put get --> KV(raft)

lab3: detect duplicate

## Overview

apply channel ->

appendEntry includes committed entry

## Why logs?

logs identical on all servers

retransmission
order
persistence
space tentative

## Log entry

leader issues AppendEntries RPCs in parallel to others.
leader will retry indefinitely until all followers eventually store all log entries

log is the item with log index, term and opt

contains:

- command
- leader's term
- log index

- elect leader
- ensure logs become identical

## Election

Leader's log must be up-to-date

- logs with later term
- logs is longer when they have the same term

happens when finding missing heartbeat

use term to determinate who is qualified.

challenge: split vote

when split vote happened, it will run into next election

and use random election timeout to avoid infinite split vote

- few heartbeats
- random value
- short enough that downtime is short

## log may diverge
