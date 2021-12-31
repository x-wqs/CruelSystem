## Multi-Leader Replication

a big problem of Leader-based replication is that all the writes must go through the master

so there's an extension to allow more than one leader.

**Use Cases**

- Multi-datacenter operation
  - you can have a leader in each datacenter
- Clients with offline operation
- Collaborative editing

### Handling Write Conflicts

conflicts is the biggest problem of multi-leader replication

- conflict avoidance

- converging toward a consistent state
- custom conflict resolution logic

### Multi-Leader Replication Topologies

the communication paths along which writes are propogated

- Circular
  - mysql default
- Star
- All-to-All



