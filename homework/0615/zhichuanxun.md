# Raft

log divergence
log catchup
persistence
snapshots
linearizability

## Divergence

leader election rule:

- majority
- at least up-to-date

## Log catchup

erasing log entries

catch up quickly

## Persistence

strategy1: re-joins, replay log
strategy2: start from persistence state

vote for:
log[]: promise the log to commit
current term: monotonic increasing

Snapshotting is the simplest approach to compaction

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

## Service recovery

- replay log: recreate state
- snapshot = state contains all ops through => cut the log through

## Using Raft

service -------apply channel--------- raft

## Correctness: linearizability

== like a single machine

1. total order of operations
2. match real time
3. real return result of the last write