## 9999-memcached-read.md
- 演变：但集群，多集群，多区域
- 读多写少，Memcache之前，数据分片，一主多从
- 需求：10亿QPS，多区域，平滑升级，最终一致性
- 缓存策略，读取，demand-filled look-aside，边顾周围，按需填充
- 缓存策略，写入，write-invalidate，写入并废除
- 缓存侧率，TODO-分布式事务
- 横向扩容，Scale-out，一致性哈希
- 横向扩容，High Fanout，有向无环图，DAG，放射状读取
- 横向扩容，减少延迟，incast congestion，添头拥塞，减少RTT，RoundTrip Time，循环时间
- 客户端分成两部分，SDK，Proxy或mcrouter，和memcached接口一直，增加一层抽象
- TODO: UDP / TCP
- TODO: connection coalescing，无缝连接
- TODO: Incast congestion，由于通信模式问题引起的拥塞

