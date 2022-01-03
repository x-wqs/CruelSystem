# DDIA: pp 181-190

### Limitations of Quorum Consistency

With a smaller w and r (w + r <= n)

- read staled data
- lower latency and higher availability



With quorum satisifed w and r, may also read staled data

- sloppy quorum
- concurrent writes
- write happend concurrently with a read
- write failed <= not roll back on the successful replicas
- the node carrying new value fails => recover from a replica with old value => disobey quorum



Monitoring staleness

- leader-based: substract a follower's current position from the leader's current position
- leader-less: hard. Expect percentage of stale reads depending on the parameters n, w, and r.



Sloppy Quorums and Hinted Handoff

- If one node dies, there might be fewer than w or r reacheable nodes => no longer reach a quorum
- *sloppy quornum*: writes and reads still require w and r successful responses, but those may include nodes that are not among the designated n "home" nodes for a value.
- *hinted handoff*: Once the network interruption is fixed, any writes that one node temporarily accepted on behalf of another node are sent to the appropriate "home" nodes.
- Suitable for increasing write availability => may read staled data, because the latest value may have been temporarily written to some nodes outside of n.



Multi-datacenter operation (Cassandra and Voldemort)

- The number of replicas *n* includes nodes in **all** datacenters
- Each write is sent to all replicas, but only waits for ack from a quorum of nodes within its local datacenter.



Detecting Concurrent Writes

- simialr to multi-leader replication



Last write win (discarding concurrent writes)

- declare that each replica need only store the most “recent” value and allow “older” values to be overwritten and dis‐ carded
- Force an arbitrary order on the writes, e.g., timestamp. However, timestamp is not sufficiently reliable.
- Only safe case: ensure that a key is only writ‐ ten once and thereafter treated as immutable.



The "happens-before" relationship and concurrency

- How to decide whether two operations are concurrent or not?
  - Not concurrent: B's operation is causally dependent on A: e.g., B update inserted value from operation A.
  - Concurrent: When each client starts the operation, it does not know that another client is also performing an operation on the same key.



Concurrency, Time, and Relativity

- We simply call two operations concurrent if they are both unaware of each other, regardless of the physical time at which they occurred.



Capturing the happens-before relationship

- Server maintains a version for every key, increments the version number every time that key is written, and stores the new version number along with the value written.
- Client must read a key before writing. When the client reads the key, server returns all values that have not been overwritten and the latest version number.
- Client write the key: include the key number, merge together all values that it received in the prior read. => the response from a write request: all current values
- Server receive the write with the version number: overwrite all values with that version numbre or below. Keep all values with a higher version number.

=> with version number, we know which previous state the write is based on.

=> If there is a write without version number, it is concurrent with all other writes, so it will not overwrite anything. => return as one of the values on subsequent reads.



Merging concurrently written values

- if several operations happen concurrently, clients have to clean up afterward by merging the concurrently written values.

- Remove operation might fail after union different versions of data.

