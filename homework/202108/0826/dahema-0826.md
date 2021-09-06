https://zhuanlan.zhihu.com/p/130332285
# 分布式一致性算法-Paxos、Raft、ZAB、Gossip

## 需求
- 数据不能存在单个节点（主机）上，否则可能出现单点故障。
- 多个节点（主机）需要保证具有相同的数据。

## 定义
一致性就是数据保持一致，在分布式系统中，可以理解为多个节点中数据的值是一致的

## 分类
- 强一致性，保证系统改变提交以后立即改变集群的状态。模型：Paxos，Raft（muti-paxos），ZAB（muti-paxos）
- 弱一致性，最终一致性，系统不保证改变提交以后立即改变集群的状态，但是随着时间的推移最终状态是一致的。模型：DNS系统，Gossip协议

## Paxos，需要进一步阅读
- Proposal提案，即分布式系统的修改请求，可以表示为[提案编号N，提案内容value]
- Client用户，类似社会民众，负责提出建议
- Propser议员，类似基层人大代表，负责帮Client上交提案
- Acceptor投票者，类似全国人大代表，负责为提案投票，不同意比自己以前接收过的提案编号要小的提案，其他提案都同意，例如A以前给N号提案表决过，那么再收到小于等于N号的提案时就直接拒绝了
- Learner提案接受者，类似记录被通过提案的记录员，负责记录提案

## Raft
- Leader总统节点，负责发出提案
- Follower追随者节点，负责同意Leader发出的提案
- Candidate候选人，负责争夺Leader

## ZAB
- 和Raft的区别：
  - 对于Leader的任期，raft叫做term，而ZAB叫做epoch
  - 在状态复制的过程中，raft的心跳从Leader向Follower发送，而ZAB则相反。

## Gossip
- Gossip算法每个节点都是对等的，即没有角色之分。
- Gossip算法中的每个节点都会将数据改动告诉其他节点
