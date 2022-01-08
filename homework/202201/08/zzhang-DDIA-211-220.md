# DDIA: pp 211-220

Rixed number of partitions:

- create many partitions and assign to servers proportional to the hardware etc. factors.

- the number of partitions is usually fixed for simple operation

  - data increase, still not split => need to choose the partition number high enough

  - need to choose *appropriate* number of partition.

- the size of each partition is proportional to the size of the dataset.

  

Dynamic partitioning (variable number of partitions)

- When a partition grows to exceed a configured size, it is split into two partitions 
- If lots of data is deleted and a partition shrinks below some threshold, it can be merged with an adjacent par‐ tition. 
- the number of partitions is proportional to the size of the dataset

- *pre-splitting*: at first, the dataset is small such that it might be put on single server => Not good since others are idle => allow an initial set of partitions to be configured on an empty database
  - key-range partitioning: pre-splitting requires the key distribution.



Partitioning proportionally to nodes

- make the number of partitions proportional to the number of nodes—in other words, to have a fixed number of partitions (e.g. 256) per node.
- keeps the size of each partition fairly stable
- New node: randomly choose a fixed number of existing partitions to split,  then takes ownership of one half of each of those split parti‐ tions while leaving the other half of each partition in place



Operations: Automatic or Manual Rebalancing

- fully automatic rebalancing: might overhead the network or harm the performance of other requests; can be dangerous with automatic failure detection.



Request routing

- one case of *service discovery*
- Common methods:
  - Clients connects to any node (via a round-robin load balancer): if the node owns the partitions, serve; otherwise, forward the request to the correct node, receive the reply, pass the reply to the lcient
  - Use a partition-aware load balancer, that determine the correct node.
  - User keep the partition-node assignment map, connect to the correct node directly.
- Problem: consensus about changes in the assignment of partitions to nodes.
  - Solution 1: *ZooKeeper*. Zookeeper maintains the authoritative mapping of partitions to nodes. 
    - Other actors subscribe to this information in ZooKeeper. 
    - Partition changes ownership, or a node is added or removed => ZooKeeper notifies the routing tier so that it can keep its routing information up to date.
  - Solution 2: (Cassandra, Riak) *gossip protocol*. Disseminate any changes in cluster state to any node. Use forward method.
  - Solution 3: (Couchbase) Not rebalance automatically. Use a routing tier which leans about routing changes from the cluster nodes.



Parallel Query Execution

- *massively parallel processing* (MPP) DB (for analytics): query contains several join, filtering, grouping, and aggregation operations.
- The query optimer breaks this complex query into multiple stages and partitions, many of which can be executed in parallel on different nodes.



### Summary

Goal: spread the data and query load evenly across multiple machines, avoiding hot spots



Two main approaches to partitioning:

- Key range partitioning: good: range search; bad: hot spot
- Hash partitioning: good: evenly; bad: destory the ordering, ineffective range search

- hybrid: one part of the key to identify the partition, another part for the sort order



Two methods for Partitioing and secondary indexes (search words etc)

- Document-partitioned indexes (local index): single partition write, scatter/gather read
- Term-partitioned index (glopbal index): all partition write, one partition read serve.



Routing Request

- Partition-aware -> parallel query execution
