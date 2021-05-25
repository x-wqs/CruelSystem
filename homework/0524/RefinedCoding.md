# Fault-Tolerant VM    
- P7 - P10    
- https://github.com/refinedcoding/CruelSystem/blob/main/paper/vm-ft.pdf    
    
    
### 3.5 网络IO的实现问题    
- TCP模式可以让FT无需切换线程上下文    
    
## 4. 另一种设计    
    
### 4.1 共享磁盘和非共享磁盘    
- 默认共享磁盘，可能磁盘崩溃    
- 或者用外部磁盘，非共享磁盘，各自维护内部状态，需要两份拷贝同步    
    
### 4.2 在备份虚拟机上执行磁盘读写    
- 备份机默认不读    
- 或者读数据，操作数据，不必复制    
- 也可以主机读    
    
## 5.性能评估    
    
### 5.1 基本测试结果    
    
### 5.2 网络基准测试    
    
## 6 相关工作    
    
## 7 总结和未来展望    
    
## FAQ    
- 物理机只能确定性执行，更难    
- 托管机模拟控制和时序    
- 容错在虚拟机监控器上实现    
- GFS FT on storage, more efficient    
- VM FT on computing, complex    
- GFS vs S3    
- 缓存网络报和磁盘块    
- 管理程序负责生成随机数，而非在主机和备份机分别生成    
    
    
## Termilogy    
subtlety，微妙，精明    
hypervisor，管理程序    
buggy    
malicious    
fail-top    
Byzantine，拜占庭问题    
    
    
    
    
    