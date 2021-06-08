# 6.824 Lecture 6 - Raft 1

## Lab
- 2A: Raft leader election
- 2B: Raft log handling/replication

## Fault Tolerance Patterns
- MapReduce, computation replication and need coordinator to organize
- GFS, data replication, need coordinator to pick primaries
- FT VM, service replication, need test-and-set to pick primary
- all need coordinator
- advantage: avoid split brain

## Split Brain
- primary conflict
- network partition

## Majority Vote
- odd number of servers
- need majorities agree, e.g 3/5
- because at most one majority
- 2f + 1 servers can tolerate f failed servers
- become quorum systems with more than f failed servers
- two majorities must overlap/intersect
- successive majority can convery old one

## Paxos
- invented around 1990 with ViewStamped replication
- used in real word in the last 15 years
- Raft for modern techniques
- Raft vs. Paxos: https://dl.acm.org/doi/10.1145/3293611.3331595

## Raft
- state machine replication
- client -> leader: send command
- leader -> follower: AppendEntries RPCs
- follower add command to log
- leader waits for reply from majority
- entry is committed when majority reply
- leader execute command, replies to client (how about crash here)
- follow execute command

## Log
- with state machine replication
- replicas with a single execution order
- leader sures followers with identical logs
- tentative commands are stored in log until committed
- commands are stored in log for deduplicate
- commands are stored in log for persistency after reboot

## Lab2
- Command(index, term, isLeader)
- Raft.Start(command)
- Leader.Send(AppendEntries) to Followers
- Leader Election
- Ensure identical logs

## Lab2A
- new leader increments term
- at most one leader per term, i.e. may no leader
- how to become a candidate

## Question
- 群友有5台服务器，1个Leader S1，四个Follower ， S2 - S5
client 发送数据D1给 Leader ，S2 和 S3回复了，Leader S1认为数据提交了
然后S1和S2宕机了，S4被选为新的Leader
请问D1如何从Follower S3同步到新的Leader S4和另一个Follower S5啊，谢谢了！
Majority可以包括自己吧，S3, S4, S5
五台机器，宕机了两台，会继续工作吧
S4有可能选为Leader吧

## Terminology
arise	,	出现
nasty	,	讨厌
despite	,	尽管
symptom	,	症状
network partition	,	网络分区
insurmountable	,	不可逾越
insight	,	洞察力
majority	,	多数
symmetry	,	对称
tolerate	,	容忍
quorum	,	法定人数
convey	,	传达
piggybacks	,	背驮
superseded	,	被取代
cast one vote	,	投一票
suppress	,	压制
precedence	,	优先级
randomness breaks symmetry	,	随机性破坏对称性
needless	,	不必要
sacrifice	,	牺牲
in return for clarity	,	以换取清晰
















