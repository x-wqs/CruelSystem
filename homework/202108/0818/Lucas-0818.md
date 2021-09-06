# MIT 6.824 L13 - 分布式事务
- https://youtu.be/B6btpukqHpM
- https://pdos.csail.mit.edu/6.824/notes/l-2pc.txt

## 问题
- 分布式事务 = 同步控制 + 原子提交
- 传统数据库事务采用 begin / end 关键字 来顺序执行代码
- ACID，原子性，一致性，独立性，可持续性
- 可序列化的，顺序相同，则结果一致

## 锁的分类
- pessimistic 悲观锁，使用前加锁，冲突会造成任务延迟
- optimistic 乐观锁，读写记录时不用锁，提交时检查检查读写是可序列化，冲突就重试
- 冲突较多是，悲观锁块，否则乐观锁快

## 两阶段锁 Two-phase Locking
- 事务必须在使用之前获得锁
- 事务在提交或失败后才会释放锁
- 每个数据库记录都有锁
- 同一个数据可以在不同的服务器上有副本

## 两阶段提交 Two-phase Commit
- TODO
