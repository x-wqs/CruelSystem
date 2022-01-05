# DDIA: pp 191-200

Tombstone:

- When one item is deleted from the databse, the system must leave a marker with an appropriate version number to indicate the item has been removed.



Version vectors: Collection of version numbers from all the replicas.

- Use a version number per replica and per key. Increment owne version number when writing.

- Allows the database to distinguish between overwrites and concurrent writes
- safe to read from one replica and subsequently write back to another replica.



### Summary:

Replication => high availability, disconnected operation (continue to work), low latency, scalability



Three approaches:

- Single-leader: no conflict resolution
- Multi-leader: robust in the presence of faulty nodes, network interruptions, and latency spikes
- Leaderless: robust in the presence of faulty nodes, network interruptions, and latency spikes



Leader fails => prmote an async updated follower to be the new leader



consistency models:

- Read-after-write consistency
- Monotonic reads
- Consistent prefix reads: logic dependency



## Partitioning (Sharding)

- Each piece of data (each record, row, or document) belongs to exactly one partition.

- Good for scalability.
  - Different partitions are put on different nodes in a shared-nothing cluster
  - query load can be distributed across many processors

- important topics: partition, rebalancing, route requests to the right partition.



Partition and replictaion can be combined.



