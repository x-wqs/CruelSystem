# 6.824 Distribute System  
- what is the distributed system  
- how did distributed system develop in last decades  
- main topics / course structure  
- 1st lab: map reduce  
  
## Goal of distributed system  
- availabity / fault-tolerance  
- scalability  
- performance  
- security / isolation  
  
## Topics  
- fault-tolerance  
- strong/weak/causal consistency  
- performance  
  
## History  
- 1980 MIT, LAN  
- 1990 Internet, Yahoo  
- 2000 Big Data, Google  
- 2010 Clouding Compution, Amazon  
- 2020 TODO  
  
## Papers  
- MapReduce    
- Raft    
- ZooKeeper    
- Chain Replication    
- 6.033 Text Book Ch9. Atomicity  
- Spanner    
- Distributed Transaction    
- Spark    
- MemCached    
- Fork Consistency    
- BitCoin    
- BlockStack    
- AnalogicFS    
  
## Questions  
1. However, MapReduce limits what apps can do:   
No real-time or streaming processing.  
Q: Map Tasks should be finished before Reduce?  
为啥不支持流啊，不能持续把input.txt写入GFS，然后启动Map Task， Reduce Task 么？  
  
2. Paper's root switch: 100 to 200 gigabits/second, total 1800 machines,     
so 55 megabits/second/machine. 55 is small, e.g. much less than disk or RAM speed.    
A: 这是路由或者交换机，不是Hub，每两台机器之间的速度都应该是200Gbps吧     
  
3. after all Maps have finished, coordinator hands out Reduce tasks each Reduce fetches its intermediate output from (all) Map workers.    
Q: How does reducer read data from mapper?    
  
4. Reduce workers read directly from Map workers, not via GFS.    
Q: Which protocol?  
  
5. Can omit re-running if Reduces already fetched the intermediate data ? 
A: No, all mapper tasks are finished.
  
6: What if the coordinator crashes?  
  
7. Probably no longer in use at Google.  
Q: Not the most efficient or flexible. Replaced by Flume / FlumeJava (see paper by Chambers et al).  
  
## Teminology  
spectrum	,	光谱  
pure deterministic functions	,	纯确定性函数  
interaction	,	相互作用  
straggler	,	散客  
flakey hardware	,	亚麻五金  
Hugely influential, 极具影响力  
  
  
   