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