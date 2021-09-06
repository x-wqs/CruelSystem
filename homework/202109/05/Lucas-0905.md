# 三篇文章了解 TiDB 技术内幕 - 说存储
- https://pingcap.com/zh/blog/tidb-internal-1

# 数据库需求
- 支持跨数据中心的容灾
- 写入速度快
- 写入后方便读取
- 支持数据修改，和并发修改
- 支持修改多条数据记录的原子性

## 键值数据库
- 巨大的TreeMap
- 高性能，分布式，可靠的，巨大的映射
- 使用RocksDB作为单节点数据库的底层
- 使用Raft协调所有节点
- Raft主要功能：领导选举，成员变更，日志复制
- TODO: Raft 优化
- https://zhuanlan.zhihu.com/p/25735592

## 区域
- 为了水平扩展
- 按照键映射到多个区域
- 以区域为单位，数据分散到所有节点
- 每个节点分配的区域数差不多
- 以区域为单位进行Raft的日志复制和成员管理

## 多版本并发控制 MVCC
- 多客户端同时修改同一个键值对
- 在键后添加版本
- 获取数据版本，然后更新第一个大于这个版本的版本？

## 事务
- 采用Percolator模型
- https://www.usenix.org/legacy/event/osdi10/tech/full_papers/Peng.pdf
