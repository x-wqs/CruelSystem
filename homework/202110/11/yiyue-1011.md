#### detailed design

- rules stored on disk. workers pull rules frequently and store in cache
- client -> request -> server, request -> rate limiter first
- rate limiter loads rules from cache

#### distributed environment

##### race condition
- two requests concurrently read the counter value before write
- each increment the counter by one without checking the other thread
- counter value incremented by 1 only

##### synchronization issue
- one rate limiter not enough
- multiple rate limiter, requires synchronization

*Solution*

1. sticky sessions, a client send traffic to same rate limiter
- not scalable
- not flexible

2. centralized data stores

#### performance optimization

- multi-data center setup - latency is high for users far away
- synchronize data with an eventual consistency model

#### monitoring

Gather analytics to check whether effective
- algorithm
- rules

- if rules too strict and valid requests droppped, relax the rules
- sudden increase in traffic, replace the algorithm


# consistent hashing

- To achieve horizontal scaling, distribute requests/data efficiently and evenly across servers

## rehashing problem

n cache servers, to balance the load
`serverIndex = hash(key) % N`

if new servers are added, existing servers are removed, unstable

## consistent hashing

- when hash table is re-sized & consistent hashing is used
- avg k/n keys need to be remapped 
- k = number of keys
- n = number of slots

## hash space and hash ring

add a server will only require redistribution of a fraction of keys