## 0922

## 2.3 数据模型
- 模式化的半关系型表，支持查询语言和通用事务
- Megastore支持模式化的半关系型表和副本同步，Google or UCB?
- Bigtable 仅支持跨数据中心副本的最终一致性
- Megastore 比 Bigtable 简单，支持跨数据中心的副本同步，其他区别？
- 使用Megastore的系统有Gmail，Picasa，Calendar，Android Market, AppEngine
- Bigtable 缺少跨行事务，所以构建了Percolator
- Spanner的两阶段提交有性能问题，所以应用程序应该避免多度使用事务
- 一个宇宙有多个数据库，一个库有无限个表，每个表和行，列和版本号
- Spanner并非纯关系型，每行有行名作为主键
- TODO:多个有层次结构的表，交错结构

## 3 TrueTime 接口
- TrueTimeInterval，一个有时间不确定界限的时间方位，注意是不确定
- TT.now() 被调用的时绝对时间，带有闰秒的Unix时间
- 底层使用GPS和原子时钟，GPS信号弱要用天线，原子时钟不是很贵
- 后台使用Marzullo算法的变体检测时间错误的机器， TODO

## 4 并发控制
- 如何实现，外部一致事务，无锁只读事务，旧数据的非阻塞读取
- 不同区Spanner客户端的写去，两阶段提交会在准备阶段生成Paxos写入
- 读写事务，单独的写入，客户端无需循环等待
- 只读事务，单独的非快照读，无锁执行，不会阻塞写，
- 快照读取，旧数据的无锁读取
