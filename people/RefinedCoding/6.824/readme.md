# MIT 6.824

## Build
- ./.check-build

## Lab 1 - Sequential
- cd src/main
- go build -race -buildmode=plugin ../mrapps/wc.go
- rm mr-out*
- go run -race mrsequential.go wc.so pg*.txt
- more mr-out-0

## Reading
- https://github.com/LebronAl/MIT6.824-2021/blob/master/docs/lab2.md
- applier / ticker
- https://github.com/crimson-gao/MIT-6.824-spring2021/blob/master/src/raft/raft.go

## Drop
- https://github.com/zhouchuyi/MIT-6.824-2020/blob/master/src/raft/raft.go
- https://github.com/siclait/6.824-cljlabs-2020

- election/leader ticker
- https://github.com/JellyZhang/mit-6.824-2021/blob/master/src/raft/raft.go

- leader loop
- https://github.com/viktorxhzj/6.824/blob/master/src/raft/raft.go

- start election
- https://github.com/h1deOnBush/6.824-2021/blob/master/src/raft/raft.go

- timeout watcher / log applier
- https://github.com/MrGuin/6.824-labs-2021/blob/master/src/raft/raft.go

- ticker / NotifyApplyCh
- https://github.com/vnguyen12/6.824-Distributed-Systems/blob/main/src/raft/raft.go

- Image Update
- https://github.com/lzlaa/6.824/blob/master/raft/raft.go


