# SCALE FROM ZERO TO MILLIONS OF USERS
## data centers
### challenges
- Traffic redirection
- Data synchronization
- Test and deployment

## message queue
- message queue can decoupe the architecture for building a scalable and reliable application
- 常见 Kafka 考点
  - 重试：https://www.infoq.cn/article/51xshw2opsmakhhmtth8
  - At least:  processing then ack
  - At most: ack then processing
  - Exactly once: distributed transaction on processing then ack
- 常见 RabbitMQ 考点
  - 重试：当一个request 超过 time out的时间， 
  - 去重：幂等

## Logging, metrics, automation
- Logging:identify errors and problemsin the system
- Metrics: gain business insights and understand the health status of the system
  - Host level metrics: CPU, Memory, disk I/O, etc
  - Aggregated level metrics: entire database tier, cache tier
  - Key business metrics: daily active users, retention, revenue, etc.
- Automation: Continuous integration Continuous deployment, CICD.

## Database scaling
### Vertical scaling
- dding more power (CPU, RAM, DISK, etc.) to an existing machine
- hardware limits
- single point of failure
- overall cost of vertical scaling is high

### Horizontal scaling
- adding more servers
- Sharding key (known as a partition key) determine how data is distributed
- Resharding data: consistent hash
- Celebrity problem(hotspot key problem)
- Join and de-normalization

## summarize 
- Keep web tier stateless
- Build redundancy at every tier
- Cache data as much as you can
- Support multiple data centers
- Host static assets in CDN
- Scale your data tier by sharding
- Split tiers into individual services
- Monitor your system and use automation tools
