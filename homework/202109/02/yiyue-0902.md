## Related work

- Consensus algorithms publication categories
  - Original Paxos
  - Elaborations of Paxos
  - Systems that implement concensus algorithms
  - Performance optimizations applied to Paxos
  - Oki & Liskov's Viewstamped Replication
- Difference between Raft and Paxos - Strong leadership
  - Raft uses leader election as part of consencus protocol, concentrates much functionality in the leader
  - Paxos uses leader election only to improve performance but not to achieve consensus
  - Paxos includes 2-phase protocol for consensus
  - Raft incorporates leader election into consensus algorithm
- VR and ZooKeeper are leader-based
- Raft has less mechanism because it minimizes functionality in non-leaders
  - log entries flow only one direction: leader -> AppendEntries
  - VR flows both directions
  - ZooKeeper log entries <-> leader, more like Raft
- Raft has fewer message types than any other consesus-based log replication algorithm - only 4 message types (two PRC requests & corresponding responses)
- Raft's strong leadership precludes some performance optimizations
- Raft has joint censensus approach - leverages the rest of the consensus protocol to have little additional mechanism for membership changes
- Lamport's alpha-based approach wasn't option for Raft - it assumes consensus reached without a leader
- Reconfiguration algorithm has advantage - membership changes without limiting normal requests
  - VR stops all normal requests
  - SMART imposes alpha-like limit
- Raft adds less mechanism

## Conclusion
- Understanding is as important as correctness, efficiency, and conciseness for algorithms
