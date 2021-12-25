# Part 2. Distributed Data
## Scaling to higher load
- Vertical scaling
  - Shared-memory architecture.
  - Not linear, twice the size not necessarily means enough processing power to handle twice the load.
  - Limited fault tolerance.
- Share-disk architecture
  - Several machines with independent CPU/RAM
  - Store data on an array of disks shared via network.
  - Contention and overhead of locking limits the scability of share-disk architecture.
- Share-nothing architecture
  - horizontal scaling or scaling out.
  - a node use its CPU, RAM, disk independently.
  - coordination between nodes is done at the software level, using a conventional network.

Replication and partition
- Replication keep a copy of the same data on several nodes, provide redundancy, improve availability.
- Partition split a big dataset into smaller subsets that are assigned to different nodes, aka, sharding.
Partition and replication can work together, e.g.: a database split into two partition and each partition have two replicas.

