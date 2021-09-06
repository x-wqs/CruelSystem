# Magnet:一个可扩展的、高性能的Apache Spark shuffle架构
- https://engineering.linkedin.com/blog/2020/introducing-magnet

## 简介
- 用于数据驱动决策分析的离线数据分析
- LinkIn 用 Spark 来做数据仓库，人工智能，机器学习，A/B测试，指标报表
- 从 2017 年起，每年应用 3 倍增长，超过一万节点
- 每天处理超过 PB 数据分析的 70%
- 在扩容和随机处理上还有缺陷

## 随机基础
- 和 MapReduce 一样，中间处理结果要随机分配到所有的节点上
- 在 Apache YARN 上运行，采用 Spark 的外部随机服务
- 采用推送的 Magnet 随机架构

## 架构特点
- 提高磁盘的读写性能
- 缓解了稳定性和扩展性
- 异步，不会被阻塞
- 和 Spark 的原生集成
