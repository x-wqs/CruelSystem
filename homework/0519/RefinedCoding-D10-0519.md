# RefinedCoding-D09-018  
- GFS P6 - P10  
- https://pdos.csail.mit.edu/6.824/papers/gfs.pdf  
  
## 2. Design Overview 概要设计  
  
### 2.7 Consistency Model 一致性模型  
- Guarantees by GFS / GFS约束  
- Implication for Applications 应用蕴含  
  
## 3. Sysetm Interfactions 系统交互  
  
### 3.1 Leases & Mutation Order 租约和变化顺序  
- GFS 3.1.3 The client pushes the data to all the replicas.   
- 一个客户端要写多个节点复制，会不会增加客户端的磁盘读取和网络负担啊?  
- 为什么不让GFS内部实现同步呢？在GFS图2里面也没看到客户端写入Secondary Relica B啊  
  
### 3.2 Data Flow 数据流  
  
### 3.3 Atomic Record Appends 原子性的记录追加  
  
### 3.4 Snapshot 快照  
  
## 4 Master Operation 集群主操作  
  
### 4.1 名空间管理和锁  
  
### 4.2 副本布置  
  
### 4.3 块的创建，重复制，和重负载均衡  
  
### 4.4 垃圾回收  
  
### 4.5 稳定副本的探测  
  
### 5 容错和诊断  
  
### 5.1 高可用性  
  
### 5.2 数据的完整  
  
### 5.3 诊断工具  
  
## 6 测试  
  
### 6.1 微基准  
  
### 6.2 现实的集群  
  
## Terminology  
distinguish	,	区分  
earliest opportunity.	,	最早的机会。  
insert padding	,	插入填充  
drawfed by	,	拉拢  
chunk version number	,	块版本号  
cache entry's timeout	,	缓存条目的超时  
premature	,	过早的  
handshakes	,	握手  
data corruption	,	资料损坏  
iirreversibly	,	不可逆转地  
accomodate the relaxed consistency model	,	适应宽松的一致性模型  
relying on	,	靠  
aotomically	,	原子地  
permanent	,	永恒的  
regardless	,	而不管  
far more	,	还有更多  
resilient	,	有弹性的  
perspective	,	看法  
prospective	,	预期  
aspect	,	方面  
discard extra paddinggragments	,	丢弃多余的填充格言  
occasional	,	偶然  
lease	,	租  
apply mutation	,	应用突变  
grant order	,	授予命令  
lease mechanism	,	租赁机制  
piggy	,	小猪  
revoke a lease	,	撤销租约  
identity	,	身份  
decoupling	,	去耦  
identify	,	确认  
indicate	,	表明  
straddle	,	跨骑  
interleave	,	交错  
identical	,	完全相同的  
identity	,	身份  
identify	,	确认  
decouple	,	解耦  
distributed lock manager	,	分布式锁管理器  
traditional write	,	传统写法  
instantaneously	,	瞬间  
stardard copy-on-write	,	标准型写时复制  
interruption	,	打断  
revoke lease	,	撤销租赁  
hence	,	因此  
reclaim	,	回收  
proper serialization	,	适当的序列化  
prefix compression	,	前缀压缩  
illustrate	,	阐明  
suffices	,	足够  
racks	,	架子  
scalability	,	可扩展性  
reliability	,	可靠性  
avialability	,	可用性  
circuit	,	电路  
exploit	,	开发  
survive	,	存活  
utilization	,	利用率  
imminent	,	即将来临  
predict	,	预测  
reliably	,	可靠地  
prioritize	,	优先  
equalize	,	均衡  
overshelming	,	压倒性的  
additionally	,	此外  
throttling	,	节流  
gradually	,	逐步地  
swamp	,	沼泽  
criteria	,	标准  
reclaim	,	回收  
undelete	,	取消删除  
reclamation	,	开垦  
uniform	,	制服  
amortize	,	缓冲  
spond	,	自发的  
irreversible deletion	,	不可逆删除  
hinder	,	阻碍  
storage is tight	,	储存空间紧张  
tune usage	,	调音用法  
expediting	,	加急  
explicitly	,	明确地  
distinguish	,	区分  
safeguard	,	保障  
bound to be	,	一定会  
routinely	,	例行地  
challenging	,	具有挑战性的  
manageable	,	可管理的  
complicated	,	复杂  
loosely coupled system	,	松耦合系统  
dominated	,	占主导地位  
elsewhere	,	别处  
canonical name	,	规范名称  
DNS alias	,	DNS别名  
lag the primary slightly	,	稍微落后于小学  
fractions of a second	,	几分之一秒  
shadow master	,	影子大师  
checksumming	,	校验和  
corruption	,	腐败  
moreover	,	而且  
divergent	,	发散的  
legal	,	合法的  
  
  
  
  
  
  
  
  