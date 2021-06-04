## Contribution
A new abstraction for parallel computation on commodity hardware with the following features: 
1. parallelization (map + reduce)
2. fault tolerance (re-execution)
3. data distribution 
4.load balancing  


## Programming model
The programming model is simply map + reduce, inspired by functional programming. 
Examples: word/url count; distributed grep; reverse web-link graph (url, list(pags)); inverted index (word, list(doc Id)); distributed sort (will touch later);

## implementation
several k commodity PCs (4GB mem, 1G NI, GFS based on hard drive), scheduling system

Say there are M Map tasks and R Reduce tasks:

1. input file => splitted into M pieces (each 64MB); duplicate these M pieces across the cluster
2. has a master and many workers;
3. master => map worker gets a map task; it reads the split (list of k-v pairs) and generates a new (intermediate) list of k-v pairs (that are in memory); 
4.  intermediate k-v pairs are persisted periodically; they are partitioned by the partition function (sth like hash(key) Mod R); worker => master the location 
5. master => reduce worker gets a reducing task; use RPC to read a partition of intermediate k-v pairs from map workers. Sort the pairs by key (use external sort if not fit in memory), so same keys are grouped together for reducing. Append output of user-specified reduce into a file for this reduce partition.

### Master data structure
For each task (map or reduce), master stores the state (idle, in progress, completed). 
Master keeps the location of intermediate files. It sends location to workers for reducing tasks. 

### Fault tolerance
Failure 
failed worker: a worker does not respond to ping
1. reschedule all map tasks (finished or in progress) and in-progress reduce tasks (this is because output of map tasks is stored locally on workers, but reduce tasks are stored in global FS. 
failed master: the state of master is periodically persisted so it is recoverable

### Locality
The Input data is managed by GFS (that is, a block has several copies across the cluster). Master schedule tasks according to the location, so that no or very few network traffic is generated.

### Task granularity
M steps of map phase and R steps of reduce phase.
Master makes O(M + R) scheduling decisions and keeps O(M*R) states.
In practise, M is decided by input data size (tend to make size of each map input <= 64 MB); R is a small multiple of #workers.

### Backup Tasks
straggler: some machines are too slow (bad disk or CPU)
Schedule backup execution for remaining tasks when a MapReduce task is to finish and use results from the faster ones. In practise, this saves a lot of time (44% for distributed sort). 
