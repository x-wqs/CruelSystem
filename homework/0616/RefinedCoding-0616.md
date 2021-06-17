# Raft Lab 2B - Log
- 参考： https://pdos.csail.mit.edu/6.824/labs/lab-raft.html
- 测试步骤： go test -run 2B -race
- 难度： 难

### 测试细节
- 实现主从节点追加日志
- 通过TestBasicAgree2B()
- 参照图2，通过AppendEntries RPC发送和接受新的日志
- 参照5.4.1，实现选举限制
- 达到一直的方法一是重复选举
- 方法二是暂时不发心跳而去赢得选举
- 实时CPU超过一分钟，或者用户CPU超过五秒钟，则超时处理
- 实验要求用循环检测事件，看看能否支持异步处理，或者事件注册

### 测试用例
- 基础一致协议
- RPC字节数量
- 过多随从掉线
- 并发启动
- 从新加入分区领导者
- 领导从错误随从日志中恢复
- RPC调用次数不高



