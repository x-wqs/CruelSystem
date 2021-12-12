#Ch3 Data store and indexing
- 哈希索引
- SSTable in Memory, LSM Tree on Disk
- B Tree
- B Tree vs LSM Tree
- Value in Index
- Multiple Column Index
- Text Index
- 模糊索引
- Store All in Memory
- OLTP or OLAP
- Data Warehouse
- Star Model vs Snowflake
- Column Store

## 简单数据库
- 键值数据对
- 无索引，仅追加

## 哈希索引
- 仅在内存？
- Bitcask，Riak 默认引擎
- 压缩段，Compaction
- 只保留最近值
- 标记删除，快照恢复
- 严格追加，单一写入
- TODO-P78
