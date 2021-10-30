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

## Performance
- partition URL DB
- cache DNS
- locality, geographically close to website hosts
- short timeout
- checksum to avoid duplicates
- max URL length for spider trap
- data noise like ads, spam URLs


## Extensibility
- Content Seen -+-> Link Extractor
                +-> PNG Downloader
				+-> Web Monitor

## Fault Tolerance
- ConsitentHash: replace a dead host and distribute load
- save Frontier queue into disk
- shift queue file to another host

## Data replication and partition
- based on hostname

## Enhancement
- Telemetry
- Analytics

## Questions:
- How to handle URLs constructed by JavaScript

## Hints
-    KB   MB   GB   TB   PB
-    K    M    B

## Terminology
- for the sake of, 为了什么目的、好处
- courteous
- appropriate
- canonical hostname
- off-limits
- canonical form
- probabilistic
