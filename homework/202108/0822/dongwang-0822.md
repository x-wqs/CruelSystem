Fault-tolerant KV servive

Use Raft to design a key/value store service which can be tolerant - the service can still work as long as a majority of
service are alive and can communicate. 

Ops:
Put(key, value) update the key with value if key is present else crate a (key, value) pair
Append(key, args) append args to the value of key
Get(key) get the value for key
All these ops should done through a Clerk interim

Requirements:
Strong consistency: for concurrent operations, the result of last completed op must can be observed before next operation.
This is called ***Linearizability*** .  With strong consistency, all clients see the same state-the latest state.  It's harder
for replicated servers to achieve strong consistency. 

part 1:  without worrying the Raft log growing with bound
Clerks send ops to the server whose Raft is a leader and then this kvserver send the op code to the Raft, which will log 
these ops sequentially. If Clerk cannot find which kvserver is leader, it will re-try by sending to a different kvserver. 

part 2:  consider the Raft log have limit size and will discard old logs, which means you have to implement snapshots.
A rebooting server replay the log to restore the state. By using previous snapshot, every time the log grows to a maxsize
it will be saved into a snapshot for later replaying .  
