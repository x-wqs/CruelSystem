# Consensus Algorithm 一致性算法  
- P4 - P6  
  
## 5 Raft一致性算法  
- 管理主节点，并接受  
- Raft把一致性问题分解成独立的子问题
- 主节点选取
- 日志复制
- 选取安全
- 主节点只追加
- 日志匹配
- 主节点完成性
- 状态机安全

### 5.1 Raft基础
- 集群有五个节点，每个节点三种状态：主节点，备选，随从
- 一般情况，一个主节点，多个从节点
- 主节点处理所有节点的通信
- 从节点如果没有收到主节点的心跳，则进入备选状态
- 主节点或者从节点发现自己的期限过期，则进入随从状态

### 5.2 主节点选举
- 主节点发送AppendEntries给从节点
- 从节点如果没收到心跳，则发起选举
- 如果开始选举，从节点则增加自己的期限，发送RequestVote给其他所有节点
- 结果有三种：赢得选举，其他节点选举成功，超时
- 如果在备选期收到AppendEntries心跳，且期限数大于等于自己的，该节点则返回从节点状态
- 如果期限数小于自己的，则继续停留在备选状态
- 如何避免split votes分票？
- 如何判断日志提交成功？
  
## Terminology    
identical,	完全相同的，  
fault tolerance	,	容错  
single cluster leader	,	单集群领导  
a separate replicated state machine	,	一个单独的复制状态机  
manage leader election	,	管理领导人选举  
survive leader crash	,	在领导崩溃中幸存下来  
a series of commands	,	一系列命令  
state machine	,	状态机  
execute in order	,	按顺序执行  
the same sequence of commands	,	相同的命令序列  
deterministic	,	确定性的  
replicated log consistency	,	复制日志一致性  
even if	,	即使  
non-Byzantine condition	,	非拜占庭条件  
fully functional	,	功能齐全  
tolerate the failure	,	容忍失败  
depend on timing	,	取决于时机  
faulty clocks	,	有故障的时钟  
majority of cluster	,	大多数集群  
single round	,	单轮  
remote procedure call	,	远程过程调用  
impact	,	影响  
synonymous with	,	与...同义  
capable of	,	能够  
single decree	,	单一法令  
facilitate a series of decision	,	促进一系列决策  
cluster membership	,	集群成员  
two significant dwawbacks	,	两个重要的dwawback  
notoriousy opaque	,	臭名昭著的不透明  
single decree subset	,	单一法令子集  
simple intuitive	,	简单直观  
develop intuition	,	培养直觉  
single decree protocol	,	单一法令协议  
composition rules	,	组成规则  
complexity and subtlety	,	复杂性和微妙性  
multiple decision	,	多重决策  
decomposed in	,	分解在  
agreed-upon algorithm	,	共识算法  
sketched	,	草图  
flesh out	,	充实  
in most cases	,	在大多数情况下  
single decree decomposition	,	单一法令分解  
melding them into	,	将它们融合成  
constrained order	,	受约束的顺序  
elect a leader	,	选举领袖  
bear little resemblance	,	没有几分相似  
error prone	,	容易出错  
formulation	,	公式  
priving theorems	,	私有定理  
importance	,	重要性  
alternative	,	选择  
practical foundation	,	实践基础  
a large audience	,	大量观众  
intuition	,	直觉  
evitable	,	不可避免的  
evit	,	维特  
evict	,	驱逐  
numerous points	,	无数点  
evaluated the alternatives	,	评估替代方案  
state space	,	状态空间  
subtle implications	,	微妙的含义  
recognize	,	认出  
a high degree of subjectivity	,	高度的主观性  
well known approach	,	众所周知的方法  
problem decomposition	,	问题分解  
simplify the state space	,	简化状态空间  
coherent and eliminating non-determinism	,	连贯和消除不确定性  
not allowed to have holes	,	不允许有孔  
inconsistent	,	不一致  
eliminate nonterminism	,	消除非终端主义  
randomized approaches	,	随机方法  
tend to reduce	,	倾向于减少  
key properties	,	关键属性  
piecewise	,	逐段地  
electing a distinguished leader	,	选出一位杰出的领袖  
consulting	,	咨询  
decompose	,	分解  
TODO, P4
  
  
  
  
