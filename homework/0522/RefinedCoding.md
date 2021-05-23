# Fault-Tolerant VM            
- P1 - P3            
- https://github.com/refinedcoding/CruelSystem/blob/main/paper/vm-ft.pdf            
            
## 0.概要            
- enterprise grade 企业级的            
- 在VMware vSphere 4.0中已实现            
- 在廉价服务器上运行，FT VM只影响10%的性能            
- 主备机器之前带宽只有20 Mbps，可以实现远距离部署            
            
## 1. Introduction            
- 重现状态的方法会造成很大的带宽使用            
- 确定状态机            
- 忽略不确定的操作，比如读取时钟，截取中断            
- 状态机方式适合用在虚拟机上            
            
## 2. 基本的容错设计            
            
### 2.1 确定性恢复实现            
            
### 2.2 容错协议            
- 输出需求：如果备份机器能够在主服务器宕机接管的话，那它就必须和之前主服务发往外界的所有输出保持一致。            
- 输出规则：主服务器可能向外界发送输出信号，知道备份机器收到和确认于产生输出的操作相关的日志。            
            
## Questions            
- FT VM比GFS生词少一些，不查字典，也知道中文是啥意思，但是就是不知道作者在说啥，感觉Primary VM和Backup VM就是数据库Replica一样，用WAL同步，然后ACK确认，加上原来VMWare vSphere共享内存的架构，同步基本能在一秒内完成。             
- 大家可以推荐一下WAL的论文么？是不是业界现在都用这个啊，和其他同步的方式有什么优劣啊            
            
            
## Termilogy            
deterministic，确定性的；命运注定论的            
practical            
lockstep            
replicate the execution            
take over            
identical            
deterministic state machine            
corordination            
hypervisor            
in contrast            
well defined            
            
            
            
            