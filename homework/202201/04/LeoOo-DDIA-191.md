
Version vector
- a single version number to capture dependencies between operations.
  - not sufficient when there're multiple replicas accepting writes concurrently.
- need use a version number per replica as well as per key.
- each replica increments its own version number when processing awrite.
  - also keep track of the version numbers it has seen from other replicas.
  - vv are sent from db to client then sent ack to db when a value is written.
  - vv allows db to distinguish between overwrites and conccurent wirtes
    - **exactly how???**
- lamport clock is also an alternative to provide logical ordering of events.


## Summary
Replication provides:
- high availability
- disconnected operation
- improved latency
- scalability

Three main approaches:
- single-leader
  - no conflict resolution
- multi-leader
- leader-less
  - more robust but weak consistency guarantees

Sync or async

- Read-after-write consistency
- monotonic reads
- consistent prefix reads
  - causal ordering, e.g. order of dialogues
