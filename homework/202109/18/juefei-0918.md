### Introduction

关键字：database，shards data across Paxos state machines
自动reshard，

应用可以通过Spanner的特性来完成高可用（因为物理隔离）
每个Application一般复制了3-5份

存储方式：semi-relational

同时提供部分事务，且支持sql查询。

Spanner提供很高的自定义：可以指定哪个datacenter存储哪部分data，data离用户有多元，多少个备份等待。

Spanner提供两大特性：
> 提供最终一致性
> 基于时间戳的全局一致性

example: 如果事务T1发生在事务T2之前，那么T1的时间戳一定在T2之前。

这个实现是基于一个叫TrueTime的api

### Implemnetation

*universe* : A Spanner deployment
*Zone* : unit of administrative deployment
