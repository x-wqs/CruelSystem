# GFS (1-10)

## abstract
scalable distributed file system for large distributed data-intensive apps

hundreds of terabytes of storage across thousands of disks on >1000 machines

## why GFS
- 节点故障频繁发生
- 文件巨大–数GB
- 大多数文件都通过在末尾附加来修改
- 几乎不存在随机写入（和覆盖）
- 高持续带宽比低延迟更重要
- 优先处理批量数据

## 典型工作 on GFS
两种读取：大型流读取和小型随机读取
- 大型流式读取通常读取1MB或更多
- 通常，应用程序会读取文件中的连续区域
- 小随机读取通常只有几个偏移量任意的KB
 还有许多将数据追加到文件的大顺序写入
- 与读取类似的操作大小
- 写入后，很少再修改文件
- 任意偏移量的小写操作不一定有效
 多个客户端（例如〜100个）同时附加到一个文件中
- 例如 生产者－消费者队列，多路合并

## 接口设计

不符合POSIX，但支持典型的文件系统操作：创建，
删除，打开，关闭，读取和写入
- 快照：以低成本创建文件或目录树的副本
- 记录追加：允许多个客户端将数据追加到同一文件
同时
- 至少保证第一个附加元素是原子的

## 架构
decouple data flow from control flow
- 客户端与主服务器交互以进行元数据操作
- 客户端直接与块服务器进行所有文件操作的交互
=> 可以通过计划昂贵的数据流来提高性能
基于网络拓扑

## master结点
