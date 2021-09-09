## 调度

### 需求

- 满足的需求
	- 副本数量刚好
	- 分布不同机器
	- 副本可迁移新加节点
	- 节点下线后需移走数据
- 优化
	- Leader分布均匀
	- 储存容量均匀
	- 访问热点均匀
	- balance速度
	- 管理节点状态

### 基本操作

- Add replica
- Remove replica
- Transfer leader between replicas

### 信息收集

- 节点定期向pd汇报信息
	- PD checks living store or newly added store through heartbeat
	- heartbeat contains status info of store
- Leader定期向PD汇报信息

### 策略

- replica 数量正确
	- PD 发现replica 数量不满足，需要 Add / remove replica
- 多个replica 不同位置
- 副本在store间分布均匀
- leader在store间分布均匀
- 访问热点在store间分布均匀
- store的存储空间相等

### 实现

- PD collects information through heartbeat from leader / store
- generates operation sequences according to info and strategy
- returns operations to region leader
- examines the result through heartbeat