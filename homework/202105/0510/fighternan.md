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
several k commodity PCs (4GB mem, 1G NI, DFS based on hard drive), scheduling system

Say there are M Map tasks and R Reduce tasks:

1. input file => splitted into M pieces (each 64MB); duplicate these M pieces across the cluster
2. has a master and many workers;
3. master => map worker gets a map task; it reads the split (list of k-v pairs) and generates a new (intermediate) list of k-v pairs (that are in memory); 
4.  intermediate k-v pairs are persisted periodically; they are partitioned by the partition function (sth like hash(key) Mod R); worker => master the location 
5. master => reduce worker gets a reduce task; use RPC to read a partition of intermediate k-v pairs. Sort the pairs by key (use external sort if not fit in memory), so same keys are grouped together for reducing. Append output of user-specified reduce into a file for this reduce partition.

To be continued....


