# ZooKeeper Lecture
- https://youtu.be/HYTDDLo2vSE
- https://blog.csdn.net/u014634338/article/details/106039491
- https://blog.csdn.net/wx1528159409/article/details/84633903
- https://club.perfma.com/article/395908
- https://www.jianshu.com/p/d01b1913cced
- 分布式协调，元数据管理，高可用，分布式锁

## Write 事务性请求
- 所有事务性请求必须求领导者来负责，或者转发给领导者
- 领导者收到半数以上节点写入成功的消息，则返回成功
- leader create zxid ?
- FT-VM test-and-save ?
- write linearizable ?
- not blocked by fuzzy snapshots
- ZooKeeper的层次模型用Java的ConcurrentHashMap来实现，一种线程安全的哈希表，使用分段锁来减少锁竞争，既安全又高效

## Read 非事务性请求
- 非事务性请求由本地服务器处理
- read not linearizable
- read not stable
- client keep the highest zxid
- reader's zxid must be no order than write ?

## Raft
- send reads to the leader
- read linearizable
