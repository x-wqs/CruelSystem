
> 名词：RAID redundant Array of Independent Disks

用来提供单机冗余

1. Tika是基于KV进行存储
2. 使用RockesDB作为存储端
3. 使用Raft来进行主从备份

数据流：data先打到Raft上，由Raft进行commit，之后Raft调用RocksDB进行存储将数据进行持久化，具体流程和6.824里说的应该差不多。

### Region
现在假设一个KV系统只有一个副本

简单地说，就是把一段key放在一个节点上，与之相对应的方法就是根据key的hash值来存

优点：水平扩容和负载均衡
Tika是基于region进行复制的

### MVCC
TiKV使用在key后面添加版本号来进行版本控制

### 事务
使用Percolator，文章也没细讲

