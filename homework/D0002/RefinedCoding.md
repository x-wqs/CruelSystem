# 20210511: Map Reduce P1 - P6

## 3. Implemation
### 3.5 Task Granualrity
- All / 64MB -> M splits -> R pieces
- M = 200,000, R = 5,000, nodes = 2,000

### 3.6 BackTask
- backup in-progress tasks when almost done
- 30% faster in section 5.3

## 4. Refinements,细化
### 4.1 Partition Function
- originally hash(key) mod R
- special partition func: hash(HostName(URL)) mod R
- output all URL of the same host to one output file

### 4.2 Ordering Guarantees
- sort intermediate key/value asc
- for efficient random access lookups by key

### 4.3 Combiner Function
- Map -> Conbine (Another Map) -> Reduce
- see example in Appendix All

### 4.4 Input & Output Types
- text input
- database to reducer
- memory data to reducer (cross cluser ?)

### 4.5 Side Effects
- write to temp file before rename
- atomic 2pc not supported

### 4.6 Skipping Bad Records
- detect bad records causing deterministic crash
- skip them in order and move forward

### 4.7 Local Execution
- alternative implementation of MR library
- execute map/reduce jobs sequentially

### 4.8 Status Information
- run HTTP server on master
- job status, progress
- bytes of input, intermediate data, output
- standard error and output

### 4.9 Counter
- another small MR job?

## 5. Performance Test
- 2G CPU / 4G memory * 1,800
- distribute grep/sort
- take 44% longer without backup tasks
- 5% longer if 200 of 1746 machine failuer

# 6. TODO: Large Scale Indexing

## Terminology 术语
constrained,受约束的  
optimization,优化  
lengthens,加长  
straggler,散客  
arise for ,为...而产生  
correctable errors,可纠正的错误  
a factor of one hundred,一百倍  
tune,调  
significant ,重大  
repetition,重复  
Zipf distribution,Zipf分布  
commutative,可交换的  
associative,联想的  
boundaries,界线  
auxiliary,辅助的  
atomic two-phase commits,原子两阶段提交  
idempotent,幂等  
deterministically,确定性地  
feasible,可行的  
last gasp,最后一口气  
facilitate debugging,方便调试  
alternative,选择  
sequentially,依序  
eliminates effects,消除影响  
tolerable fraction,容许分数  

## Notes  
Q: git remote set-url origin git@github.com:refinedcoding/CruelSystem.git
error: chmod on /mnt/d/perforce/CruelSystem-1/.git/config.lock failed: Operation not permitted
fatal: could not set 'remote.origin.url' to 'git@github.com:refinedcoding/CruelSystem.git'
A:  sudo git remote set-url origin git@github.com:refinedcoding/CruelSystem.git
  




