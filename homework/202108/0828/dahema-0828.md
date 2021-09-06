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
- job message queue
- job worker
- failure message queue
- failure email notification server

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
- job message queue: job_id, time_out, job_content_path, ...


## work flow
TODO
classical question to be cleared: at least one
- worker alive
  - success
  - job fail
  - time out
- worker down
  - success
  - job fail
  - time out

## Q&A
TODO

## Scale
TODO
