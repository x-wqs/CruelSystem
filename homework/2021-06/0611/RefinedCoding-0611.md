# MIT 6.824 Lab 2 Requirements
- https://pdos.csail.mit.edu/6.824/labs/lab-raft.html

## 概要
- 实现一个支持容错的键值存储系统
- 使用复制状态机协议
- 服务分片
- 通过存储状态的完整拷贝来实现容错
- 失败会造成副本保存不同的数据拷贝
- 把客户端请求作为日志放入队列
- 每个日志都带有索引序号
- 带有序号的日志最终都会提交，会丢么
- 论文图2： 状态，Vote RPC，AppendEntries RPC，Server Rules
- 无需实现集群成员变化

## Lab 2 A
- rf := Make(peers, me, persister, applyCh) 四个参数
- 领导选择，
- 实现VoteRequest和VoteResponse
- 如何实现不同节点不同的超时等待


## 参考
- Raft流程图
- https://pdos.csail.mit.edu/6.824/notes/raft_diagram.pdf
- Students' Guide to Raft
- https://thesquareplanet.com/blog/students-guide-to-raft/
- Paxos / Bolosky
- https://static.usenix.org/event/nsdi11/tech/full_papers/Bolosky.pdf

## 术语
flaky network	,	片状网络
live replica	,	活复制品
leave off	,	离开
log entry	,	日志条目
blindly	,	盲目地
simulate	,	模拟
stimulate	,	刺激
volatile state	,	不稳定状态







