# Consensus Algorithm 一致性算法    
- P7 - P9  
  
## 5.3 复制日志  
  
## 5.4 安全性  
  
### 5.4.1 选举限制  
  
### 5.4.2 从上阶段提交条目   
  
### 5.4.3 安全性争议  
  
## 5.3 随从和备选奔溃  
  
  
## Terminology (P4)  
latest term server	,	最新术语服务器  
volatile state	,	不稳定状态  
monotonically	,	单调地  
timeout elapse	,	超时时间  
repeat during idle period	,	在空闲期间重复  
log inconsistency	,	日志不一致  
trigger independently	,	独立触发  
completeness	,	完整性  
key safty property	,	关键安全财产  
involves an additional restriction	,	涉及额外的限制  
role of timing	,	时机的作用  
raft cluster	,	筏式集群  
passive	,	被动的  
issue no requests	,	不发出请求  
divide time into term	,	将时间划分为术语  
arbitray	,	任意  
result in	,	造成  
split vote	,	分裂投票  
initiate an election	,	发起选举  
majority of the full cluster	,	整个集群的大部分  
no emerging leader	,	没有新兴的领导者  
a single leader manages the cluster	,	一个领导者管理集群  
without choosing a leaderensure	,	不选择领导确保  
ovserve the transition	,	监督过渡  
observe an election	,	观察选举  
logical clock	,	逻辑时钟  
detect obsolete information	,	检测过时的信息  
monotonically	,	单调地  
a timely manner	,	及时地  
issue RPCs in parallel	,	并行发出 RPC  
heartbeat mechanism	,	心跳机制  
as long as	,	只要  
maintain their authority	,	维护他们的权威  
election timeout	,	选举超时  
begin an election	,	开始选举  
increment	,	增量  
establish its authority	,	确立其权威  
lose the electionlegitimate	,	失去选举合法性  
leader's term	,	领导任期  
menting its term	,	提到它的术语  
without extra measures	,	无需额外措施  
repeat indefinitely	,	无限重复  
randomized election timeout	,	随机选举超时  
fixed interval	,	固定间隔  
spread out	,	扩散;传播开  
handle split vote	,	处理分裂投票  
elapse before	,	过去  
likelihood	,	可能性  
committed	,	坚定的  
design alternative	,	替代设计  
unique rank	,	唯一等级  
subtle issues	,	微妙的问题  
reset progress towards	,	重置进度  
new corner cases	,	新的角落案例  
randomized retry approach	,	随机重试方法  
servicing client requiest	,	服务客户要求  
replicated state machine	,	复制状态机  
in parallell	,	并行  
be safely replicated	,	被安全复制  
eventually	,	最终  
inconsistencies between logs	,	日志之间的不一致  
durable	,	耐用的  
log entry	,	日志条目  
subtleties	,	微妙之处  
keep track of 	,	不要跟丢  
highest index	,	最高指数  
high level of coherency	,	高度一致  
predictable	,	可预测的  
constitute	,	构成  
log matching property	,	日志匹配属性  
consistency check	,	一致性检查  
induction step	,	诱导步骤  
preserve	,	保存  
identical	,	完全相同的  
leader crashes	,	领导崩溃  
compound over	,	复合  
series of leader	,	系列领导者  
differ from	,	与......不同  
scenarios	,	场景  
its term	,	它的期限  
missing and extraneous entries	,	缺失和无关的条目  
span multiple terms	,	跨越多个术语  
handle inconsistencies	,	处理不一致  
conflicting entries	,	冲突的条目  
coupled with	,	加上  
rejection	,	拒绝  
decrements	,	递减  
infrequently	,	很少  
restore log consistency	,	恢复日志一致性  
log replication mechanim exhibits	,	日志复制机制展示  
leader commits several log entries	,	领导者提交了几个日志条目  
leader completeness property	,	领导者完整性属性  
precise	,	精确的  
proof sketch	,	校样草图  
  
TODO-541  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
    
  