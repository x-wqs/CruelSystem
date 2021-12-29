# DDIA: pp 161-170

### Problems with Replication Lag

Read-scaling arch: situable to workloads with mostly reads => create many followers and distribute read requests across followers

=> async replication => read req might be responded with out-dated data

- *eventual consistency* 
  - Response of the sam query on leader and a follower might be different. But if you **stop writing** to the db and wait a while, the follower will **eventually catch up**.

- *replication lag*
  - the delay between a write happening on the leader and being reflected on a follower.
  - lag large => serious inconsistencies problem



Three examples of problem when there is replication lag:

1. Reading your own writes

   - Example: User write to reader, and read from follower. User might not see his recent write.
   - *read-after-write consistency* (*read-your-writes consistency*): user can view the new own updates.
   - Implementation technique:
     - When read something that the user may have modified from the leader, otherwise, read it from a follower. Example: read user's own profile from the  leader and others' profiles from followers. => Not efficient if most things are potentially editable.
     - Track the time of the last update, for one minute after the last update, make all reads from the leader. Might track the replication lag on follower, and read from followers less than 1 min behind the leader.
     - Remember the timestamp (*logical timestamp*, e.g., log seq number or actual system clock) of most recent write.

   - *cross-device read-after-write consistency*. Example: If user writes on one device, he can view the new updates on another device
     - new issues: remembering the timestamp is difficult => the metadata need to be centralized; distributed cluster => no guarantee that connections from different devices will be routed to the same datacenter.

2. Monotonic Reads
   - Problem: see things *moving backward in time* if the user makes several reads from different replicas. It's because different replicas  have different lags.
   - *Monotonic reads*: if one user makes several reads in sequence, they will not see time go backward.
   - Implementation: read from the same replica, e.g., choose replica based on hash of ID.
3. Consistent Prefix Reads
   - Problem: Some data with dependency out-of-order
   - *consistent prefix reads*: if a sequence of writes happens in a certain order, then anyone reading those writes will see them appear in the same order.



*transactions*: they are a way for a database to provide stronger guarantees so that the application can be simpler. ??



## Multi-leader replication

- allow more than one node to accept writes

- appropriate for multi-data-center situation.



Multi-datacenter operation: have a leader in each datacenter

- inside datacenter: regular leader-follower replication
- across datacenter: each datacenter's leader replicates its change to the leaders in other datacenter

Advantage of multi-leader replication

- performance: not all writes go over internet, async write updates to other datacenter
- tolerance of datacenter outages: each datacenter operate independently
- tolerance of network problem: async replication usually tolerate network problems better.
