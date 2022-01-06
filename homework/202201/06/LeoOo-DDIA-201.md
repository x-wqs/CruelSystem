# Partitioning

Replication is about having multiple copies of the same data on different nodes.
For very large datasets, or very high query throughput, we need to break the data up into partitions, also known as sharding.
- improve scalability
  - different partitions on different nodes
- disjoint

## Partition and Replication
- work together, each partition are stored on multiple nodes.
- each node can be the leader of some partition and the follower of another partition
- goal
  - spread the data and query load evenly across nodes.
  - skewed partition leads to hot nodes
  - to avoid skewed partitioning
    - randomly distribute the tuples.
- equal depth partition
  - by key range, but make sure each partition with roughly same amount of data.
  - used by bigtable, hbase, etc.
  - does not work with timestamp attribute, the node for the most recent time will be a hot node.
- by hash of key
  - hash function does not need to be cryptographically strong.
    - but they do need to generate the same key given the same input
      - java's object.hashCode generate different key in different process for the same input.
  - assign each partition a range of hashes.
  - efficient for point query, aka looking a particular key
  - but cannot handle range queries.
    - keys are scattered cross all partitions.
- compromise of the two partition strategy used by Cassandra
  - use a compound primary key consisting of several columns
    - first part of it is hashed to determine the partition.
    - other columns are also used as the concatenated index for sorting the data in SSTable.
    - e.g. uid, postid
      - partition by uid, all the posts by the same user will be on the same partition.

## Skewed workload and hot spot
- no perfect solution, application need to handle it.

