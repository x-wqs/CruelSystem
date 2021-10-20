0000-wormhole-read

# Wormhole: Reliable Pub-Sub to support Geo-replicated Internet Services
- https://www.jianshu.com/p/538d46cb3621

## 概要
- Facebook 的发布订阅系统
- 地理位置上的数据复制
- 为 TAO，图搜索，Memcache 等业务服务
- steady 稳定状态，每秒 35GB
- 低延迟发布更新
- 订阅者可以从灾难中恢复，无需 compromising 必要的性能损失？


## 简介
- 用户上传内容，写到数据库
- 其他业务对 learning of wirte 比较感兴趣
- 用户需要即时的通知
- 内存失效流水线，索引服务流水线
- 业务可以长轮询数据更新
- 和存储系统的生产环境负载 workload 交互
- 发布者从MySQL，HDFS，RocketDB 读取事务日志，产生更新
- 订阅者从上次的位置读取更新
- 如果给业务发送通知，有两个缺点
- 一是跨软件栈的通知
- 二是将通知写入存储，会引入可能失败的存储

## 术语
interpose
