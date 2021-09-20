# 从零用户到百万用户的扩容
- Alex Xu Chapter 1

## Database Scaling

- 垂直扩展
单节点提供更强的计算能力；如提供更好的CPU、更大的磁盘等。
限制：1. 单节点有明确的上限 2. 单点风险 3. 成本更高

- 水平扩展

- - sharding
- - resharding
- - celebrity problem
- - join and denormalize

## 总结
1. 保持web层的无状态
2. 每一层都需要冗余
3. 可以尽量多的缓存数据
4. 支持多数据中心
5. CDN 静态资源加速
6. 将数据进行sharding获得更好的扩展性
7. 将业务拆分到不同的service
8. 打造良好的监控系统
