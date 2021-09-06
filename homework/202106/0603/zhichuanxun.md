# Raft

## Abstract

raft is a consensus algorithm for managing a replicated log

## Introduction

Raft is similar with many existing consensus algorithms
novel features:

- strong leader
- leader election: use random timer to elect leader
- membership change

## Replicated state machines

state machines on a collection of servers compute identical copies of the same state

Replicated state machines are used to solve FT problems

Replicated state machines are typically implemented using a replicated log

The consensus algorithm is to keep the replicated log consistent

Properties:

- ensure safety under all non-Byzantine[TODO] conditions(network delays, partitions, packet loss, duplication and recording)
- majority of servers are operational and can communicate with each other
- consistency do not depend on timing
- a command can complete as soon as majority of the cluster has responded to a single round of rpc

## Paxos

- Paxos is exceptionally difficult to understand
- Paxos does not provide a good foundation for building practical implementations(implementation will be changed due to real-world situation, and the final system will be based on an unproven protocol)

## Designing for understandability

- problem decomposition: divide problems into separate pieces that could be solved, explained and understood relatively independently(leader election, log replication, safety and membership changes)
- simplify the state space by reducing the number of states to consider(logs are not allowed to have holes)

## The Raft consensus algorithm

Raft use a distinguished leader to manage all logs

consensus problem is decomposed into 3 questions:

- Leader election
- Log replication: leader must accept log entries and replicate them across the cluster
- Safety: the safety in state machine

## algorithms in consensus

TODO

- [ ] State
- [ ] Append Entries RPC
- [ ] RequestVote RPC
- [ ] Rules for servers

### basic

machine's role: leader, follower, candidate(in election term)

raft divide the time into different terms, there is at most one leader in each term

term plays a logical clock in Raft they allow servers to detect obsolete information

if candidate or leader discover that its term number is smaller, it will become follower

only two rpc is required:

- RequestVote(election)
- AppendEntries(heartbeat)
  
- Transfer snapshot

### Leader Election

Leader use heartbeat to establish its authority and prevent new election

Election happens when follower receive no heartbeat

Election:

1. followers increment term
2. transitions to candidate
3. RequestVote RPC

first-come-first-served basis, vote for a most one candidate

use term number for comparison

no candidate obtains majority: use randomized election timeouts to ensure this resolved quickly

### Log replication

leader issues AppendEntries RPCs in parallel to others.
leader will retry indefinitely until all followers eventually store all log entries

log is the item with log index, term and opt

ensures:

- If two entries in different logs have the same index and term, then they store the same command.(leader create at most one entry and entries never change their position)
- If two entries in different logs have the same index and term, then the logs are identical in all preceding entries.(AppendEntries RPC contains index and term that immediately precedes the new entries)

use strong leader to ensure consistency: forcing followers' log to duplicate its own

leader maintain a nextIndex for all server, when a new leader elected, it will change the nextIndex according to its own log entries

if AppendEntries RPC fails due to nextIndex, leader will decrement nextIndex until success

### Security

#### Election restriction

Leader's log must be up-to-date

- logs with later term
- logs is longer when they have the same term

#### Committing entries from previous terms

never committed by counting replicas

#### Safety argument

[TODO] review the Figure 9
