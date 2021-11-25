# Ch1: System Design Tips and Introduction

需要掌握的程度：

Not really required for new grads:

- Focus more on fundamental. SQL vs noSQL, ACID, etc.

SDE 2, ask hint, no major red flags to pass the interview.

- Okay not to know what is rainbow table but should not store the pwd in plaintext
  - Rainbow table: precomputed table for the output of cryptographic hash functions, usually for cracking password hashes.

Senior or above. 一票否决

- Red Flag 例子：Use database to store multi-media file. 用object storage
- Red Flag 例子：Cassandra / redis for bank transaction. 用Relational database。



Tips:

1. Build blocks and modules 



### Truth, Knowledge and Lies

- You should know which design does NOT work fundamentally



超算中心: 

- One component fail will halt the whole progress.

  

Commercial DC - Clos topology 🐶

- 每个component都很便宜

- Built a reliable system from unreliable components: TCP on top of unreliable network

- Error correcting code for corrupted data



Network failures

- Indicator: timeout. If failed, retry few times to confirm.
  - 确定Retry 相隔的时间：做实验测一下或者用已有算法
- Falsely declared dead => split brain
  - HDFS用ZooKeeper 保证只有一个Node active。Fencing mechanism迫使之前active的node不能访问resources 
  - Kafaka用Epoch number
  - Cassandra 用generation number去记录node的restart次数
- Idempotent 
  - 一个请求发了很多次，结果应该不变 比如 GET, PUT, DELETE （但POST不）



Network delays/lost

- Delay is chosen by design （因为是为了提高resource利用率）
- 设计的是maximize throughput

- Real time OS 更关注 response time



Clock

- Cluster中常用 Network Time Protocol (NTP)。根据 A group of servers 调整时钟
- GPS receiver 取时间更准

- 时间衡量
  - Time-of-day clock:  `clock_gettime(CLOCK_REALTIME)`: highly unreliable.
  - Monotonic clock: `clock_gettime(CLOCK_MONOTONIC)`: get duration. relatively accurate in distributed system

- Spanner: Built on top of atomic and GPS clock
  - 好处：可以用来判断events的先后顺序 （例如 处理 Last write win 数据库的数据不一致问题）



Process Pause

- "stop-the-world" garbage collection
- OS context switch; hypervisor switch to a different VM.
- wait for slow IO when applying synchronous disk access
- Thrashing: swapping to disk
- Real time OS: every library call has a worst time guaranteed

推荐的reading：

- How to do distirbuted locking: http://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html
- Redlock: fault-tolerant distributed locks : no fencing token



和Process Pause相关的问题（HBase 出现过）

- Client A 和 Client B 要往同一个storage 写data，预期先A后B。Client A先拿到lease，但被stop-the-world GC pause打断，lease过期了还没写。Client B拿到了lease，写完data B。Client A 做完GC继续写data A，corrupt the file in storage.

解决方法：用fencing token，get lease的时候有一个递增的token，写的时候check token。Token相比前一个operation的小了，就reject这个操作。



Truth:

- Majority defines the truth: Quorum, more than half. 
- Majority里的disagree, need consensus.



Lies:

- Byzantine failure



Different models in distributed system

- Synchronous: 不现实，因为 unbounded delays and pause 真实存在。
- Partially synchronous: A realistic model of many system. 大部分时候synchronous，有时候delay很大。
- Asynchronous: 不能用任何timing assumption。clock不能用。concurrency control 里最高的隔离级别SERIALIZABLE 没法实现。



有关Faults的assumption

- Crash-stop faults：挂了就gone forever
- Crash-recovery faults: 挂了存在disk里的还在，in-memory的没了。perhaps start responding again after some unknown time.
- Byzantine (arbitrary faults): Nodes do absolutely anything. 在DC里一般没问题，P2P网络里可能有。
