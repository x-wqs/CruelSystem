# Consensus Algorithm 一致性算法      
- P10 - P12  
    
### 5.5 随从和备选奔溃    
- 从备恢复更简单  
- 领导发现失败，会一直重试  
  
### 5.6 时序和可用性  
- 安全性不依赖于时序  
- 牺牲可用性，换区安全？  
  
## 6 集群成员变化    
  
    
## Terminology  
leader based consensus algorithm  
Viewstamped Replication	,	视图标记复制  
Raft uses a simpler approach	,	Raft 使用更简单的方法  
log entries from oder terms	,	来自其他术语的日志条目  
voting process	,	投票过程  
hold all the committed entries	,	保存所有提交的条目  
the voter denies its vote	,	选民拒绝投票  
determine	,	决定  
previous term	,	上一届  
old log entry	,	旧日志条目  
eliminate problems	,	排除问题  
Log Matching Property	,	日志匹配属性  
conservative approach	,	保守方法  
incurs this extra complexity	,	导致这种额外的复杂性  
more precisely	,	更确切地说  
Leader completeness	,	领导者完整性  
leader failures	,	领导失误  
idempotent	,	幂等的  
a timely manner	,	及时地  
aspect	,	方面  
inequality broadcaseTime	,	不等式宽限时间  
magnitude less than the election timeout	,	幅度小于选举超时  
MTBF	,	平均无故障时间  
occasionally	,	偶尔  
changeover	,	转换  
they risk operator error	,	他们冒着操作员错误的风险  
incorporate them	,	合并他们  
atomically	,	原子地  
potentially	,	潜在地  
ensure safety	,	确保安全  
transitional configuration	,	过渡配置  
joint consensus	,	共同共识  
individual servers	,	个人服务器  
compromising safety	,	危及安全  
future decisions	,	未来的决定  
the rules of C(old)	,	C（旧）的规则  
unilateral	,	单方面  
without approval	,	未经批准  
irrelevant	,	无关的  
unilateral decisions	,	单方面决定  
initially	,	原来  
avoid availabity gaps	,	避免可用性差距  
additional phase	,	附加阶段  
step down	,	下台  
disrupt the cluster	,	扰乱集群  
eventually	,	最终  
poor availabity	,	可用性差  
disregard RequestVote	,	无视请求投票  
minimum election time of hearing	,	听证的最短选举时间  
avoid disruption	,	避免中断  
incorporate more client request	,	合并更多客户请求  
a pratical system	,	实用的系统  
grow without bound	,	无限成长  
replay	,	重播  
discard obsolete information	,	丢弃过时的信息  
accumulated in the log	,	在日志中积累  
snapshotting	,	快照  
discarded	,	丢弃  
incremental approaches	,	渐进式方法  
log structured merge tree	,	日志结构化合并树  
a fraction of the data	,	数据的一小部分  
spread the load of compaction	,	分散压实负荷  
complexity	,	复杂  
last included index	,	最后包含的索引  
cluster membership	,	集群成员  
occasionallylag behind	,	偶尔落后  
InstallSnapshot RPC	,	安装快照 RPC  
all superseded	,	全部取代  
a prefix of its log	,	其日志的前缀  
depart from	,	离开  
decision conflict	,	决策冲突  
reorganize the data	,	重新组织数据  
TODO 5.4.3 - 5.5  
  
  
  
  
   
  