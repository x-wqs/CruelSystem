# DDIA: pp 141- 150

## Distributed Data

Reason for applying distribution of databases.

- Scalability: spread the load
- Fault tolerance/high availability: get redundancy by using multiple machines
- Latency: Users around the world => set up servers worldwide.



Vertical scaling (scaling up)

- shared-memory architecture: not linear speedup, limited fault tolerance
- shared-disk architecture: contention, overhead of locking => bad



horizontal scaling (scaling out)

- many nodes with independet CPUs, RAM, disks
  - node: each machine or VM running the database software

- Disadvantage: complex 



How to distribute data

- Replication: provides redundancy
- Partition: 
  - *sharding*: different partitions canb e assigned to different nodes

