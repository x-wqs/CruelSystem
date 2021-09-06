# # Lecture 5: Fault Tolerance: Raft (1)

## Pattern
Pattern: single point of failure
- map reduce: coordinator 
- gfs: master
- ft-vm: storage server
single machine (instead of multiple machines) -> avoid split brain

Idea: test-and-set server replication
possible causes:
- server fails
- network partition

Network partition: majority rule
- example: s1, s2, s3
- overlapping in majorities
- 2f + 1 (odd number of servers)
- majority: live + down servers

Protocols using quorum:
- ~1990
	- Paxos
	- View-stamped replication
- ~2014
	- Raft
		- rsm (replicated state machine) with raft

K/V servers implemented with Raft
- append put/get operations into logs
- on failure:
	- new leader election
	- detect duplicates

Overview
- appendEntry
Why logs?
- retransmission
- order
- persistence
- space for tentative commands
- logs should be identical on all servers
Log entry
- comand
- term
	- leader's term: term id implicitly signals who is the leader
- log index

Elect leader
- missing heartbeats (appendEntry) from leader -- (election timeout) --> start leader election
	- increase term number
	- vote for itself
	- requestVote from other servers
	- no split brain
		- solutions:
			- majority rule
			- term number
- challenge
	- split vote
		- solution: election timeout randomized
	- election timeouts:
		- >= few hearbeats period
		- random valve
		- short enough that downtime is short
Ensure logs identical
