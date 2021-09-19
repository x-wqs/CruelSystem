## 0918

## 2 实现
- 使用抽象的目录来管理副本和局部性
- 一个Spanner的部署为Universe，全球大概几个Universe
- 每个Universe有一个master，一个放置驱动，若干区域集合
- 区域是物理隔离单位，有一个master和成百上千的Spanner服务器
- Universe master控制和显示所有区域的状态信息，可以交互调试
- 放置驱动每分钟检查需要一定的数据，完成区域迁移，保证负载均衡

## 2.1 Spanner 运行栈
- 多版本数据库（key: string, timestamp: int64) -> (value: string)
- Participant Leader -> Transaction Manager -> Lock Table -> Leader
- Replica -> Paxos -> Tablet (存在B树和WAL) -> Colossus / GFS
- Paxos状态机基于定时租约，只有为主的时候才可写，每个副本可读
- 副本之间用LockTable实现并发控制，包括两阶段锁，可事务性读取
- 副本之间使用事务管理器来支持分布式事务，Paxos Group使用两阶段提交
