ddia ch3 storage and retrieval

- 为什么要学习数据库进行存储和查询的实现细节? 因为工程上需要理解一定程度的实现原理来进行技术选型
- 基础的kv数据库: append-only文件, 写入o(1), 查询O(n)
- 提高查询效率? 建立索引

# hashTable(Bitcask): 
append一个kv的时候更新hash, k映射到v的位置, 而非值, table记录在mem. 同时更新mem和log文件  
  - 适合写入非常频繁, k范围有限的case (ytb视频点赞的数目, 频繁增加)
  - log compaction: 定期清除掉k相同的值, compact成一个小文件
  - 细节: csv/binary, deleting via tombstone, snapshot for mem hashtable

# SSTable
常见于各大数据库实现中
- 写入kv数据时保持有序, sorted, 便于查找
- 多个kv组合作为一个segment, 通过merge sort来合并多个segment
- segment有时间, 因此如果一个k出现了多次, 取最新的作为v
- 检索segment: 通过in-mem sparse index, 二分找到对应的segment
- 存储时compress segment, 提高disk io效率, 提高查询速度
- 如何构建 segement? 内存中记录一个search Tree, 定期刷写到disk
- 写入流程:
  - kv写入in-mem search tree (RB Tree/AVL)
  - in-mem tree根据size或者timeout, 刷写到disk, 生成最新的segment
  - kv查找, 首先看in-mem, 然后看segment, 二分.
  - 定期做segment compact

# LSM-Tree
层级化的SSTable, 形成了树形结构
- LevelDB, RocksDB, Cassandra, HBase, BigTable

To-Be-Continued
