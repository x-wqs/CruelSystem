# 三篇文章了解 TiDB 技术内幕 - 谈调度
- https://pingcap.com/zh/blog/tidb-internal-3

## 调度的本质
**解决Replica的分布问题，而Replica的本质意图就只有3个。分别是提高容量、提高吞吐、容灾备份。我们这里关注的主要是后面两点**

### 基本需求点
- 副本数量不能多也不能少 （不太理解为什么完全不能多）
- 副本需要分布在不同的机器上
- 新加节点后，可以将其他节点上的副本迁移过来
- 节点下线后，需要将该节点的数据迁移走

### 优化点
- 维持整个集群的 Leader 分布均匀
- 维持每个节点的储存容量均匀
- 维持访问热点分布均匀
- 控制 Balance 的速度，避免影响在线服务
- 管理节点状态，包括手动上线/下线节点，以及自动下线失效节点

### 基本操作
基于 Raft 的 AddReplica、RemoveReplica、TransferLeader 去增删replica和做leader的transfer。

### 信息收集
调度决策需要基础的节点信息作为依据。我们任命节点PD，去监听Region和TiKV的状态。

### 调度策略
发现节点数量增多或减少，通过 Add ｜ Remove Replica 调整节点数量。
通过labels可以标记Replica的存储单元，如在同一个机架上，我们应该尽量让replica分散在不同的存储单元。
副本在Store分配均匀。
Leader在Store分配均匀。
访问热点在Store分配均匀。
各Store的存储空间大致相等。
控制调度速度。
支持手动下线。

### 实现
基于策略调用Leader的接口，等待Leader寻找合适的时机执行。

### TODO
对Store和Region的关系还不太了解。
