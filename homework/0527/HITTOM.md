## source 
> 6.824 Lab 1 MapReduce: https://pdos.csail.mit.edu/6.824/labs/lab-mr.html

> lab guidance: https://pdos.csail.mit.edu/6.824/labs/guidance.html

> set up go tutorial: https://pdos.csail.mit.edu/6.824/labs/go.html

> debugging by pretty printing: https://blog.josejg.com/debugging-pretty/

## lab guidance:
### hardness
> easy: a few hours

> moderate: 6 hours per week

> hard: more than 6 hours per week

> a few hundred lines but conceptually complicated and detailed

### definition
> fault: logic bug

> error: caused by fault

> forward approach: fault --> error, guess-and-check

> backward approach: error --> fault, more applicable to distributed progress

## set up go
## go version
> recommand version: 1.15

> my version: https://golang.org/dl/  go1.15.12.linux-amd64.tar.gz

## lab1: MapReduce
### introduction
> implement a worker  process that calls application Map and Reduce functions, and a coordinator process

### get started
```
go build -race -buildmode=plugin ../mrapps/wc.go
go run -race mrsequential.go wc.so pg*.txt
more mr-out-0
```

### my job(moderate/hard)
> implement a worker(bunch of machines in real system, but many processes on one machine this lab) and a coordinator(one process).

> workers talk to coordinator via RPC

> routines are in main/mrcoordinator.go and main/mrworker.go, don't change these files. put implementation in mr/coordinator.go, mr/worker.go, mr/rpc.go

```
go build -race -buildmode=plugin ../mrapps/wc.go
go run -race mrcoordinator.go pg-*.txt # expectation: Done() argument error
go run -race mrworker.go wc.so # in other windows
cat mr-out-* | sort | more
bash test-mr.sh # ??
```

### a few rules
> The map phase should divide the intermediate keys into buckets for nReduce reduce tasks, where nReduce is the argument that main/mrcoordinator.go passes to MakeCoordinator().

> main/mrcoordinator.go expects mr/coordinator.go to implement a Done() method that returns true when the MapReduce job is completely finished;

### hint
> to avoid partially written file, use a temporary file(ioutil.TempFile) and rename(os.Rename) it when task done

### no-credit challenge
> implement distributed grep

> run coordinator and workers on different machine

## terms
> moderate 中等

> deviate 偏离

> enshrine 铭记

> manifest 显现

> deliberate 刻意

> hypothesis 假设

> bisection 二等分

> percieve 感知

> condense 凝结

> facilitate 促进

> mileage 里程

> tweak 调整

> resort 采取

> obscure 朦胧

> glue 胶水

> sporadical 零星的

> grain 谷物

> interleaving 交织

> extraneous 外来的
