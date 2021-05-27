# Lab 1
- https://pdos.csail.mit.edu/6.824/labs/lab-mr.html

## Target
- implement rpc, coordinator/master & worker

## Learn
- without MapReduce, src/main/mrsequential.go
- word counter: mrapps/wc.go
- text indexer: mrapps/indexer.go
- build with -race, run with -race

## Functions
- worker ask the coordinator for a task
- coordinator create a handler to return the task
- worker talk with the coordinator with RPC/UnixSocket ?
- implement backup task with a worker cannot finish the task in 10 seconds
- how to test backup task
- the coordinator will set task as new when not finished in 10 seconds
- test with main/test-mr.sh, the script will hang as the coordinator never stops
- reducer output to the file mr-out-X
- wc test
- indexer test
- map paralelism test
- reduce paralelism test
- crash test
- which file is output by mapper?
- intermediate key/value pairs in JSON files
- reducer will use ihash(key) to pick reduce tasks
- pseudo task to let worker quit
- reduces cannot start until the last map job has finished
- the coordinator use a single thread to handler task request from worker
- the coordinator will not be blocked when handling the task request from one worker
- use mrapps/crash.go application plugin to test crash recovery
- 在mapper/reducer生成结果的时候使用临时文件，然后重命名
- test-mr-many.sh will test test-mr.sh with timeout

## Teminology
cope with
consisting of,
as et unstarted
straightforward
extraneous tasks
bare-bones script
 



