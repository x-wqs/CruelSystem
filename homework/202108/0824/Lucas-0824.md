
为什么要选举leader？  
[Lucas] 比如数据库集群，同一个分区有多个副本，只能有一个副本可写，不然就会有数据冲突。这个可写的副本，就是Leader，其他可读的副本就是Follower。  
CAP原理，最后一个是要处理网络分区，比如网络一分为二，出现两个可写怎么办？一致性算法会保证只有一个Leader  

不能随便定一个活着的人就是leader吗？  
[Lucas] Leader只有一个，原来Leader挂了的话，大家会选一个唯一的新的Leader。  

分布式系统之间是怎么通信的？   
[Lucas] 一般节点都在不同的机器上，有可能在不同的网络和数据中心，大家互相通过网络协议，比如TCP通信，具体到应用层，就是HTTPS，GRPC。  

以上是群友的个人理解，如有不对，请大家指正。谢谢了！  
