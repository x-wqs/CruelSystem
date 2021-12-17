## Sort order in column storage
An order in column storage can be helpful:
- range queries on certain column
- compression
But the column store need to be sorted an entire row at a time in order to match all columns that stores in different file. We cannot sort the columns independently.

### Multiple (different) sort orders
Store copies of the same data in different sort orders, each sort order can benefit a certain type of query.  (Vertica used it)
- Data needs to be replicated to multiple machines anyway, store the redundant data sorted in different ways does not necessarily introduce additional storage overhead.

Writes become more difficult. 
- In-place update like B-tree does not work with compressed columns.
- To insert a row, you most likely have to rewrite all column files.
- Use LSMT
  - write to memtable, 
  - if large enough, **merged with column files on disk and write to new files.**
    - not a trivial task. Vertica does it.
    - Query check recent writes in memtable as well as files on disk.
