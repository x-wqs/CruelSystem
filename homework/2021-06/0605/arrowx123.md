# Raft Paper

Highlights:
- More understandable than Paxos;
	- problem decomposition: leader election, log replication, safety, and membership changes
	- stronger degree of coherency -> less number of states
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

# Raft 

## Guarantee
- Election Safety: at most one leader
- Leader Append-Only
- Log Matching: if logs contain an entry with the same index and term, all previous entries are identical
- Leader Completeness: if entry commited in a term, all higher terms in leader logs contain that entry
- State Machine Safety: log entry at a given index does not conflict


## Summary (quick reference)
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
		- If election timeout elapses w/o AppendEntries RPC or granting vote to candidate: convert to candidate
	- Candidates
		- start election
			- increment currentTerm
			- Vote for self
			- Reset election timer
			- Send RequestVote to all others
		- votes received from majority of servers: become leader
		- AppendEntries RPC received from new leader: convert to follower
		- election timeout elapses: start new election
	- Leaders
		- Upon election
			- send initial empty AppendEntries
			- repeat during idle periods
		- command received from client
			- append entry to local log
			- respond after entry applied to state machine
		- If last log index ≥ nextIndex for a follower
			- send AppendEntries starting at nextIndex
				- successful: update nextIndex and matchIndex for follower
				- fails because of log inconsistency: decrement nextIndex and retry
		- exist N, if (N > commitIndex) && (majority of matchIndex[i] >= N) && (log[N].term == currentTerm), commitIndex = N

## Basics

Algorithm
1. electing a distinguished leader
	- accepts log entries from clients
	- replicates them on other servers
	- tells servers when it is safe to apply log entries to their state machines
2. Submodule
	- Leader election
	- Log replication
	- Safety

Server states:
- leader: exactly one
- follower: only respond to requests
- candidate: used to elect a new leader
Term:
- arbitrary length
- numbered with consecutive integers
- begin with election
- at most one leader in a given term
- act as a logical clock
- increases monotonically over time

Leader election:
- Leaders send periodic heartbeats to all followers
- If election timeout, begins an election
- election
	- follower increments its current term and transitions to candidate
	- votes for itself
	- issues RequestVote RPCs to other servers
	- consequences
		- (a) it wins the election
			- receives votes from a majority of the servers
			- Each server will vote for at most one candidate
		- (b) another server establishes itself as leader
		- (c) a period of time goes by with no winner
	- when waiting for votes and received a heartbeat claiming to be leader from another server
		- leader’s term >= candidate’s current term: return to follower
		- leader’s term < candidate’s current term: reject the leader, and continue in candidate
- split votes issue could happen indefinitely (w/o extra measures)
	- randomized election timeouts

Log Replication
- client req contains a command
- leader appends the command (entry) to its log
- leader applies the entry and returns the result to clients, after it hs been replicated
- leader keeps retrying AppendEntries indefinitely until all followers store all log entries
