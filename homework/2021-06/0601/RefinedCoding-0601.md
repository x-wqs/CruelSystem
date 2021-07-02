# Consensus Algorithm 一致性算法  
- P1 - P3

## 1 简介  
- Raft分离了一致性算法的关键因素，比如主节点选举，日志复制，日志安全。  
- 日志只能从主节点流向其他服务器  
  
## 2 副本状态机  
- 处理相同的命令序列，结果则相同  
- ZooKeeper，复制状态机  
- 复制日志的一致性，接受了再返回  
- 不依赖时序  
- 日志的一致性  
- 完成只要集群的主题回应了  
  
## 3 Paxos的缺陷  
- 单一决定  
- 第二不太好理解  
- Chubby 没有开源？  
- 结果导致 Single decree decomposition  
  
## 4 为易理解性设计  
- 系统设计的权衡  
- 反义消除  
- 主节点选举，日志复制，安全，成员变化
  
## 5 Raft一致性算法  
- 管理主节点，并接受  
  
  
  
## Terminology  
Consensus	，	共识  
Consensus Algorithm	，	共识算法  
replicated log	，	复制日志  
replicated state machine	，	复制状态机  
practical system	，	实用系统  
key elements of consensus	，	共识的关键要素  
a stronger degree of coherency	，	更强的一致性  
a new mechanism	，	一种新机制  
overlapping majority	，	重叠多数  
a coherent group	，	一个连贯的群体  
survive the failures	，	在失败中幸存下来  
dominate the discussion	，	主导讨论  
influenced by	，	被影响  
primary vehicle	，	主要车辆  
unfortunately	，	很遗憾  
in spite of	，	尽管  
numerous attempts	，	无数次尝试  
approachable	，	平易近人的  
primary goal	，	首要目标  
understandability	，	易懂  
significantly	，	显着地  
facilitate the development of intuitions	，	促进直觉的发展  
be essential for	，	必不可少  
obvious	，	明显  
specific techniques	，	具体技术  
improve understandability	，	提高可理解性  
decomposition	，	分解  
state space reduction	，	状态空间缩减  
significantly easier	，	明显更容易  
novel features	，	新颖的特点  
flow from	，	从  
randomized timer	，	随机计时器  
resolve conflict	，	解决冲突  
a new joint consensus	，	新的共同共识  
majority	，	多数  
superior	，	优越的  
safety properties	，	安全特性  
replicated state machine problem	，	复制状态机问题  
strength	，	力量  
weakness	，	弱点  
arise	，	出现  
compute identical copies	，	计算相同的副本  
replicated state machine	，	复制状态机  
Single decree decomposition	，	单法令分解  


