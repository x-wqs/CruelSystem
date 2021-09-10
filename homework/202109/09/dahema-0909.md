# TiDB 技术内幕 - 谈调度

## 需求分析
- 同一个region的不同replica分布在不同的节点，机房
- 如何增加节点，将其他节点的副本迁移过来，
- 如何处理节点下线，需要将该节点的数据迁移出去
- 如何调节和确定replica的个数，副本的数量不能多也不能少，多了影响性能，少了影响容错
- leader对于性能的影响，维持整个集群leader的均匀分布
- 维持每个节点的储存容量均匀
- 维持访问热点 hot spot 分布均匀
- 如何处理负载均衡，控制 Balance 的速度，避免影响在线服务

## 调度的基本操作
- 增加/删除一个副本
- Leader在一个 Raft Group 的不同 Replica 之间 transfer

## 信息收集
- 每个 TiKV 节点会定期向 PD 汇报节点的整体信息
  - 检测store是否存活
  - 是否有新加入的store
  - 携带store的状态信息，总磁盘容量，可用磁盘容量，承载的 Region 数量，数据写入速度
- 每个 Raft Group 的 Leader 会定期向 PD 汇报信息
  - leader位置
  - followers 位置
  - 掉线的Replica的个数
  - 数据写入/读取速度

## 调度的策略
- 保证replica的数量满足要求，通过add/remove replica实现
- 保证replica不在同一个位置
- replica/leader/hot spot在store之间均匀分配
  - TODO 读写都是通过leader 怎么保证hot spot 在store之间均匀分配？？？
- 各个 Store 的存储空间占用大致相等
- 控制调度速度，以避免过度消耗资源影响在线服务
- 支持手动下线节点

## 调度的实现
- PD 不断的通过Store和Leader的心跳包收集信息，获得整个集群的详细数据
- 根据集群的详细数据以及调度策略生成调度操作序列
- 每次收到Region Leader发来的心跳包时，PD都会检查是否有对这个Region待进行的操作，并将需要进行的操作返回给Region Leader，
- 在对应的Region Leader的心跳包中监测执行结果
- 上述操作只是给Region Leader的建议，并不保证一定能得到执行

## 总结
- 调度的需求讲述清晰
- 信息收集的过程和接口讲述清晰
- 具体的调度策略不够具体，TODO 进一步挖掘
