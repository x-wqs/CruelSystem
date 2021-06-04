## Lab1: MapReduce

Install go

```
brew install go
```

Compile the `wc.go` as a plugin with `-race` detector

```
go build -race -buildmode=plugin ../mrapps/wc.go
```

Run a sequential mapreducer with word count execuable 

```
go run -race mrsequential.go wc.so pg*.txt
```

### Description

A MapReducer of a word count app.

1. One coordinator: the implementation should be in `mr/coordinator.go`
2. Multiple workers: the implementation should be in `mr/worker.go`
3. workers talk to coordinator via RPC - the implementation should be in `mr/rpc.go`
4. Each pg-xxx.txt is the input to one mapper. #pg-xxx files == #mappers
5. Each mr-out-xxx file is the output of a reducer

6. Testing: 

   ```
   bash test-mr.sh
   ```

### Other Stuff

- the results together from all mr-out-xxx should be the same to the result from mrsequential
- `nReducer` is 10 by default. This number is the number of reducers/buckets to all the keys from results of mapper
- Change codes only in `mr/`, and rebuild
- The results from mappers can be saved in current dir in json files (using encoding/json package)
- all the intermidate/final files are saved in the cur files
- Read `mrsequential.go` first
- lock shared data in coordinator (?)
- Coordinator detects a dead worker by waiting 10s, and then send the task to another worker

## Ref

- the lab https://pdos.csail.mit.edu/6.824/labs/lab-mr.html
- go command https://golang.org/cmd/go/
- go language https://gobyexample.com/
- 

