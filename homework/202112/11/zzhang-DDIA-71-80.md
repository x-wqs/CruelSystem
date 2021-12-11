# DDIA: pp 71-80

Two families of stiroage engines:

- log-structured: append-only data file.
- page-oriented such as B-trees.



### Indexes

#### Hash Indexes

- Key-value
  - value: a byte offset in the data file
  - Example: Bitcask (in Riak) : 
    - Hash map is kept compeletly in memory. Only one disk seek.
    - Suitted: larger number of writes per key.

- Log-based model

  - Too large => Use segments: 
  - Merge and compact with background thread
  - Format: CSV is not best. Use a binary format that first encodes the length of a string in bytes, followed by the raw string
  - Delete record: append a deletion record, *tombstone*. Delete when merge segments
  - Crash recovery: rebuild hash map through log file => Can store snapshot of hash map on disk to speed up the recovery process (Bitcask's method).

  - Partially written records: crash on the halfway through appending a record to the log => use a checksum to detect this and ignore it.
  - Concurrency control: one write, multiple read concurrently.

- Advantage of log-based model
  - Fast writing
  - Concurrency and crash recovery easy
  - Merge old segments to avoid the problem of fragmented data file.

- Limitation:
  - Hash table must fit in memory. Maintaining a hash map on disk is hard and expensive.
  - Hash collision require fiddly logic
  - Range queries are not efficient



#### SSTables and LSM-Trees

- Requirement: 
  - the sequence of key-value pairs is **sorted by key**
  - Each key only appears once within each merged segment file
- SSTable (Sorted String Table)
  - **Merging** is easy even if the files are bigger than the available memory. <= Use Mergesort-like algorithm
  - **Search** is easy. Do not need index of *all* keys in the memory, and can use a much sparser key hash-map (one key for every few kilobytes of segment file).
  - **Compression**: Group records within key ranges into a block and compress it before writting it to disk => save disk space and I/O bandwidth use.

- Storage engine working process

  1. Maintain a *memtable* in memory with an in-memory balanced tree data structure
  2. Write *SSTable* to disk when *memtable* reaches the threshold size.
  3. Compact *SSTable* with background thread

  * Search: First search from *memtable*, then check *SSTable* from recent to older.
  * Write: Always write into *memtable* first.

- Crash problem (in-memory records lost) => Append records to the log in the disk.

- This indexing structure => LSM-Tree: Log-Structured Merge-Tree
  - Used in LevelDB, RocksDB, Cassandra, HBase, Bigtable
- Optimization
  - Optimize the search: bloom filters
  - Optimize the SSTable compaction: 
    - Size-tiered compaction (HBase, Cassandra): newer and smaller SSTables are successively merged into older and larger SSTables
    - Leveled compaction (LevelDB, RocksDB, Cassandra): the key range is split up into smaller SSTables and older data is moved into separate “levels,” which allows the compaction to proceed more incrementally and use less disk space.

Other Similar work

- Lucene: indexing engine for full-text search
  - Use similar method for storing its term (word) dictionary: sorted files and merging process in the background.



#### B-Trees

- Most widely used
- break the database down into **fixed-size** *blocks* or *pages*, traditionally 4 KB in size
  - Log-structured indexes break the database down into **variable-size** *segments*

- More closely to the underlying hardware: disks are arranged in fixed-size blocks
- Root page. Use other page references in the tree.