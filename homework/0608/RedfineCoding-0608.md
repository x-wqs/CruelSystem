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


## Question
- 群友有5台服务器，1个Leader S1，四个Follower ， S2 - S5
client 发送数据D1给 Leader ，S2 和 S3回复了，Leader S1认为数据提交了
然后S1和S2宕机了，S4被选为新的Leader
请问D1如何从Follower S3同步到新的Leader S4和另一个Follower S5啊，谢谢了！

## Terminology
arise
nasty
despite
symptom
network partition
insurmountable
insight
majority
symmetry
tolerate
quorum
convey
piggybacks

