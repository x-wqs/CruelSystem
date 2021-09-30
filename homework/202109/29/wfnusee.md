# Robotron: Top-down Network Management at Facebook Scale
目标： 通过减少人们对设备的直接操作，降低事故和错误发生的概率。

网络配置管理是更高层网络功能的基础。
网络管理就是需要人们高度交互的任务。
传统的方式需要引入很多的运维操作。

一些挑战：
- 分布式配置文件
- Multiple domains
- 版本
- 依赖
- 不同的供应商
- 代码即配置
- 验证
- 可扩展性

Facebook - network of networks
包括 POP 集群 ； backbone ； 多个大型数据中心

### POP 
building new pop 是最复杂的任务
还有 改变 BGP 配置； provision new peering 等操作

### DC
由DR提供外部连通性
DC的配置改变的更频繁一些

### Backbone

## Robotron
- FBNet
- Nerwork Design
- Config Generation
- Deployment
- Monitoring

## FBNet
对网络的建模，我理解就是存储了所有网络组件元数据。