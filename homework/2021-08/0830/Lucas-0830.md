# Design Distributed Job Scheduler
- https://www.youtube.com/watch?v=Zs4O-Oo5aTc
- https://leetcode.com/discuss/general-discussion/1082786
- https://captionstrendupdategb.blogspot.com/2021/04/system-design-job-scheduler.html
- https://medium.com/airbnb-engineering/dynein-building-a-distributed-delayed-job-queueing-system-93ab10f05f99

## High Level Design
- MicroService ----> Kafka ---> Job Scheduler Service ----> DB Cluster ----> Job Executor  -----> Kafka  ----> Job Result Service
- MicroService schedules a one-time/recuring job, and put request in Kafka
-                                                         
