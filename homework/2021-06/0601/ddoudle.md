Raft Paper (P1-3)	

Consensus algorithm

​	allow a collection of machines to work as a group that can survive the failures of some members.

​	job: keep the replicated log consistent

​	replicated state machines: implemented using a replicated log, to solve fault tolerance problem

​	safety, fully functional, not depend on timing, 

Paxos: 

​	ensure both safety and liveness, support changes in cluster membership, correct, efficient.

​	difficult to understand, not provide a good foundation for building practical systems

​	not provide 

Raft: 

​	provide a complete and practical foundation

​	safe, efficient

​	understandability

​		problem decomposition

​		reduce the number of states

​	novel features: strong leader (simplify the management), leader election (randomized times to resolve conflict), membership changes (new joint consensus)

​	