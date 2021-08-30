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
- 





