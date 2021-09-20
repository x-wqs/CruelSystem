## Data centers
- goDNS - a DNS service allows domain names to be resolved to IP addresses based on location
- significant data center outage - direct all traffice to a healthy data center
- multi-data center setup challenges
  - Traffic redirection
  - Data synchronization - replicate data across multiple data centers
  - test and deployment - test at different locations. Automated deployment tools are needed
- need to decouple different components of the system to scale them independently

## message queue
- a durable component, stored in memory, supports asynchronous communication
- buffer - distributes asynchronous requests
- producer -> publish -> message queue -> consume -> consumer -> subscribe -> message queue
- producer can post a message -> queue when consumer unavaiable to process

## logging, metrics, automation
- logging: monitoring error logs
- metrics
  - host level metrics - CPU, Memory, disk I/O
  - aggregated level metrics - performance of database tier, cache tier
  - key business metrics - daily active users, retention
- automation: continuous integration

## data scaling

## verticle scaling
- Add more power
- draw backs
  - hardware limits
  - greater risk of single point of failures
  - overall cost is high

## horizontal scaling (sharding)
- add more servers
- servers share same schema with different data
- sharding key - one or more columns determine how data is distributed
- resharding data
  - a single shard couldn't hold more data
  - certain shards experience shard exhaustion faster than others
- celebrity problem - excessive access to a specific shard causes server overload
- join and de-normalization