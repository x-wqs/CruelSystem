# Replication

reasons to replicate:

- reduce latency
- increase availability
- increase read throughput

## Leaders and Followers

when writing, write to leader, and leader update followers accordingly

when reading, both are ok

**Replication Strategy:**

- synchronous
- asynchronous
- semi-synchronous

there's a trade-off between efficiency and consistency

In reality, asyn is more widely used

**Setting Up New Followers**

- take a snapshot of the leader
- copy the snapshot to the new follower
- the follower requests all changes since the snapshot
- finally the follower catch up

**Handling Node Outage**

- Follower failure: Catch-up
  - from its log it can get all the updates from the leader and then catch up
- Leader failure: Failover
  - automatic failover process
    - determine that the leader has failed
    - choose a new leader
    - reconfigure the system