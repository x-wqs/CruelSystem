# MIT 6.824 Lecture 12 - Frangipani
- Thekkath, Mann, Lee, SOSP 1997
- https://youtu.be/jPrUxfIcWWs
- https://pdos.csail.mit.edu/6.824/notes/l-frangipani.txt

## 概要
- 缓存一致性，cache oherence
- 分布式事务，distributed transation
- 分布式灾难恢复，distributed crash recovery

## 架构
- 用户，Frangipani，网络，Petal
- Petal，可复制的，带分区的，块存储服务
- Frangipani，存储目录，inode，文件内容，和空余位图？
- frangipani，去中心化的文件服务，高性能缓存

## 设计
- 强一致性
- 每个工作站都有缓存写回
- 所有操作都先写入缓存，速度非常快，比如文件创建和更名
- 一旦缓存过，再通过远程调用更新其他工作站？

## 缓存一致性协议
- 工作站A先写入缓存，而不是Petal
- 文件忙碌，正在使用数据
- 文件空闲，不用数据，但是加锁了，然后从Petal读数据，缓存，写入Petal，释放锁

## 分布式事务
- 使用锁服务
- 所有修改操作都加锁，其他服务器不会看到部分修改
- 锁服务具有缓存一致性，和事务原子性

