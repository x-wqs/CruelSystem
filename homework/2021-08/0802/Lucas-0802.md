# 计算机系统设计原理 - CH9 原子性
- All or Nothing，序列协调
- Before or After，协调并发线程
- https://zhuanlan.zhihu.com/p/336136705
- http://www.lucienxian.top/2019/08/18/Atomicity-All-or-Nothing-and-Before-or-After%E2%80%94%E2%80%94MIT6-824/
- https://ocw.mit.edu/resources/res-6-004-principles-of-computer-system-design-an-introduction-spring-2009/online-textbook/atomicity_open_5_0.pdf

## 9.1 原子性
| 领域     | 序列原子性       | 协同原子性         |
|----------|------------------|--------------------|
| 数据库   | 更新多个数据记录 | 多线程更新共享数据 |
| 硬件架构 | 处理中断和异常   | 注册重命名         |
| 操作系统 | 管理程序接口     | 打印机队列         |
| 软件工程 | 应用层容错       | 有界缓存           |

### 9.1.1 数据库的原子性 All or Nothing

### 9.1.2 中断接口的原子性 All or Nothing

### 9.1.3 层级应用系统里的原子性

### 9.1.4 原子性可有可无的操作

### 9.1.5 协同并发线程的原子性 Before or After
- 如果线程并发执行的结果，和不同线程按顺序逐个完成的结果一致
- 当前线程不需要知道其他线程的操作

### 9.1.6 数据更正和序列化

### 9.1.7 系列原子性和协调原子性
