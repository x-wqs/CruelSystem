# key-value store

## single server key-value store

- Store key-value pairs in a hash table
- optimizations
  - data compression
  - frequently used data -> memory, rest -> disk

## distributed key-value store

### CAP theorem
- impossible for a distributed system to simultaneously provide more than two of: consistency, availability, partition tolerace
- *Consistency* - all clients see the same data at the same time regardless of node connected
- *availability* - client requests data gets a response even if some nodes are down
- *partition tolerance* - system continues to operate despite network partitions

e.g. n1, n2, n3 three nodes

### Ideal situaion
- network partition never occurs
- data written to node n1 automatically replicated to n2 & n3, consistency & availability achieved

### real-world
- partitions cannot be avoided
- partition occurs, must choose between consistency & availability, n3 down
  - choose consistency - block all write operations to n1 & n2
  - choose availbility - keep accepting reads and writes

### Data partition
- split data into smaller partitions and store them in multiple servers
- challenges
  - distribute data evenly
  - minimize data movement when nodes are added or removed
- consistent hashing to partition data advantages
  - automatic scaling
  - heterogeneity - # of virtual nodes for a server is proportional to the server capacity

### data replication
- data replicated asynchronously over N servers
- N chosen
  - after key is mapped to a position on the hash ring
  - walk clockwise from that position
  - choose the first N servers on the ring to store data copies