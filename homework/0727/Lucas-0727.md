# Frangipani 一种可扩展的分布式文件系统
- https://pdos.csail.mit.edu/6.824/papers/frangipani-faq.txt
- https://zhuanlan.zhihu.com/p/237989253
- https://www.zhihu.com/column/c_1273718607160393728
- https://juejin.cn/post/6844903953352622087
- https://mit-public-courses-cn-translatio.gitbook.io/mit6-824/
- https://keys961.github.io/2019/04/26/%E8%AE%BA%E6%96%87%E9%98%85%E8%AF%BB-Frangipani/

## Frangipani 和 GFS 的区别
- GFS 具有文件系统的逻辑，没有缓存
- GFS 处理顺序化的大文件读写，这些大文件不适合缓存
- Frangipani 把文件系统的逻辑分发给工作站
- Frangipani 有许多机制来确保工作站缓存滞留一致性
- Frangipani 更接近于日常的文件系统

## 文件的创建
- Frangipani 比 GFS 需要花更多的时间
- 因为日志，所有的元数据需要写入两次，第一个是日志，第二次是文件系统

## 系统故障
- Frangipani 只恢复系统元数据，比如 i-node， 目录，空余的位图
- 文件内容在日志之前写入？
- 特殊情况是，程序使用了 fsync 操作

## 和 Petal 的联系和区别
- Petal 管理虚拟的磁盘
- Frangipani 调用 Petal 的服务
- Petal 和 Frangipani 的日志都存储在 Petal 的虚拟磁盘上
