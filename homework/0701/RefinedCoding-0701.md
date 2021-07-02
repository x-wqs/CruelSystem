# MIT 6.824 Lab 2C - Persistence 持久化
- https://pdos.csail.mit.edu/6.824/labs/lab-raft.html

## 需求
- 基于Raft协议的服务在重启之后会恢复重启之前的服务
- 要求Raft支持持久化，才能客服重启故障
- 论文图2提高了Raft的持久化

## 实现
- 在现实视线中，Raft的持久化状态每次在改变的时候都会写入磁盘
- 然后重启的时候，会从磁盘读出
- 在本实验中，无需使用磁盘
- Raft会从持久化类直接读入，估计持久化类Persister.go已经实现好了
- 接口有Persister类的ReadRaftState()和SaveRaftState()

## 要求
- 实现Raft类的persist()和readPersist()
- 还要编码序列化Raft的状态，使用labgob编码器
- Labgob编码器类和Go的Gob编码器类似


