# Raft

Highlights:
- More understandable than Paxos;
- Structured different
-- separate leader election, log replication, safety, stronger degree of coherency
- use overlapping majority for changing cluster membership


Raft Design Philosophy:
- understandability
-- decomposition
-- state space reduction
- facilitate intuition

Raft Features:
- Strong leader: log entries only flow from leader
- Leader election: randomized timers to elect leaders, implemented on top of heartbeats
- Membership changes: joint concensus

# Intro
## Replicated state machines
Consensus algorithms typical properties:
- safety: never returning an incorrect result
- high available if > 1/2 servers are operational
- do not depend on timing for consistency
- minority of slow servers need not impact performance

## Paxos 
Properties
- single-decree Paxos: reaching agreement on a single decision
- combine multiple single-decree Paxos -> multi-Paxos

Cons
- difficult to understand
- architecture needs complex changes to support practical use
