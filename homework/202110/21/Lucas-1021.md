0000-wormhole-read

# Wormhole: Reliable Pub-Sub to support Geo-replicated Internet Services
- https://www.jianshu.com/p/538d46cb3621

## 2 问题
- 存储大量用户产生的数据
- 比如状态更新，评论，点赞，分享
- 根据读写优化选择不同的存储
- vast user base 大量用户基础
- geo-replicated 地理上的数据复制
- employ caching system 大量使用缓存系统，比如Memcached，TAO
- underlying 底层存储系统没有 inundated 没有更新？
- 让应用程序重定向的数据库轮询
- dissemination 全套系统？
- 1. 不同的消费速度
- 2. 至少传递一次
- 3. 更新的有序传递
- 4. 容错
- heterogeneous 异构特性，imposed 自带特性1和4

## 3 架构

### 3.1 Caravan 拖车分配

### 3.2 过滤机制

### 3.3 可靠传递

#### 3.3.1 单拷贝可靠传递

#### 3.3.2 多拷贝可靠传递

## 4 负载评估

### 4.1 可扩展性

#### 4.1.1 应用程序数量的扩展

#### 4.1.2 订阅者的吞吐量

#### 4.1.3 生产环境中发布者的吞吐

## 术语
- interpose
- impose
