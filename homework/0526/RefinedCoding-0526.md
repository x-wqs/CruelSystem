# Lecture 4: Primary-Backup Replication
- 45 - 90 minutes
- https://www.youtube.com/watch?v=gXiDmq1zDq4

## 概要
- Hypervisor监测器
- 主机和备份机只要一个停止网络通讯，将启动另一个机器，提供单点服务

## 磁盘IO
- 监测器模拟磁盘接口
- 外部磁盘更快，无需拷贝

## 什么时候
- 主机发生变化的时候
- 不是确定结果的指令？

## 什么变化
- 在主备机器上执行后结果相同

## TODO-因差异产生的灾难流程

## 反弹缓存
- 在主备机器上的内存同一点

## 非确定的指令

## 数据库的同步
- 用输出原则
- 等备份机确认后，在回复客户端
- 回复只读操作，主机无需等待

## TODO-Page-6-7

## Terminology
- ordinarily
suppressed by
goes live, （系统）启动
sole
underneath,下面的，底层的
divergence，差异；分歧；分散，发散；（气流或海洋的）分开处
deterministic consequence，确定的结果
identical,同一的；完全相同的
by induction，通过归纳法
uniprocessor,单核处理器
bounce buffer,反弹内存，Stack？
lag by,滞后的
yield different，产生不同
constraint,约束；局促，态度不自然；强制
emit，发出，放射；发行；发表
eliminate,消除；排除







