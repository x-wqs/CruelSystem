# GFS (1-10)

## abstract
scalable distributed file system for large distributed data-intensive apps

hundreds of terabytes of storage across thousands of disks on >1000 machines

# why GFS
- 节点故障频繁发生
- 文件巨大–数GB
- 大多数文件都通过在末尾附加来修改
- 几乎不存在随机写入（和覆盖）
- 高持续带宽比低延迟更重要
- 优先处理批量数据

