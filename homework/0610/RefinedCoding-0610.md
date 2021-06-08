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
flavor
needlessly
tutorial
provision
pipelining
mechanism
in flight
batching
in return for
sacrifice
impression
indefinite
introductory
in contrast
efficiently
complex
introductory
in contrast
relatively
fairly
contribution
inherently
sacrifice
resemble
more or less
derived
satisfactory
diverging
mutate


