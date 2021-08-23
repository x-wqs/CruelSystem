Design: a fault-tolerant key/value storage service with Raft library

requirements:
1. the service should support 3 operations: put(key, value), append(key, arg), get(key)
2. the service provides strong consistency calls to the clerk methods
concurrent calls: the return values and final state must be the same as if the operations had executed one at a time in some order, a call must observe the effects of all calls that have completed before the call starts (linerizability)
3. achieven the consistency for the service which is replicated, since all servers must choose the same execution order for concurrent requests, and must avoid replying to clients using state that isn't up to date

The retults of the consistency:
all clients see the same state and they all see the same last state


Two parts of tasks:
A: implement the service without the raft log
No dropped message, no failed servers
-- The clerks send put(), append() and get() rpcs to kvserver and kserver submits them to raft
-- the kvservers should not directly communicate and they should only interact with each other -- through raft
-- The servers will not execute the request twice with the re-send


B: implement the snapshots, which allow raft to discard the old entries
Whenever your key/value server detects that the raft state size is approaching the threshold, it should have a snapshot using snapshot

