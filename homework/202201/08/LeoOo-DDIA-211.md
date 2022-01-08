
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

