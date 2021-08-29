# Design Distributed Job Scheduler
- https://www.youtube.com/watch?v=Zs4O-Oo5aTc
- https://leetcode.com/discuss/general-discussion/1082786
- https://captionstrendupdategb.blogspot.com/2021/04/system-design-job-scheduler.html
- https://medium.com/airbnb-engineering/dynein-building-a-distributed-delayed-job-queueing-system-93ab10f05f99

## 功能性需求
- 其他服务或者微服务可以调用
- 协调单次任务，或者周期性任务
- 提供统一的接口，调用者要遵循，比如可执行脚本，可执行文件，Java/Jar/API
- 保存任务执行的结果，并支持查询

## 非功能性需求
- 高扩展性，水平扩展
- 高可用性，容错
- 高性能，单节点限流
- 持久性，保存任务，不会丢

## 数据规模
- TODO

## 数据库设计
- Job(id, name, executorClass, priority, status, parameters)
- JobExecution(id, jobId, startTime, endTime, result)
- Trigger(id, type [OneTime/Cron], startTime, endTime, cronExpression)
- JobExecutor(id, lastSuccessTime, classPath)

## 接口设计
- int createJob(name, classPath, startTime, endTime, cronExpression)
- int listJob()
- int queryJob(jobId)
- int stopJob(jobId)

