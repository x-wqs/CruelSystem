ddia p199-210 

chapter 6 partition

# intro
- replication vs partitioning: 
  - both can improve query throughtput & reduce latency
  - replication -> high availibility
  - partitioning -> handle large dataset
- each piece of data (record, user info, document, txn) belongs to a single partition
- every partition can be viewed as a small database instance
- why partition? __scalability__ -> handle more data than current system via shared-nothing arch
- who uses partition? NoSQL (OLTP), Hadoop-based data warehouses (OLAP)

# kv data partition
- key question: how to assign kv data to partitions? want balanced partition, no skew, no hot spot
- by __key range__
  - e.g. partition a dict, keeps keys sorted
  - range scan is easy
  - the key range need to be fixed
  - need to make sure each partition has almost the same load
  - sensor id -> increasing ts
- by __hash of key__
  - FNV, md5
  - consistent hashing for nodes changes
  - hard to perform range scan
- by __both key range and key's hash__
  - use 2 fields as PK
  - use 1st key's hash to determine first partition
  - use 2nd key range to determine second sub-partition
  - range scan on the 1st key is not fast
- avoid hot spots
  - celebrity activities: many people comment or view one single data
  - naturally skewed workload, hard to avoid it, can only alleiviate it in application layer
  - solution: for such hot data, add prefix/suffix to make multiple keys
  
# partition & 2nd indexes
- 2nd indexes: search for value, not key
- __partition by document__:
  - each partition maintains a 2nd idx (local idx)
  - use scatter/gather approach to 
- __partition by term__:
  - 1. value -> PKs, then use PK to find records
