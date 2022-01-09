Rebalancing partitions

- need to add or remove nodes in the cluster
- partitioned data set need to be rebalanced
- want to minimize data transfer between nodes

# how to rebalancing
how to partition data to avoid rebalancing issues
- mod by N: not ideal
- fixed partition cnt
    - e.g. 10 nodes, 1000 parts
    - __steal__ parts and assign them to new nodes
    - __release__ parts when node exits
    - use to use in heterogeneous hardware architecture
    - sometimes difficult to choose a suitable partition cnt
- dynamic partition
    - in key ranged-partitioned approach, fixed partition may cause imbalance as system grows
    - want to achieve dynamic rebalancing
    - B tree like: partitions can split and merge, according to their size
    - pre-splitting: when start up, a single partition is not enough, need to make sure basic throughout.
- partition proportionally to nodes
    - fixed num & dynamic partition only consider size of dataset, not how many nodes you have
    - assign each node a fixed num of partitions
    - newly added nodes split existing partitions and take them

# how to route requests
- service discovery like
- where is partition knowledge? in server, in lb tier, in clients, 3 choices
- how to make sure partition knowledge is correct & up-to-date? consensus algorithm
- key range-partitioned approach can use __Zookeeper__ to record partition strategy (mapping from key to partition). routing tier load strategy from zookeeper
- gossip can also be used to update partition information
