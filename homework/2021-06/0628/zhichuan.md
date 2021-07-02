# Zookeeper

## Introduction

Locks constitute a powerful coordination primitive

When designing API of ZooKeeper, move away from blocking primitives(locks)

wait-free is important for performance, zookeeper provide order guarantees for operations

linearizability, leader-based atomic broadcast protocol

zookeeper use watch mechanism to enable clients to cache data without managing the client directly

- Coordination kernel(wait free, relaxed consistency)
- Coordination recipes(build higher level coordination primitives, blocking, strongly consistent primitives)
- Experience

## The ZooKeeper service

client: user of the ZooKeeper service
server: process providing service
znode: in-memory data node
data tree: hierarchical namespace

### Service overview

clients access a set of data nodes(znodes)

znode:
regular: create and delete explicitly
ephemeral: system can remove automatically

datamodel:
full data r/w or kv with namespace

znodes map to abstractions of the client application

session: timeout, watch, persistence

### Client API

create, delete, exists, getData, setData, getChildren, sync

all methods have both synchronous and asynchronous

### ZooKeeper guarantees

Linearizable writes: all requests that update the state of ZooKeeper are serializable and respect precedence

FIFO client: request from a given client are executed in order

Scenario:

- As the new leader starts making changes, we do not want other processes to start using the configuration that is being changed;
- If the new leader dies before the configuration has been fully updated, we do not want the processes to use this partial configuration.

- if a majority of ZooKeeper servers are active and communicating the service will be available
- if the ZooKeeper service responds successfully to a change request, that change persists across any number of failures as long as a quorum of servers is eventually able to recover.

### Examples of primitives

Configuration Management, Rendezvous, Group Membership, Simple Locks, Simple Locks without Herd Effect, Read/Write Locks, Double Barrier

## ZooKeeper Applications

Fetching service, katta, Yahoo! Message Broker

## ZooKeeper Implementation

### Request Processor

request will be translated into transaction and calculate data, version and timestamp

### Atomic Broadcast

use zab to guarantee the order

### Replicated Database

no need to lock since state change are idempotent

### C/S Interactions

server will notify client and clean the watch when update happened
zxid

## Related work

a service that mitigate the problem of coordinating process in distributed applications