## source
https://pdos.csail.mit.edu/6.824/papers/gfs.pdf
9 ~ 15

## fault tolerance and diagnosis
### high availability
- fast recovery, both for chunkserver and master

- chunk replication

- master replication
> when master fails, monitoring infrastructure outside GFS starts a new master process with replicated operation log

### data integrity
> each chunkserver must independently verify integrity of own copy by maintaining checksums

> each 64KB block has 32 bit checksum, kept in memory and persistent by log

> chunkserver check itself to avoid propagate corruptions to clients or other chunkservers

> chunkservers scan checksums of inactive chunks during idle periods.

### diagnostic tools
> GFS servers generate diagnostic logs

## measurements
### micro benchmarks
#### read
> efficiency drops from 80% to 75% when readers increases from 1 to 16, because of increasing probability of multiple readers simultaneously read from same chunkserver

#### write
> collision is more likely for multiple writers than multiple readers, as write involves three replicas

#### record appends
> performance drops due to congestion

### real world clusters
> read rates much higher than write rates

> master load is not bottlenect.(After a optimization of binary searches through namespace rather than scanning through large directories for a particular file)






## terms
integrity 完整性

divergent 不同的

propagate 传播

idle 空闲

transient 短暂的

saturate 饱和

culprit 祸首

congestion 拥塞

hobble 蹒跚

illegible 难以辨认

sustaining 维持

leeway 退路

caveat 注意事项

infer 推断

cumbersome 麻烦的

mutual 相互的

bimodal 双峰的

outpace 超越

portion 部分

versus 相对

skew 歪斜

deliberate 故意的

automated 自动的

conceive 构想

rudimentary 初级的

sporadic 零星的

aggregate 总计的

opt 选择

magnitude 震级

anticipated 预期的

IDE integrated development environment

delegate 代表

lease 令牌

lift 举起

innovate 创新

shepherd 牧羊人
