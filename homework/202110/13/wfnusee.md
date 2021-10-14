# design key-value store
kv存储一般说的是key value数据库，是一种非关系型数据库。
value可以是各种东西 比如 stirng list object 等。

```
put(key, value)
get(key)
```

foundationdb 就是一种有序的key-value存储

* key-value 一般比较小 小于 10kb
* 可以用来存储大数据
* 高可用
* 高扩展
* 自动扩展
* 一致性
* 低延迟

## 单节点
最简单的实现用hash table即可；将一切存储在内存中即可。

### 分布式
CAP理论：
一致性；高可用；分区容错性 三者取其二；而分区容错性一般是必须支持的

我觉得文章里对CAP的解读有问题。 CP的系统；n3挂了 一定要拒绝写吗？3个有两个work也可以继续工作吗？