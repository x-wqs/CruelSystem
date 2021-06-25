# 0623 - ZooKeeper - P1 - 3

## 概要
- 协调进程
- 协调原语
- 促成共享寄存器和分布式锁服务
- 无需等待方面
- 消息驱动来清空缓存
- 基于客户端的先进先出保证

## 简介
- 支持配置，见参考27
- 领导选择
- 互斥访问
- 协调内核
- 加锁原语
- 简单无需等待的数据对象
- 无需等待的属性
- 保证顺序
- 实现一致性
- 简单的流水线架构，来保证先进先出的客户端顺序
- Zab协议
- 缓存主服务的标记，而不是每次探测
- 检测更新，Chubby使用租约来防止一个损坏的客户端阻挡系统
- 协调内核：提供无需等待的协调服务
- 协调暂停、食谱、厨房：强一致性原语
- 协调体检

## 2 ZooKeeper服务
- 客户端接口调用服务，也维护和服务器的连接
- 术语
- znode：内存里的节点，表示ZooKeeper数据
- 数据树，层次性信息
- 客户端创建会话，拥有会话句柄

### 2.1 服务概要
- 一组数据节点的抽象组合
- 常规节点，显示创建和删除
- 短暂节点，在会话结束的时候自动删除
- 监测只能收到一次事件
- 支持元数据或者配置
- 当前主服务写入这种信息
- 时间戳和版本计数器
- 观察状态变化的后续

### 2.2 客户端接口

## Termonology
primitive
incorporate
sophisticate
opted
faulty
herlihy
comprise
compromise
sub-second
manipulate
desirable way
illustration
ephemeral
monotonically
deliberately
exploit
ensemble





