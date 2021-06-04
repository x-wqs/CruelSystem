# MapReduce continued
## 一致性/确定性：
Map Reduce 的输出均为原子操作，保证系统一致性（相同输入产生相同输出）

## 性能分析
- Bottleneck: RPC 的网速带宽

    输入是在 GFS 而不是本地

- Google 解决方案：
    
    GFS 每个文件有3份拷贝，分配 mapper 的时候大部分读取都可以在本地完成

## 总结
Pros：
- 适合处理海量数据的分布式计算模型
- 简洁通用的编程模型，封装了各种复杂繁琐的细节
- 高容错性
- 高扩展性

Cons：
- 不适合实时、迭代（iterative）数据处理
- 不能处理图形数据
- 不适合小数据、小文件处理
- 延迟高

未解决的问题：
- 高昂的维护成本：

    商业场景的复杂，协调 MapReduce 与业务逻辑困难
- 时间性能不足：
    
    比如分片等环节


## Q&A
1. Mapper 和 Reducer是同时工作还是先Mapper 工作还是 Reducer工作的么？

Mapper要结束了后Reducer才能运行

1. 运行过程中一个Mapper或者Reducer挂了怎么办？

重新分配一台机器做 具体见D2 notes

2. Reducer一个机器Key特别大怎么办？

加一个random后缀，类似Shard Key

3. Input 和 Output 存放在哪？

存放在GFS里面

4. Local Disk 上面的Mapper output data有木有必要保存在GFS上面？要是丢了怎么办？

不需要，丢了重做就好

5. Mapper 和 Reducer 可以放在同一台机器么？

这样设计并不是特别好，Mapper 和Reducer之前都有很多需要预处理的工作。两台机器可以并行的预处理。

