Raft(section 1- 4)
Summary
Paxos is the acronym of consensus algorithm but it has drawbacks: 1)hard to understand 2)hard to implement, in other words
its lacks good foundation for building large/complex software system. Raft achieves the comparable performance but improves
the understandability and introduces new thoery foundation to make it easy to implement. 
Raft has 1)strong leaders to reduce the non-deterministic of Paxos  2)randomized timers to elect leaders 3)use a new joint 
consensus approach for consensus so that cluseter can continue to work when config changes

Replicates state machines are the foundations to fault-tolerance. Consensus algorithms  should 1)ensure safty 2)continue to
function as any majority servers are operational 2)don't depend on timing to ensure the consistency 3) a command is completed
when any majority servers responded to a PRC call, minority of the servers will not impact the final result.

Raft improves Paxos in aspects of 1)undertandability and 2)foundations for building complex software. Specifically, it 
decomposes the problem into several parts 1)leader election 2)log replication 3)safty 4)membership changes. Also Raft reduces
the number of states to eliminate nondeterminism as much as possible, although randomization is used in the leader election to simplify
the leader election. The final result is that study shows people tend to understand Raft better than Paxos and there are 
more implementations of Raft than Paxos which is error-prone and different among different implementations.
