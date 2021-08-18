# MIT 6.824 L13 - 分布式事务
- https://youtu.be/B6btpukqHpM
- https://pdos.csail.mit.edu/6.824/notes/l-2pc.txt

## 问题
- 分布式事务 = 同步控制 + 原子提交
- 传统数据库事务采用 begin / end 关键字 来顺序执行代码
- ACID，原子性，一致性，独立性，可持续性
