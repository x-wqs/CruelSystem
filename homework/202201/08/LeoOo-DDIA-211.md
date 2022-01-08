
## Rebalance partition (repartition)
- add more nodes
- size increased
- node fails.

Strategies

- avoid hash % N
  - rebalance everything when number of nodes changes.
- fixed number of partitions
  - create many more partitions than there are nodes, and assign several partition to each node.
  - each node will be responsible to multiple partitions.
    - add new node, the new node steal a couple partitions from each node.
    - remove, each existing node will take a couple partition from the removed node.
- Dynamic partitioning
  - key-range partition
  - need to find proper partition boundaries.
  - split and merge as the data get inserted/deleted. i.e. dynamic.
  - the total number of partitions adapts to the total data volume.
  - HBase, MongoDB
- Partitioning propoertionally to nodes
  - when add a node, the ndoe randomly choose a fixed number of existing partitino to split and take ownership of half of each partitin.

auto vs manual rebalancing. auto rebalancing can be unpredicatble and may sacrifice performance.


### Reuqest routing
more general problem is service discovery.
- allow clients to reach any node (load balancer)
  - if the requested node cannot handle, it will forward the request to proper node
- build a routing tier
  - act as a partition-awar LB
- require clients to know partition and ndoes info

how do the rounting service learn about change of partition
- challenging because it involves multiple nodes
- ZooKeeper
  - each service register itself in ZK.
  - subscribers of ZK will get notified.
- alter native: gossip protocol.
  - request any node then forward to proper ndoe.

## Summay
- key range partition
  - sort
- hash partition
- local vs global index
