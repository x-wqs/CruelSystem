# Lab1 Mapreduce
## Intro
### Worker process
It calls the map and reduce function, which could deal with the files(read/write).
### Coordinator process
Take care of the workers if failure happens
## Environment settings
### Tools
GIt, GO
The initial software is in Github. 
### mrsequential
It implements a simple sequential mapreduce.
Url -> src/main/mrsequential.go.
### mrapps
Word count app
### indexer
Text indexer
## The assignment needs to implement
DIstributed MapReduce, which includes the coordinator and the worker
### Coordinator
One coordinator process is needed.
### Worker
Could be one or more worker processes.
### Some reference code
main/mrcoordinator.go <br />
main/mrworker.go
### Code needs to implement
mr/coordinator.go <br />
mr/worker.go <br />
mr/rpc.go	
## How to run the code
### Build the word count plugin
`$ go build -race -buildmode=plugin ../mrapps/wc.go`
### Run the coordinator in main directory
`$ rm mr-out*`<br />
`$ go run -race mrcoordinator.go pg-*.txt`
### Run workers in other window(terminal)
`$ go run -race mrworker.go wc.so`
### Run below commands after worker and coordinator finish their work
`$ cat mr-out-* | sort | more`
## Supported test script
main/test-mr.sh
## Some rules need to follow
Map should work to divide the intermediate keys into n buckets. <br />
The worker should produce output of the n’th reduce task mr-out-X <br />
Format of mr-out-X: Contains one line per Reduce function output with a "%v %v" format <br />
Code should work with the original version <br />
Worker should produce output files in the current directory for later reads <br />
Implement a Done() function in the mr/coordinator.go <br />
Work process should exit after the job is done. <br />
