# Lecture 4: Primary-Backup Replication
- 45 minutes
- https://www.youtube.com/watch?v=gXiDmq1zDq4

## 概要
- 这一章讲如何用主备复制来容错，用VMware FT VM来举例

## 支持情形
- 机箱风扇，处理器过热
- 地震
- 不支持硬件或者软件错误，或者人为的配置错误

## 两种容错的方法
- 状态转移，状态大，慢
- 复制状态机，流量小，复杂，比如FT VM
- Lab 2/3/4会使用第二种办法

## 副本相同的层次
- 软件层面，GFS，更有效，主机向备份机发送高层操作
- 应用程序需要支持容错，比如转发操作
- 机器层面，寄存器或内存
- 这篇论文只讨论机器层面，对软件透明，可以运行现有的操作系统和服务器软件

## 其他问题
- Q1: 如何快速恢复？
- Checkpoint + Log？



## Terminology
sparse，稀疏的
fail-stop failure, 故障停止失效，恢复正常？
overhead
overheat
trip over，绊倒
overhead
defect，缺陷
correlated
