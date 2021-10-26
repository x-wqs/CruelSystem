# 0000-system-web-crawler

## Purpose
- search engine indexing
- web archiving
- web mining
- web monitoring

## Functional Requirements
- 1B pages / month
- HTML only, store for 5 years

## Non-functional Requirements
- Scalabilty
- Robustness, handle exception
- Politeness, follow rebots rules
- Exensiblity, support new content/format

## Capacity Estimation
- 1B pages / month
- 10^9 / 30 / 24 / 3600 = 10^4 / 30 = 300 QPS
- Peak QPS = 600
- Storage: 1B * 500K * 12* 5 = 30PB
- Bandwidth: 600 * 500K / s = 300MBps = 2Gbps

## High Level
								DNS Resolver							ContentDB  
									^										^  
                                    |										|  
- SeedURL ---> URLFrontier ---> HTML Downloader ---> Content Parser ---> Content Seen?  
					^														|  
					|												Link Extractor  
					|														|  
					|													URL Filter  
					------------------------------------------------------- |  
																		URL Seen?  
																			|  
																		URL DB  
- Seed URL: starting point
- URL Frontier: URL Queue, Kafka? 参考文献？
- URL Filter: filter content types

## Design Design
- DFS, Stack应该不支持
- BFS, FIFO
- Politeness Constraint: Delay for each Queue
- Mapping Table ---> Queue Router ---> Q1 ---> Queue Selector ---> Worker Thread 1  
                                   +-> Q2 -+                   +-> Worker Thread 2   
								   +-> Q3 -+                   +-> Worker Thread 3  
- Mapping Table <Host, Queue>


## TODO
- Priority

## Hints
-	KB	MB	GB	TB	PB
-	K	M	B
