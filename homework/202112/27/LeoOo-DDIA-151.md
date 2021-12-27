
# Replication
Keep a copy of the same data on multiple machine that are connected via a network.
- geographically close to user
- increase availability
- increase read thruput.

In this chapter we assume each node can hold a copy of the *entire* dataset.

## Leaders and followers
Each node that stores a copy is called a replica. Every write to the database needs to be processed by every replica.

Leader-based replication (aka master-slace, active-passive)
- one of the replica work as the leader, write requests are first sent to leader, who writes the new data to local storage.
- others are followers. whenever leader writes, it also sends data change to all of its followers as part of a replication log or change stream. followers then apply writes in the same order as processed by the leader.
- client reads from any of the follower or the leader but only write to the leader.

## Synchronouse vs Asynchronous replication
- Sync: the leeader send the write 'data change' message to follower, waits until followers confirm that it received the write before report success to user.
  - followers is guaranteed to have an up-to-date copy of the data, that is consistent with the leader.
  - but leader must lock all writes and wait until the sync replica is available again.
    - if sync follower does not respond due to crash?
    - impractical for all followers to be sync.
    - enable sync replication usually means *one* of the follower is sync. others are async.
    - at least two nodes have the up-to-date data.
      - **Semi-synchronouse**
- Async: the leader does not wait for a response from the follower.
  - Completely async risk to suffer data loss.
    - if non of the followers finish the update before the leader crashes. the write is not durable.
    - but the leader can continue process writes even if followers fall behind.
