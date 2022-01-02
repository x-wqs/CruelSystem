
## Limitation of quorum consistency
Quorum consistency makes sure you always get updated result because the set of nodes you write to and read from must overlap.

However, even with w + r > n, there can be edge cases where stale values are returned.
- sloppy quorum
- two writes occur concurrently, there's no clear order which leads to a merge of the values. e.g. based on timestamps, writes can be lost due to clock skew.
- write happens concurrently with a read. such a write maybe reflected on part of the replicas.
- a write might fail on some replica.
- node with updated value fails and get restored to an older version.
- unlucky with timing.

Therefore, quorum consistency is not the silver-bullet. Do not take them as absolute guarantees. 
Dynamo-style dbs are generally optimized for use cases that can tolerate eventual consistency.

### Monitoring staleness
Even if your application can tolerate stale reads, you should be aware of the health of your replication.
- leader-based replication, replication lag can be feed into a monitoring system.
  - writes are applied to leader and followers in the same order.
  - each node has a position in the replication log. the diff in position is the lag.
- leader-less replication is difficult to monitor, no common practice yet.
  - but we should still somehow find a way to quantify the 'eventual' of eventual consistency.
  
