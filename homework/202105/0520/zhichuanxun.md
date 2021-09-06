# GFS

## introduction

- component failures are the norm rather than exception
- files are huge by traditional standards
- most fils are mutated by appending new data rather than overwriting existing data
- co-design benefits the overall system
  
## design

### assumptions

- 系统由许多普通组件组成，出错是常态
- 系统存储了大量大文件，小文件需要被支持但无需优化
- 工作负载为：
  - 大量的流读取和小量的速记读取
  - 大量的append操作
- 生产者消费者队列
- 可持续高带宽比低时延更重要

### 接口

基本的：
  create, delete, open, close, read, write

扩展的:
  snapshot: 低开销创建copy或者文件树
  record append: 允许多个客户端同时追加数据，并保证原子性

### 架构

cluster: 1 Master +  1...* GFS chunkserver
(master 高可用？)

文件被分为固定大小的chunk，每个chunk由一个全局唯一的chunk handle标识。

chunk server 通过linux 文件的方式存储，并有三个副本来保证可靠性。

Master有文件的元数据： 命名空间，访问控制信息（[TODO]:啥访问控制信息，干啥的），文件到chunk的映射，当前chunk的位置。Master还管租约，gc，迁移和心跳。

应用是直接与chunkserver进行数据交换，不提供posix api，不需要linux vnode层的hook([TODO]:知识盲点)

没有缓存机制可以简化一致性问题（！），client会缓存元数据。
而在chunkserver中，通过文件的存储形式将缓存让渡给linux的缓存机制。

### 单一的master

确实，单一master很简化设计。
因为client在一次和master通信之后会缓存下这个信息，之后就直接和chunkserver通信来，缓存信息有过期时间，所以就算master挂了，在过期时间内能得到重启就没问题，在一定程度上解决了可用性问题，但万一在获取之前master就挂了呢。

### chunk大小

64MB，是一个相对比较大的值

lazy的空间分配可以防止碎片造成的空间浪费。

优点：
- 可以减少与主节点的交互
- 客户端可以在一个快上进行很多操作，并保持tcp链接减小网络的开销
- 减少了在master上的元数据的大小，这样可以将元数据存在内存中（内存/元数据*64MB的inmemory快速访问）

### 元数据

命名空间，映射，chunk位置

操作日志会持久化并备份

master会询问chunkserver关于chunk的位置信息

#### in-memory的数据结构

命名空间信息通过前缀压缩来减小大小([TODO]复习前缀树)

#### chunk的位置

chunkserver应该对自己disk的信息内容有最终话语权（感觉可以应用到设计中，就是将一部分控制权利让渡给从节点来达成一些目的）

#### 操作日志

日志记录了元数据，也可以用来确定并发操作的时间线。

日志有compact，checkpoint是一个btree。

master在一个单独的进程中切换以及创建checkpoint（和COW有啥区别呢？）

### 一致性模型

#### GFS保证的一致性

命名空间的操作是原子的：
  - 锁保证正确性
  - 日志确定了顺序

_consistent_: 所有客户端都可以观察到同样的数据
_defined_: consistent的，并且可以看到写了什么（有点不太理解）

这里设计了很多接下来的章节，回头再看

#### 客户端

- append rather than overwrite
- checkpoint
- write self-validation
- self-identifying records

## 系统交互

### 租约和变化顺序

数据流和控制流是解耦的

master给一个replica分配一个租约，这个租约叫primary，对于一个mutation，primary确定了所有replica变化的顺序。

租约持续60s，如果仍然处在mutation中，那primary可以通过心跳来延长租约。

control flow：

1. client向master询问当前lease以及其它replica的位置
2. client缓存这些数据
3. client push da1a，chunkserver存在LRU cache里
4. 所有replica收到数据后，client发送写的请求给primary
5. 所有replica按照primary安排的顺序一个个写1
6. 所有replica写完后回复primary
7. primary回复client

### 数据流

数据流的目标是最大利用网络。

数据流的传输是一个线性链路而不是拓扑图，这样可以最大利用outbound
网络最近的节点传输，distance通过ip来计算

### 原子性的记录追加

client只确定数据，gfs保证原子性。

记录追加是一个耗时的操作，传统情况下会变成生产者/消费者队列

record append受到maximum size的限制

### 快照

copy-on-write

## Master的操作

master 管namespace和infrastructure

### namespace

namespace有锁

对于一个路径，会给所有父路径加上读锁，给自己加上读（文件写）锁

### replica placement

数据可靠性和可用性，带宽

### creation, re-replication, rebalancing

平均的disk利用率

倾向于复制活跃文件的chunk

### GC

GC基于文件和chunk level的

#### mechanism

和操作系统删除文件类似，不过由master和server协调完成

#### discussion

回收策略可配置

### Stale Replica Detection

通过版本号来区分是否过期，和raft的leader election有点像

## 容错和诊断

### 高可用

#### 快速恢复

master和chunkserver都可以在数秒内重启

#### Chunk的复制

checksum校验

备选方案：奇偶校验，Erasure Code([知识盲区TODO])

#### Master的复制

master会快速重启

只是单纯复制操作日志，然后可以快速重启(既然都复制日志了，为啥不连进程也一起复制，多个master一起干活多好啊)

shadow master可以在master down的时候提供只读访问

### 数据完整性

跨服务器校验不现实，每个chunkserver独立维护checksum

chunkserver空闲时候会检查

### 诊断工具

日志

## 度量

### 小规模benchmark

1master + 2replica
16chunkserver + 16client

读取的话，client越多越慢

写入比想象的慢，网络协议栈和pipeline不适应(?)

记录追加速度受限于最后一个chunkserver的带宽

### 实际应用

master启动很快，但会颠簸一段时间（用来便利所有chunkserver以获取到位置，这段时间有的就不可用）

### 工作负荷

## 经验

## 相关工作

primary-copy的方法