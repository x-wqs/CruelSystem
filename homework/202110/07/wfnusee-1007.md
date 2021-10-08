# 一致性哈希
在一致性哈希的基本设计里 我们有一个问题：
1. 很难保证每个服务的负载很均衡。很可能部分节点需要负责的key range很大 另一些很小。
（原文中另一个问题我觉得本质是一样的 无非就是初始化分配是否均衡及节点上下线时分配是否均衡）

解决方案也很简单，在环上放置一些重复的虚拟节点即可，多个重复的节点会对应同一个物理节点。
这样keys的分布也自然会变得比较均衡。
