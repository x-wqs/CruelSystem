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