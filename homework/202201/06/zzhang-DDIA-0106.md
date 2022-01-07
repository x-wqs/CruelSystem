# DDIA: pp 201-210

Partitioning of Key-value data

- Skewed: Some partitions have more data or queries than others

- Hot spot: the partition with disproportionately high load



Partitioning by Key Range

- Assign a continuous range of keys to each partition
- Partition boundaryies need to adapt to the data
- hot spot problem example: partition by days. Writing data to one day, then the partition holding the writing day become hot spot. => use other keys to partition first



Partitioning by Hash of Key

- Cassandra, MongoDB: MD5; Voldemort: Fowler-Noll-Vo function
- Good: distribute keys fairly among the partition. 
- Bad: key-range partition => bad range queries
  - Cassandra: *compound primary key* => use the hash of first part of key to determine the partition, and the left keys are sorted.
- *Consistent Hashing*: randomly chosen partition boundaries to avoid the need for central control or distributed consensus



Skewed Workloads and Relieving Hot Spots

- In the extreme case where all reads and writes are for the same key => still skewed.
- Simple solution: Add a random number to the beginning or end of the key, then the keys are distributed to different partitions. Write to different partitions with the same original key => read operation requires reading from all partitions.
  - Additional bookkeeping: which keys are being split



### Partitioning and Secondary Indexes

Secondary index: usually doesn't identify a record uniquely; used to search for occurrences of particular value: e.g., articles containing the word xxx.



Two main approach to partitioning a db with secondary indexes:

1. Document-based partitioning: *local index*
   - Each partition maintains its *own* secondary indexes, covering only the documents in that partition. 
   - Read records with secondary index key:value => query all partitions and combine the result. *scatter/gather*

2. Term-based partitioning: *global index*
   -  The term we’re looking for determines the partition of the index
   - More efficient reads: a client only needs to make a request to the partition containing the term that it wants
   - Slow writes: Write to a single document => multiple partitions of the index (every term in the document might be on a different partition).
     - Asynchronous



### Rebalancing Partitions

- Move load from one node in the cluster to another

Methods:

1. Use hash range
   - divide the possible hashes into ranges and assign each range to a partition: 0 <= hash(key) < b_0 to partition 0, b_0 <= hash(key) <= b_1 to partiton 1, etc
   - mode N problem: when N increases, the partition will change.
2. Fixed number of partitions
   - Create many more partitions than there are nodes, and assign several partitions to each node
   - Only entire small partitions are moved between nodes.





