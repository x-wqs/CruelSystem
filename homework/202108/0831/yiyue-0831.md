## 8 Client interaction

- client connects to a random server
- if server is not leader, rejects the request and sends info of most recent leader
- if leader crashes, client requests timeout
- to ensure linearizable semantics
  - clients assign unique serial numbers to every command
  - state machine tracks the latest serial number & response
  - if command's serial number executed, responds without re-executing
- to ensure no stale data returned for read-only operations
  - stale data happens when a leader responding to a message is already superceded by another leader but being unaware
  - Leader cCompleteness Property guarantees a leader has all the committed entries. 
  - each leader commits a blank no-op entry at the start of term to find out those committed entries
  - leader should check whether it has been deposed before processing a read-ony request - done by leader exchange heartbeat messages with a majority of cluster

## 9 Implementation and evaluation

- An experiment is conducted to compare Raft and Paxos, and found Raft easy to understand and implement
- Correctness is proved
- Performance
  - established leader replicates new log entries using the minimal number of messages
  - performance is measured by crashing the server uniformly at random
  - downtime can be reduced by reducing the election timeout
  - cannot further lower the timeout because should achieve system availability