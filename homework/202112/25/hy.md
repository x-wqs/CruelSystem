# Distributed Data

reasons to distribute databases across multiple machines:

- Scalability

- Fault Tolerance / High Availability

- Latency

### Vertical Scaling

simply buy more powerful machine.

- shared-memory architecture

- shared-disk architecture

**limitations:**

- cost grows faster than linerly

- limited fault tolerance

### Horizontal Scaling

multiple nodes (whatever machine) coordinated by software using networks

**advantages:**

- be able to use machine which has the best price/performance ratio

- reduce latency by distribute across multiple regions

- fault tolerance

- take advantage of cloud deployments

**limitations:**

- additional complexity

- limiting the data model to use

### Two ways to distribute data across nodes

- Replication

- Partitioning
