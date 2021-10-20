# Wormhole: Reliable Pub-Sub to support Geo-replicated Internet Services

wormhole 自身不存储消息而是非侵入式的
wormhole 和数据源部署在一起 
* 直接读取 transaction logs
* geo-replication by data source

解决的问题：
* Delivery gurantees: 不用ack每一条消息 采用周期的 datamarker 本质就是批量的ack； 提供至少一次语义
* Dealing with shards: 
* Dealing with recovering and slow consumers:
慢节点不会对快节点产生直接的阻塞；可以用多个reader从不同的地方开始读；