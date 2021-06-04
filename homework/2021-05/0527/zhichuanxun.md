# Map Reduce[补]

## Abstract

MapReduce is a programming model

## Introduction

- parallelize the computation
- distribute the data
- handle failures

- hide the messy details of parallelization
- fault-tolerance
- data distribution
- load balancing

## Programming Model

key/value == map ==> intermediate key/value == reduce ==> merge same intermediate key

## Implementation

### Execution

1. input data => M splits, starts up many copies of the program on a cluster of machines
2. one of the copy is master, rest are workers, M map tasks and R reduce tasks to assign.
3. map worker read the contents of the corresponding split, parses key/value pair out of the input data and passes each pair to the use-defined Map. intermediate kv paris produced by Map are buffered in memory
4. buffered pairs are written to local disk, partitioned into R regions by partitioning function.locations are passed back to master
5. reduce worker use rpc to read the buffered data, use external sort to sort data by intermediate key
6. worker pass the key and set of intermediate value to Reduce function, append result to a final output file
7. return back to user code

### Master Data Structures

state: idle, in-process, completed

store the locations and size of the R

### FT

#### Worker Failure

completed worker are marked by idle and eligible for scheduling on other works.

completed map tasks are re-executed because output is stored on fail machine and inaccessible
completed reduce tasks do not need to be re-executed because output is stored in a global file system

reduce task will be notified of re-execution

#### Master Failure

checkpoint

single master

retry MR operation if they desire

#### Semantics in the Presence of Failures

when task completes, worker send a message to the master, master ignore if already received

reduce task will rename the output file, and it's atomic is guaranteed by underlying file system.

majority operators are deterministic

### Locality

most input data is read locally

### Task Granularity

O(M+R) scheduling, O(M*R) state

### Backup Tasks

Master schedule backup executions of remaining in-process tasks

## Refinements

### Partitioning Function

default is hash(key) mod R
hash(Hostname(urlKey)) mod R

### Ordering Guarantees

intermediate kv are processed in increasing key order

### Combiner Function

user can specify combiner to do partial merging before sending it 

### Input and Output Types

reader
different types: file, db ...

### Side effects

multiple output files with cross-file consistency requirement should be deterministic

### Skipping bad records

can detect which records cause deterministic crashes and skip these records

### Local Execution

use local execution to debug and test

### Status Information

master provides dashboard

### Counters

Counter is useful for validating

## Performance

Performance is good that time but hard to improve(for now)

## Experience

It seems that Google doesn't use MapReduce for now

## Related work

active disks -> locality optimization

eager scheduling + fault skip