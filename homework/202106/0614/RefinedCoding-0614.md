# 6.824 Raft FAQ  
  
## Simplicity  
- log for persistency  
- follower reject AppendEntries out of date  
- with snapshot  
  
## Senarios  
- Docker  
- etcd  
- MongoDB  
  
## Paxos  
- invented in the latest 1980s, early 1990s  
- no leader  
- make servers to agree on a single value  
- far simpler than Raft, solves less problems than Raft  
- replicate need to agree on an indefinite sequence of values  
- ViewStamped Replication in 1988  
- an agreement canbe started by any replica, how to guareetee its correctness  
  
  
## Raft  
- invnted in 2012  
  
## Others  
群友在做Lab1的时候有没有遇到Connection Refused的错误啊，我拉MIT的库啥东西没改也有这个问题，DDIA群主建议我们去做6.824 2020的作业  
嗯嗯，谢谢群友，还是群友学习仔细，Log里面是不是有term和index啊？term和index都要比较么？  
  
## Terminology  
flavor	,	味道  
needlessly	,	不必要地  
tutorial	,	教程  
provision	,	条款  
pipelining	,	流水线  
mechanism	,	机制  
in flight	,	飞行中  
batching	,	分批  
in return for	,	作为回报  
sacrifice	,	牺牲  
impression	,	印象  
indefinite	,	不定  
introductory	,	介绍性的  
in contrast	,	相比之下  
efficiently	,	有效率的  
complex	,	复杂的  
introductory	,	介绍性的  
in contrast	,	相比之下  
relatively	,	相对地  
fairly	,	相当  
contribution	,	贡献  
inherently	,	天生的  
sacrifice	,	牺牲  
resemble	,	类似  
more or less	,	或多或少  
derived	,	衍生的  
satisfactory	,	满意的  
diverging	,	发散  
mutate	,	变异  
  
  
  
  