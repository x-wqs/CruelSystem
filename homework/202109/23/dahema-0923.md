# design consistent hashing

## The rehashing problem
- serverIndex = hash(key) % N
- works well when the size of the server pool is fixed, and the data distribution is even
- if add or delete servers, need a lot data redistribute

## consistent hashing
- hash space and hash ring
- hash servers：map servers to the ring using a uniformly distributed hash function.
- hash keys:
- server lookup, clockwize find the cloest hash server by hash key
- add a server, require redistribution of a fraction of keys.
- remove a server require redistribution of a fraction of keys
- issues:
  - after add/remove server, hash server space is not evenly distributed
  - hot spot
- solution
  - virtual node: each server is represented by multiple(large) virtual nodes on the ring
  - as number of vritual nodes increase, the distribution of keys becomes more balanced, avoid hot spot.
- find affected keys
  - add server: moves anticlockwise around the ring until a server is found, keys located in this range redistributed to new virtual node, repeat this for all the virtual nodes.
  - remove server: moves anticlockwise around the ring until a server is found, keys located in this range redistributed to next virtual node, repeat this for all virtual nodes.


## wrap up
- consittent hashing
  - minimized the number of keys need to redistributed 
  - scal horizontally
  - avoid hot spot by divide the range to small virtual nodes
- usages
  - Dynamo db
  - Cassandra
  - load balancer
