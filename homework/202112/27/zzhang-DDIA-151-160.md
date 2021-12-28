# DDIA: pp 151 - 160

## Ch 5: Replication

Replication:

- reduce latency, increase availability, increase read throughput



Assumption: 

- each machine can hold a copy of the entire dataset



Challenge:

- Handle changes to replicated data



Solution:

- single-leader, multi-leader, leaderless



### leader-based replication

alias: active/passive, master-slave replication

- One replica is *leader*. First write the new data to its local storage
- Others are *followers* (*read replicas*, *secondaries*, *hot standbys*). When leader writes, it sends the data change to all followers as part of *replication log* (*change stream*). Each follower takes the log and update its local copy.

- Read: can query either the leader or follower.



### Sync vs. Async Replication

Sync: the leader waits until follower has confirmed that it **received** the write before reporting success to the user

- Sync advantage: follower has up-to-date copy => high availablity
- Sync disadvantage: if follower doesn't respond, the write cannot be processed
  - In pratice: *semi-synchronous*: Only one of the followers is sync, others are async; If sync follower is slow, change another follower to sync mode.

Async: the leader sends the msg, but doesn't wait for a response from the follower.

- Async advantage: leader can continue processing writes, even if all followers have fallen behind.

- Async disadvantage: write is not guaranteed to be durable (write might be lost)



### New followers

1. take a consistent snapshot of the leader's db.
2. copy the snapshot to the new follower
3. connects to the leader and requests all the data changes since the snapshot
4. follower has *caught up*, and can continue to process data changes from the leader.



### Node Outages

Goal: keep the system as a whole running despite individual node failures. Keep the impact of a node outage small.

- Follower failure: follower always keeps a log of the data change on its local disk, so it can recover from its log and then connect to leader and requests the data changes during the disconnected time.
- Leader failure: Failover
  - One follower will be promoted to be the new ledaer: user send writes to the new  leader; other followers requests data changes from the new leader
  - If async mode, new leader might not have all up-to-date data. Usually, discard unreplicated writes => violate clients' durability expectations.
    - Dangerous if the database is related to other service.
  - *split brain* problem: two node both believe that they are the leader. Solution: some systems have a mechanism to shut down one node if two leaders are detected.
  - Appropriate *timeout* to declare the leader dead: trade-off



### Implementation of Replication Logs

1. Statement-based replication
   - Logs every write request (*statement*): e.g., SQL statements; 
   - Problems:
     - nondeterministic function (e.g., `now()`)
     - autoincreamenting column, or depend on the existing data in the db (e.g., `UPDATE ... WHERE ...`): If not execute in exactly the same order, may have different effect.
     - statements that have side effects (e.g., triggers, stored procedures, user-defined functions): might have different side-effects

2. Write-ahead log (WAL) shipping

   - Append-only sequence of bytes containing all writes to the database.
     - log-structured storage engine (SSTables, LSM-Trees): log is main place for storage. Compact and garbage-collect in the background
     - B-tree: write to write-ahead log first lest crash.

   - Disadvantage:
     - If storage format changes, the previous log might not be read correctly.

3. Logical (row-based) log replication

   - Decouple the replication log from the storage engine
   - A sequence of records describing writes to database tables at the granularity of a row.
   - Insert row: log contains new values of all columns
   - Delete row: log contains enough info to uniquely indentify that row.
   - Update row: log contains enough info to unqiuely indentify that row and the new value.
