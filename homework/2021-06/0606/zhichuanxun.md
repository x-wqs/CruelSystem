# Raft

## Abstract

raft is a consensus algorithm for managing a replicated log

## Introduction

Raft is similar with many existing consensus algorithms
novel features:

- strong leader
- leader election: use random timer to elect leader
- membership change

## Replicated state machines

state machines on a collection of servers compute identical copies of the same state

Replicated state machines are used to solve FT problems

Replicated state machines are typically implemented using a replicated log

The consensus algorithm is to keep the replicated log consistent

Properties:

- ensure safety under all non-Byzantine[TODO] conditions(network delays, partitions, packet loss, duplication and recording)
- majority of servers are operational and can communicate with each other
- consistency do not depend on timing
- a command can complete as soon as majority of the cluster has responded to a single round of rpc

## Paxos

- Paxos is exceptionally difficult to understand
- Paxos does not provide a good foundation for building practical implementations(implementation will be changed due to real-world situation, and the final system will be based on an unproven protocol)

## Designing for understandability

- problem decomposition: divide problems into separate pieces that could be solved, explained and understood relatively independently(leader election, log replication, safety and membership changes)
- simplify the state space by reducing the number of states to consider(logs are not allowed to have holes)

## The Raft consensus algorithm

Raft use a distinguished leader to manage all logs

consensus problem is decomposed into 3 questions:

- Leader election
- Log replication: leader must accept log entries and replicate them across the cluster
- Safety: the safety in state machine

## algorithms in consensus

TODO

- [ ] State
- [ ] Append Entries RPC
- [ ] RequestVote RPC
- [ ] Rules for servers

### basic

machine's role: leader, follower, candidate(in election term)

raft divide the time into different terms, there is at most one leader in each term

term plays a logical clock in Raft they allow servers to detect obsolete information

if candidate or leader discover that its term number is smaller, it will become follower

only two rpc is required:

- RequestVote(election)
- AppendEntries(heartbeat)
  
- Transfer snapshot

### Leader Election

Leader use heartbeat to establish its authority and prevent new election

Election happens when follower receive no heartbeat

Election:

1. followers increment term
2. transitions to candidate
3. RequestVote RPC

first-come-first-served basis, vote for a most one candidate

use term number for comparison

no candidate obtains majority: use randomized election timeouts to ensure this resolved quickly

### Log replication

leader issues AppendEntries RPCs in parallel to others.
leader will retry indefinitely until all followers eventually store all log entries

log is the item with log index, term and opt

ensures:

- If two entries in different logs have the same index and term, then they store the same command.(leader create at most one entry and entries never change their position)
- If two entries in different logs have the same index and term, then the logs are identical in all preceding entries.(AppendEntries RPC contains index and term that immediately precedes the new entries)

use strong leader to ensure consistency: forcing followers' log to duplicate its own

leader maintain a nextIndex for all server, when a new leader elected, it will change the nextIndex according to its own log entries

if AppendEntries RPC fails due to nextIndex, leader will decrement nextIndex until success

### Security

#### Election restriction

Leader's log must be up-to-date

- logs with later term
- logs is longer when they have the same term

#### Committing entries from previous terms

never committed by counting replicas

#### Safety argument

[TODO] review the Figure 9

### Follower and candidate crashes

ensure by retrying indefinitely

### Timing and availability

leader election is the key:

broadcastTime << electionTimeout << MTBF(average time between failures for a single server)

0.5ms - 20ms << 10ms - 500ms << several months

## Cluster membership changes

Issue an joint consensus that combines old and new configuration:

- Log entries are replicated to all servers in both configuration
- Any server from either configuration may serve as leader
- Agreement in election requires separate majorities from both old and new

new server can join the cluster in a duration when they can not vote(to catch up logs)

leader will not be deposed by larger term if it is able to send heartbeats

[TODO] review samples

## Log compaction

Snapshotting is the simplest approach to compaction

other methods:[TODO]

- log cleaning[ROSENBLUM, M., AND OUSTERHOUT, J. K. The design and implementation of a log-structured file system. ACM Trans. Comput. Syst. 10 (February 1992), 26–52.
],
- log-structured merge tree[ REED, B. Personal communications, May 17, 2013.
]

each server takes snapshots independently

- *last included index*[the last entry the state machine has applied]
- *last included term*[the term of this entry]

leader use Install Snapshot RPC to send snapshots to followers that are too far behind[Size ?], follower decide what to do with it

if follow the strong leader's role(only leader can create snapshot):

- waste network bandwidth
- leader's implementation will be more complex(need to send snapshots in parallel with replicating new log entries)

copy-on-write

## Client interaction

consensus-based system's general problem: how clients interact with Raft(how clients find the cluster leader and how raft supports linearizable semantics)

clients send all requests to leader, the server will inform clients who is the leader

to unique serial numbers to solve linearizable semantics(each operation appears to execute instantaneously, exactly once, at some point between its invocation and its response)

- a leader must have the latest information on which entries are committed(it have the all committed entries but don't know which those are)
- leader must check whether it has been deposed before handling the read-only request

## Implementation and evaluation

raft resources: https://raft.github.io/

### Understandability

[Really interesting to assign such a understandability test XD]

### Correctness

proof[TODO if time is available]

### Performance

The most important case: an established leader is replicating new log entries

minimal number of messages(a single round-trip from the leader to half the cluster)

batching and pipelining requests

randomization in election helps with the split votes situation(10s - 287ms)

35ms on average to elect a leader

## Related work

Raft use leader election as an essential part of the consensus protocol

this approach makes raft more simple and easy to understand
(
  it precludes some performance optimizations
)

raft minimizes the functionality in non-leaders.(log entries in raft flow in only one direction)

raft only contains 4 message types

cluster membership changes:[TODO]

- Lamport's original proposal[L A M P O RT, L . The part-time parliament. ACM Transac- tions on Computer Systems 16, 2 (May 1998), 133–169.]
- VR[LISKOV, B., AND COWLING, J. Viewstamped replica- tion revisited. Tech. Rep. MIT-CSAIL-TR-2012-021, MIT, July 2012.]
- SMART[LORCH, J. R., ADYA, A., BOLOSKY, W. J., CHAIKEN, R., DOUCEUR, J. R., AND HOWELL, J. The SMART way to migrate replicated stateful services. In Proc. Eu- roSys’06, ACM SIGOPS/EuroSys European Conference on Computer Systems (2006), ACM, pp. 103–115.]

## Conclusion

decomposing the problem and simplifying the state space will improve the understandability

[Due to the CAP theory, every design and implement in distributed system is the subclass of Tradeoff]
