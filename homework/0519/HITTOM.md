## source
https://pdos.csail.mit.edu/6.824/papers/gfs.pdf
5 ~ 10

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

##terms
lease 约定
revoke 撤销
decouple 解耦
straddles 跨越
duplex 双工
congestion 拥塞
