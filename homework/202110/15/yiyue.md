### Consistency
N = the # of replicas
W = A write quorum of size W. For a write operation to be considered successful, write operation must be ackowledged from W replicas
R = A read quorum of size R. 

- If W = 1 or R = 1, operation is returned quickly
- If W or R > 1, the system offeres better consistency, query will be slower
- If W + R  > N, strong consistency is guaranteed

- If R = 1 and W = N, the system is optimized for a fast read
- If w = 1 and R = N, the system is optimized for fast write
- If W + R > N, strong consistency guaranteed
- If w + R <= N, strong consistency not guaranteed

### inconsistency resolution: versioning
Vector clock - [server, version] pair associated with a data item
downsides
 - adds complexity to client because it needs conflict resolution logic
 - [server: version] pairs grow rapily 
  - can set a threshold for the length
  - once exceeds limit, oldest pairs removed

### handling failures

#### failure detection
- requires at least two independent sources of information to mark a server down
- gossip protocol
 - each node maintains a node membership list - member IDs and heartbeat counters
 - periodically increments heartbeat counter
 - periodically sends heartbeats to a set of random nodes, which propagate to another set of nodes
 - nodes receive heartbeats, update membership list
 - if hearbeat has not increased for a predefined periods, member considered offline

#### handling temporary 

sloppy quorum
- choose first W healthy servers for writes and first R healthy servers for reads
- if a server unavailable, another server will process the request temporarily
- when down server is up, changes will be pushed back