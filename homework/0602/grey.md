Note: this includes both 0601 and 0602 homework.

# Raft Paper

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
- does not provide a good foundation for building practical implementations
-- details are missing for building multi-Paxos from single-decree Paxos
- architecture needs complex changes to support practical use
-- little benefit melding a collection of log entries into sequential log
-- uses a symmetric peer-to-peer approach

## Raft 
Properties
- understandability
-- problem decomposition: leader election, log replication, safety, and membership changes
-- simplify state space

Algorithm
1. electing a distinguished leader
	- accepts log entries from clients
	- replicates them on other servers
	- tells servers when it is safe to apply log entries to their state machines
2. Submodule
	- Leader election
	- Log replication
	- Safety

Components
- state
	 - persistent state (persisted in stable storage)
		 - currentTerm
		 - votedFor
		 - log[]
	- volatile state on servers
		- commitIndex
		- lastApplied
	- volatile state on leaders (reinitialized after election)
		- nextIndex[]
		- matchIndex[]
- appendEntries RPC (replicate log entries; heartbeat)
	- Arguments
		- term
		- leaderId
		- prevLogIndex
		- prevLogTerm
		- entries[]
		- leaderCommit
	- Results
		- term
		- success
	- Receiver Implementation
		- false if term < currentTerm
		- false if log doesn't contain an entry at prevLogIndex of prevLogTerm
		- delete existing entry and all the following if existing entry conflicts with a new one
		- append new entries not in the log
		- if leaderCommit > commitIndex, commitIndex = min(leaderCommit, index of last new entry)
- RequestVote RPC (gather votes)
	- Argumetns
		- term
		- candidateId
		- lastLogIndex
		- lastLogTerm
	- Results
		- term
		- voteGranted
	- Receiver Implementation
		- false if term < currentTerm
		- votedFor is null or candidateId, and candidate’s log is at least as up-to-date as receiver’s log, grant vote
- Rules for Servers
	- All Servers
		-  commitIndex > lastApplied: increment lastApplied, apply log[lastApplied] to state machine
		- If term T > currentTerm in RPC request or response: currentTerm = T, convert to follower
	- Followers
		- Respond to RPCs from candidates and leaders
