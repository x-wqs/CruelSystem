# MIT 6.824 Lab 2C - Persistence 持久化
- https://pdos.csail.mit.edu/6.824/labs/lab-raft.html

## 需求
- 基于Raft协议的服务在重启之后会恢复重启之前的服务
- 要求Raft支持持久化，才能客服重启故障
- 论文图2提高了Raft的持久化

## 实现
- 在现实视线中，Raft的持久化状态每次在改变的时候都会写入磁盘
- 然后重启的时候，会从磁盘读出
- 


