# 分布式一致性算法汇总
- https://zhuanlan.zhihu.com/p/130332285

## 需求
- 多个节点解决单点故障
- 多个节点保持数据一致
- 分布式系统中，保持多节点的数据一致性

## 分类
- 强一致性，提交后即改变集群状态，如 Paxos，Raft，ZAB
- 弱一致性，最终一致性，如 DNS，Gossip 协议
- Google 的 Chubby 分布式锁采用 Paxos
- etcd 分布式键值数据库，采用 Raft
- ZooKeep 分布式协调服务，Chubby的开源实现，采用 ZAB

## Paxos
- TODO

## Raft 
- TODO

## ZAB 
- 对Multi Paxos 算法的改进，和 Raft 大致相同
- 对于 Leader 的任期， Raft 叫 Term，ZAB 叫 epoch
- 复制状态时， Raft 的 心跳从 Leader 发往 Follower，ZAB 从 Follower 发给 Leader

## Gossip
- 无主设计，每个节点对等
- 每个节点会将数据改动发给其他节点
- 其他节点可在收到节点后继续发给其他的节点
- 没收到就转发，收到就停止？
- DynamoDB ？















