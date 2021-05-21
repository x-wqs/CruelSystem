## source
https://pdos.csail.mit.edu/6.824/papers/gfs.pdf
5 ~ 9

## system interactions
designed to minimize master's involvement

### lease and mutation order

> clients query master for lease and locations of replicas

> master reply

> clients push data to all replicas

> clients send write request to primary

> primary forwards write request to all secondary replicas

> secondary reply to primary indicating completed

> primary reply to client

### data flow
> pass data by chain, choosing nearest replications

> B / T + RL, T is network throughput, typically 100Mbps

### atomic record appends
> not use offset as traditional, always append to end

> guarantee data is written at least once as an atomic unit

### snapshot
> copy on write to implement snapshot

## master operation
### namespace management and locking
> read lock on /d1, /d1/d2, write lock on /d1/d2/leaf

### replica placement
> tradeoff between reliablility and bandwidth(multiple rack)

### creation, re-replication, rebalancing
#### creation
> create on below-average disk space utilization

> limit number of recent creation on each chunkserver, avoid imminent heavy write traffic

> spread replicas of a chunk across racks

#### re-replication
> do when number of available replicas fall below a user-specified goal

- reasons
> chunkserver avalible

> replication goal increased

- priority factors
> delta to goal

> prefer chunks for live files than recently deleted files

> boost priority of chunk blocking client progress

- clone destination
> goals are same as creation

> limit numbers of active clones

> limit bandwidth of each clone

#### rebalance
> same goals as creation

> gradually fill up new chunkserver

> remove beyand-average disk space utilization

### garbage collection
> after file deletion, gfs reclaim avalible physical storage lazily during regular garbage collection at both file and chunk levels

#### mechanism
- file lazy delete
> rename file to hidden name including timestamp, deleted during master's scan if exceed TTL

- chunk lazy delete
> identify orphaned chunks during master's scan, erase metadata for those chunk. when regular heartbeat, reply erased meta compared to each chunkserver's report

#### discussion
- pros
> First, simple and reliable, fault tolerant

> Second, done in batch and cost amortized, merge storage reclamation into regular background activities of master, done only when master is relatively free

> Third, delay provides safety against accidental and irreversible deletion

- cons
> hinder user's effort to fine tune usage, such as frequently delete and create temporary files

- patch for cons
> expedite storage reclamation if deleted file explicitly deleted again

> user-defined replication and reclamation policies

### stale replica detection
> master maintains chunk version number to distinguish between up-to-data and stale replicas

> master removes stale replicas in regular garbage collection

> master reply to clients or cloning chunkserver version number, thus they perform check operations

## terms
lease 约定

revoke 撤销

decouple 解耦

straddles 跨越

duplex 双工

congestion 拥塞

rack 架子

imminent 即将来临

throttle 节流

amortize 摊销

expedite 加快

hinder 阻碍

diagnosis 诊断

inevitably 不可避免地

minor 次要的

canonical 典范

lag 法律

poll 轮询
