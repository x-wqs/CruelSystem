# 设计一个分布式的定时任务调度器

## Functional Requirements 
- schedule job with Cron format
- report failures

## Non Functional Requirements
- scalibility 
- minimize duplicate execution
- job time out
- availbility

## Estimation
- number of job: 100k, 100/s
- 

## Components
- LB
- web server
- S3(job storage)
- schedule service
- job execution message queue
- job worker
- job result message queue
- email notification server

## API
- schedule job
  - request: job_name, job_schedule, job_content, start_at, user_info, time_out
  - repose: job_id, job_info, s3_path
- send failure by email 

## Storage
- metadata table: job_id, cron_format, time_out, job_status, carete_at, job_content_path, start_at, requester_id
- history table: job_id, seq_num, finish_time, log_path, job_status
- requester table: requester_id, ...
- schedule server: priority queue: epoch_run_time, job_id, ..., which sorted by epoch_run_time
- job execution message queue: job_id, time_out, job_content_path, ...


## work flow
- over flow: client -> LB -> web server -> S3/DB
                                 |->  schedule service -> job excution message queue -> job worker -> job  message queue -> failure email notification
- web server: after get request will save job content in S3, and other metadata in DB
- shedule service: 
  - maintain a priority queue, which sorted by job next run time(in epoch) in descending order
  - when push job next run time to the priority queue, set timer to the top of the priority queue
  - if the timer is up, then pop the top out and send to job execution message queue, 
  - repeat check the top and send job to job execution message queue
  - set the time to new top of priority queue
- job execution message queue
  - at least once: worker ack the message after job is done.
  - at most once: woker ack the message after got the message. if job fail, the job is never done. 
  - exactly once: idempotence, refer to https://www.confluent.io/blog/exactly-once-semantics-are-possible-heres-how-apache-kafka-does-it/
- job worker, 
  - fetch content and other infos from s3 and metadb
  - update metadata table status to in_progress, update machine id, start time and etc
  - job success: update metadata table status to success, add history table, like log_path, status
  - job fail: similar as success, but check if need trigger retry loggic. 
    some errors are retryable, like netwoker issue, machine issue
  - job timeout : similar as fail.
- zookeeper: monitor weather job workers' health, if job worker is down, then mark all the in_progress to fail/timeout. 

## Q&A
TODO

## Scale
- LB, web server can easily scale up
- s3/DB shard by job_id
- schedule service can easily scale up
- message queue can use Kafka


## availability
- s3/DB, schedule service, Kafka all should have replica
- replica is only for availability
