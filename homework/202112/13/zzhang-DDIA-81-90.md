# DDIA: pp 81-90



### B-Trees

- *branching factor*: the number of references to child pages in one page of the B-tree

  - depends on the amount of space required to stroed the page references and the range boundaries.
  - Typically, several hundred.

- B-tree with n keys always has a depth of $O(log\ n)$

  

To make it more reliable => Use a *write-ahead log (WAL, or redo log)*. Every modification must be written to this log before applit to the page of the tree.



Concurrency control => Use *latches* (lightweight locks).



B-tree optimization

- reliable and concurrency control requirements => *copy-on-write* scheme => A modified page is written to a different location, then a new version of the parent pages in the tree is created, pointing at the new location.
- higher branching factor => abbrev key inside tree
- hard to maintain sequential order of leaf pages on disk => no solution for b-tree. use LSM-trees.
- Better range search => Use references on each leaf page, pointing to sibling pages

- Reduce disk seeks => Variants:  *fractal trees* 



B-tree vs LSM-tree

- B-tree: faster for reads; LSM-tree: faster for wires

- B-tree: fragmentation problem

- LSM-tree:

  -  *write amplification* problem: several compaction processes cause several rewrite

  - Reads/writes might have to wait for compaction process
  - Disk's write bandwidth preempt: Memtable flushing to disk vs. compaction.
  - If compaction cannot keep up with the rate of incoming writes => easily fill the disk, slow read (more checks)



### Other Indexing Structures

- Secondary indexes
  - Might not unique
  - Solution: 1. Make each value in the index a list of matching row identifiers; 2. Make each key unique key appending a row identifier to it.
- *heap file*
  - key-value pair: the value is the reference to the row stored elsewhere. 
  - Advantage: update efficiently <= but if the data gets large, might require pointing to a new place with updating all key-value pairs.
- *clustered index*
  - in MySQL’s InnoDB storage engine
  - store the indexed row directly within an index ???
  - primary key is a clustered index, and secondary indexes refer to the primary key ?
- *covering index* 
  - compromise between a clustered index (storing all row data within the index) and a nonclustered index (storing only references to the data within the index) 
  - stores *some* of a table’s columns within the index

- Multi-column indexes:
  -  *concatenated index*
    - combines several fields into one key by appending one column to another 
  - friendly for geospatial data, color, etc
    - e.g., `SELECT * FROM restaurants WHERE latitude > xxx AND latitude < xxx AND longitude > ... AND longitude < ...;`
- Full-text search and fuzzy indexes
  - *Levenshtein automation*: supports efficient search for words within a given edit distance [Lucene].
  - Sparse collection of some of the keys [LevelDB].



In-memory databases

- Memcached: intended for caching use only, acceptable for data lost
- Others (care about data lost): batery-powered RAM; write a log of changes to disk; write snapshots to disk; replicate in-mem state to other machine, ...
- Reduce overhead: no need for encoding the in-mem data structure in a form that can be written to disks.
- Support other data structures: e.g., priority queues and sets
- *anti-caching*: evict least recently used data



## Transaction processing

- needn't necessarily have ACID properties
- make low-latency reads and writes -- as opposed to batch processing



Different access pattern: 

- online transaction processing (OLTP)

  - Interactive operations. insert, update, ...

- Data analysis

  - scan huge records, read a few columns, cal aggregate statistics.

  













