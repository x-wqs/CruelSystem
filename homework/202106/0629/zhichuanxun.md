# lesson 9

lab3

- widely-used system
- high performance
  - asynchronous
  - consistency
- coordination service

## replicated state machine

raft is similar to zab
service on raft

## lab3

the time spend in put
1-round-trip

## zookeeper throughput

21k/s

- asynchronous bach operations
- read by any server

## read from many peers

back in time problem

## Linearizability

behaves like single machine

- a total order of operations
- order matches real time
- read operations return value last write

the easiest way is to put and get from leader

lab3:
get ops go-through raft

## Zookeeper

change correctness definition

- linearizable writes
- FIFO client order
  write n client order
  read observe last write form same client, prfix of the log => stable data
  no read from previous one

client -w->    <-zxid-

client -r+axid->    wait azxid

stale data