ddia p161-p167

# problems with replication lag
- recap: why do we need replicated db?
  - fault-tolerance
  - higher r/w QPS scalability
  - lower latency to clients
- leader based replication
  - all writes go to a single leader node
  - reads can go to any replicas or leader
- async replication style
  - effect: eventual consistency
  - reason: replication lag

next we will discuss methods to __alleviate__ the problem caused by replication lag
- reading your own writes
- monotonic reads
- consistent prefix reads
  
# read-your-writes consistency
- effect: 
  - if the user reloads the page, they will always see any updates they submitted themselves.
  - however, other users’ updates may not be visible until some later time.
- how to achieve:
  - 1. always read from the leader, not from followers for user's own profile
  - 2. update time < 1 min, from leader, > 1 min, from followers
  - 3. use write timestamp make sure only updated followers can respond
- other issues:
  - cross-device: use both userID & timestamp to decide which dc & replica to read
  
# monotonic reads consistency
- effect:
  - never read data in this way: older -> newer -> older -> newer ...
  - wish to read data: older -> older -> newer -> newer -> (always)
  - stronger than read-your-writes consistency
- how to achieve:
  - make sure that each user always makes their reads from the same replica
  - different users can read from different replicas
  - the replica can be chosen based on a hash of the user ID, rather than randomly. 
  - if that replica fails, the user’s queries will need to be rerouted to another replica..

# consistent prefix reads
- effect:
  - causal dependency btw different writes
    - "How far into the future can you see" -> "About ten seconds usually"
  - reads shouldn't violate causal dependency
  - if a sequence of writes happens in a certain order, then anyone reading those writes will see them appear in the same order.
- how to achieve:
  - 1. make sure that any writes that are causally related to each other are written to the same partition
  - 2. mark causal dependency to msg, use algorithm to make sure msgs are delivered based on causal relationship
