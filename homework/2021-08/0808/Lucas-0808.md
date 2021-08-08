# 计算机系统设计原理 - CH9 原子性
- https://ocw.mit.edu/resources/res-6-004-principles-of-computer-system-design-an-introduction-spring-2009/online-textbook/atomicity_open_5_0.pdf
- P21 - P26
- 详解MySQL两阶段加锁协议
- https://segmentfault.com/a/1190000038163191
- MySQL 行锁、两阶段锁协议、死锁以及死锁检测
- https://www.huaweicloud.com/articles/d49a4e7f54e8980c9b0136501e1f6378.html

### 9.1.7 序列原子性和协调原子性

## 9.2 序列原子性 1 - 概念

### 9.2.1 达到序列原子性 - PUT

### 9.2.2 系统原子性 - 提交和金律

### 9.5.2 简单锁
- First, each transaction must acquire a lock for every shared data object it intends to read or write before doing any actual reading and writing. 
- 每个事务在读写共享数据对象之前，都要获取锁。
- Second, it may release its locks only after the transaction installs its last update and commits or completely restores the data and aborts.
- 在事务结束之后释放锁

### 9.5.3 两阶段锁 Two-Phase Locking



