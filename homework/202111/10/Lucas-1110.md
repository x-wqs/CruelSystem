## Cache
- urls.checkSum, 1B * 8 bytes = 8GB
- TODO: bloom filter
- 一个URL用N个哈希计算出不同的BitIndex
- 如果这个URL存在的话，那么这N个BitIndex都应该为1
- 但是一个URL不存在，所对应的N个BitIndex也可能都为1
- False Negative 假阴性不可能
- URL哈希的BitIndex没标记，那就肯定不在，不存在假阴性
- URL没有访问，有可能被标记成访问过，从来不去爬取，可以扩大位向量大小来解决
- URL访问过了，那就肯定能检测到，不可能出现假阴性

## Detail Design
- DFS, reuse the connection for the same site, avoid handshake, Stack应该不支持
- BFS, FIFO
- Path-ascending crawling
- Politeness Constraint: Delay for each Queue
- Mapping Table ---> Queue Router -+-> Q1 ---> Queue Selector -+-> Worker Thread 1  
-                                  +-> Q2 -+                   +-> Worker Thread 2   
-                                  +-> Q3 -+                   +-> Worker Thread 3  
- Mapping Table <Host, Queue>
- Queue Selector: worker 1 : n queue
- Priority: PageRank
- InputURL -+-> Prioritizer ---> F1 -+-> Front Queue Selector ---> Back Queue Router -+-> Q1 -+-> Back Queue Selector -+-> Worker Thread 1
-                            +-> F2 -+                                    |           +-> Q2 -+                        +-> Worker Thread 2
-                            +-> F3 -+                             Mapping Table      +-> Q3 -+                        +-> Worker Thread 3
- Freshness, update history or priority
- Frontier: Queue/Kafka, Hybrid, majority on disk, buffers in memory
- EnqueueBuffer -> Disk -- DequeueBuffer
- robot.txt, first check robots exclusion protocol, put in Map<host, rule> in memory
- Snapshot and restore from checkpoint
