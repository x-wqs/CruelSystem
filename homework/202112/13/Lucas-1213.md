# CH3 Storing and Indexing
- SSTable: Sorted String Table
- 合并后，只保留每个键的最新值
- LSM Tree: Log Structured Merge Tree
- 分成三步，MemTable, Immutable MemTable, SSTable
- 先通过WAL，Write Ahead Log 预写日志写入磁盘，避免数据丢失
- 把最近的更新操作，保存到内存中的MemTable
- 然后MemTable达到一定大小的时候，变成Immutable MemTable
- 同时创建新的MemTable
- SSTable是LSM树在磁盘上的数据结构
- 不支持Delete/Update In Place 而是插入新的键值对，DelKey：abc
- LevelDB, RocksDB等键值数据库

## B树
- B树，或者B-树，m路平衡查找树，非根节点至少m/2个关键字，内部节点也存索引
- B树用于文件系统，MongoDB
- B+树，内部节点不存数据，索引数据只存在叶子节点，叶子节点也指向相邻叶子节点
- B+树用于 MySQL 索引 InnoDB，索引和数据放在一起，主键必需，但不能过长，主索引，聚簇索引
- MyISAM 也是用B+树，数据和索引分离，可以没有主键
