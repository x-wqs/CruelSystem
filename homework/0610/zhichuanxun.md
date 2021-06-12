# MR

## implementation

tasktype {
  map, reduce, done
}

asktask
finishtask

1. implement RPC
2. coordinator handlers for RPC
3. work loop to send out RPC
4. helper to handle intermediate files
5. worker map
6. worker reduce
7. assign task

[TODO] go condition variables

## Alternative Sync

- wait in worker instead of coordinator
- use time.Sleep() in coordinator instead of CondVar
- Channels

coordinator assign task to worker

## Common Design mistakes

- pushing too much work to the coordinator
- sending redundant RPCs

## Debugging

- Printfs for debugging
  - Conditional Printf(DPrintf in raft/utils.go)
  - color scheme for RPCs, columns, server IDs
  - redirect output files
- ctrl-\ kill a program and dump all its goroutine stacks

## MR

- MR seems fairly easy to use for counting and sorting, other complex
  - data mining
- shuffle step
  - combing occurs at map, result written to intermediate file
  - sort occurs and reduce after all outputs from map are read by reduce

google cloud dataflow
spark

- mapper store files locally and not on GFS
  network bandwidth
- leaders necessary to have distributed systemd
  no bitcoin

- timeout to retry failed tasks
  - paper describe timeouts are used to restart tasks when worker fail
  - backup tasks are used to speed lagging tasks at the end of execution
  - timeouts both detect worker failure and detect slow tasks


## Code design

organize massive pieces of code:

- separate by RPC sender+ handler
- put all definitions of state(structs) together
- factor out common pieces of code into functions

[TODO]go interface cast