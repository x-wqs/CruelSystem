ddia p144-p151

previous read progress (p131-p141) https://github.com/refinedcoding/CruelSystem/blob/main/homework/202112/21/monsooooon.md#message-passing-dataflow-overview

part ii. distributed data

# intro
- previous part: data on a single machine, this part: multi-machines & data r/w
- why multi-machines
  - scalability: more CPU & mem to handle reqs
  - fault-tolerance/HA: redundancy
  - latency: make dc closer to user
  
# scaling methods
- scale up /  vertical scaling
  - shared-memory arch: Many CPUs with many RAM, connected together
  - problem: limited fault-tolerace via hot-swappable components (disk/mem/power), but still single location
- scale out / horizontal scaling
  - shared-nothing arch: each (virtual) machine running the db is called a node, which uses CPUs, RAM, and disks independently
  - benifits:
    - not speciail hardware
    - multi-region deployments (via cloud tech)
    - main focus of this book (distributed system)
  - powerful but complex

# replication vs partitioning
- replication: keeping a copy of the same data on several different nodes, potentially in different
locations
- partitioning: splitting a big database into smaller subsets called partitions so that different partitions
can be assigned to different nodes
- partitioning & replication can co-exist
