# Lecture 3 - GFS
- https://youtu.be/6ETFk1-53qU
- https://github.com/refinedcoding/CruelSystem/blob/main/paper/02-gfs.pdf
- https://github.com/refinedcoding/CruelSystem/blob/main/paper/gfs-faq.pdf


## Questions
- Q1: Primary1挂了之后，启动Primary2了，然后Primary1又恢复正常了，GFS怎么处理的?
- A: 不同的Primary之间的租期不重复，等Primary1租期结束了，Primary2的租期才正式生效

- Q2: GFS如何快速实现快照？
- A: GFS创建快照的时候，并没有创建副本，而是将块引用数量加一，如果有修改，才会复制并追加。这样就避免频繁的拷贝。

- Q3：GFS块文件大小为64MB，如何才能节省磁盘？
- A: 创建块文件分配磁盘也是lazy的，这样能省磁盘空间和频繁的磁盘读写操作。

- Q4: How to handle replica failure?
- A: Master发现可用Replica的数量不够的时候会执行re-relication

- Q: Why does GFS use write ahead logs?
- Q: Why does GFS use checkpoints?
- Q: 描述GFS记录追加的流程，见图2.
- Q: 描述GFS Master奔溃重启的过程。
- Q: 如何处理多个客户端对同一块文件追加的冲突？
- Q: Seconday没有收到Primary的追加命令，客户端可以从Seconday读数据么
- Q: 描述Primary奔溃的处理流程
- Q: 带有过期数据的Chunk 4在其他副本宕机的时候是否可用？
- Q: Seconday写入失败的时候，Primary该如何处理？

## Teminology
many themes	,	许多主题
all the way to	,	一路走到
tradeoffs between performance and consistency	,	性能与一致性之间的权衡
SOSP	,	SOSP
weak consistency	,	弱一致性
single master	,	单主
awkward	,	尴尬的
circumstance	,	环境
anomalous	,	异常的
anomalies	,	异常现象
anomaly
retrospective	,	回顾性的



