system design basic -- grokking the system design interview

# key characteristics of distributed systems
## scalability
- ability to handle increased demand
- horizontal scaling: cassandra, mangodb via add machines

## reliability
- probability of a system will fail in a given period
- redundancy, distributed, replications

## availability
- like 99, 999, up time
- HA: most of time online, but poor reliability: easy to attack

## manageablity
- monitoring facilities: promethus, opentracing, fault detection

# Load balancing
- spread traffic across a clusters of servers
- increase responsiveness & availability
- between client & web servers & application server & db server
- LB algorithm
  - least connection
  - least response time
  - round robin
  - IP hashing
- redundant lb: monitor the health of each other and take over whe fails

# cache
## cache invalidation
- write through
- write around
- write back
## cache consistency

# sql vs nosql
- acid: choose SQL
- schema: static or dynamic
