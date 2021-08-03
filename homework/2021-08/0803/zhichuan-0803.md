# Memcache in Facebook

TODO： https://xie.infoq.cn/article/fa1f0f9ac1cfee7845f7b29fe cache policy

## Background

读 >> 写，缓存层可以保护DB

数据分片和一主多从

## Design

- QPS 10亿
- 多区域
- 平滑升级
- 最大努力最终一致性（best-effort eventual consistency）

## Cache Policy

- read：demand-filled look-aside
  先从memcache中读，读取失败则从db中放数据到cache
- write：write-invalidate
  先更新db，然后删除cache

## Latency and Load

主要目标
低 latency + load（cache miss）

## Scale-out

常见的consistent hashing

[!遇到问题，可以分析请求的依赖关系，依赖关系是一个容易忽略的隐藏条件]

不同数据读取顺序有着依赖关系，可以表示成DAG，fanout

all-to-all communication，每个server可能需要与所有的cache server通信

- 多个数据同时到来会造成网络拥堵，incast congestion
- 每个cache都有一部分数据，可能成为瓶颈

## Solution:

latency：

- 应用层 Parallel Requests and Batching batch读
- 缓存层 在server 和 cache之间抽象一个client，client通过UDP与cache通信，自己handle各种情况，包括congestion control

[!TCP->UDP也是一个常用的解决方法，但要注意丢包率]

load：

引入lease

- cache miss返回lease id
- write-invalidate时，lease id失效
- 写数据时加lease并verify

## Thundering Herds

扩展lease机制，控制lease的发放速率

## Stable Values

如果能容忍过期数据，可以在数据被删除后暂时移到其它地方，server决定等待还是读过期数据

## Pooling

[Bulkheads设计模式]

## Gutter

备胎， 应急电源

## Scaling

[单纯的HPA可能会导致更严重的incast_congestion]

将cache server分成热点和其它

## Regional Invalidation

多个集群时，同一数据可能有多个版本

广播[链式或tree]

## Warmup

cache上线前warmup一下

## remote maker

中间状态，写入master DB但未同步到replica DB

## 总结

在面对网络问题时，多一层抽象（比如Router）可以用来解决很多问题：congestion（hierarchy）, latency(TCP->UDP, router来解决丢包重试问题等)