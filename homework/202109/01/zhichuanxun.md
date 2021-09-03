# 秒杀系统

ref： 
https://time.geekbang.org/column/article/40153
https://www.cnblogs.com/54chensongxia/p/13609148.html
https://www.zhihu.com/question/54895548

## 需求

高性能： 动静分离，热点发现与隔离，请求的削峰与分层过滤，服务端优化

高可用：兜底

一致性：秒杀减库存方案

秒杀系统本质上是大并发，高性能和高可用的分布式系统
有基础的分布式系统需求

其它需求：
1. 数据要少: 用户请求的数据尽可能少和小，尽量避免序列化之类
2. 请求量要尽可能少（合并css js 文件）
3. 路径要短，RPC-》JVM
4. 依赖要少
5. 不要有单点

## 具体API
//TODO