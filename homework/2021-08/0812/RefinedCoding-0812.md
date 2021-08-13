# 计算机系统设计原理 - 多点原子性 - 分布式两阶段提交
- https://ocw.mit.edu/resources/res-6-004-principles-of-computer-system-design-an-introduction-spring-2009/online-textbook/atomicity_open_5_0.pdf
- Section 9.6.3
- https://developer.aliyun.com/article/765176
- https://segmentfault.com/a/1190000012534071

## 介绍
- 两阶段提交协议
- 支持分布式事务
- 投票 和 事务提交 两个阶段
- 集群有一个协调者节点，其他的为参与者

## 流程 - 投票阶段
- 协调者向所有参与者发送执行请求
- 参与者记录事务日志，但不提交事务
- 参与者返回结果给协调者，组着协调者的后续命令

## 流程 - 事务提交
- 协调者向参与者发出事务提交的通知
- 参与者提交事务，释放资源
- 参与者返回结果给协调者

## 缺点
- 单点问题
- 同步阻塞，所有参与者都要听从协调者的统一调配
- 数据不一致行

## 解决办法
- 超时机制
- 轮询机制


