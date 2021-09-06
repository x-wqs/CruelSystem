## MapReduce p1-p10



overall

- functional programming model: map→k/v→reduce(merge)
- user-friendly abstraction: hide messy parallelization details

contribution: 

- simple and powerful interface that enables large-scale parallelized computation

programming model:

- map(k1, v2)→list(k2,v2)
- reduce(k2, list(v2))→list(v2)

important usage examples:

- inverted index
- distributed sort

implementation (algo steps):

1. split
2. master assign tasks to workers
3. mappers map and write to local disk
4. reducers remote read, sort keys(why?) 
5. reduce

fault tolerance

- worker failure: reschedule another one to work, and re-execute if a mapper fails
- master failure: checkpoints or simply abort
- deterministic(?)

Optimization & Refinement:

- Backup tasks to resolve the stragglers
- well-balanced partitioning, eg, hash(key) mod R
- combiner: partial merging locally to speed up
- reader: support different I/O types
- skipping bad records: instead of failing all
- local execution on single machine: for debugging purposes
- HTTP server and status pages: for monitoring
- counter facility: sanity checking

performance...

experience/usages:

- large-scale ML, extraction of texts, graph computations at Google