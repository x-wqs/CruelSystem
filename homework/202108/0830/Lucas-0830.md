# Design Distributed Job Scheduler
- https://www.youtube.com/watch?v=Zs4O-Oo5aTc
- https://leetcode.com/discuss/general-discussion/1082786
- https://captionstrendupdategb.blogspot.com/2021/04/system-design-job-scheduler.html
- https://medium.com/airbnb-engineering/dynein-building-a-distributed-delayed-job-queueing-system-93ab10f05f99

## High Level Design
- MicroService ----> Kafka ---> Job Scheduler Service ----> DB Cluster ----> Job Executor  -----> Kafka  ----> Job Result Service

## Caller
- MicroService schedules a one-time/recuring job, and put request in Kafka

## Job Scheduler
- consume request message
- TODO: Snowflake ID Generator
- Distribute jobs to database partition

## RDBMS
- RDBMS with ACID, sharding to distribute load, Active/Passive or Leader/Follower for each parition
- TODO: DB sync workflow
- Follower may be promoted when Leader fails
- TODO: At least one follower hold up-to-date data

## Job Executor Service
- TODO: read database partition info from ZooKeeper
- Job executor will choose a database partition with fewest job executors
- update database partition back in ZK
- send heartbeats to ZooKeeper
- pull jobs from the database partition, and update running status and start time for the job row
- if the node fails, other node will detect and get jobs of failed node
- after job executed, it will send job result successful/failed to Kafka
- TODO: broadcast failed node status to all of nodes

## Job Result Service
- store job result in no-sql database
- high write thoughput, TODO-leaderless
- replicates asynchronously, eventual consistency

## Coordinator / ZooKeeper
- store info about above nodes

## Message Queue / Kafka
- scale producer / consumer independently
- decouple producer and consumer
- lower latency for the producer
- durability and reliability for consumer
- control traffic throttle / limit
- message ordering

## TODO
- unreliable clock, multiple NTP servers?
- exactly once delivery





