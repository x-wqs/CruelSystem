# 设计数据密集型应用
- https://vonng.gitbooks.io/ddia-cn/content/

# CH1 - Reliability, Scalability, Maintainability
- data intensive instead of computing intensive
- data store, database, cache, search index, stream processing, batch processing
- reliability, hardware/software failure, human mistake, fault tolerance, resilient
- fault is not failure, MTTF (mean time to failure) of hard drive is 10 to 50 years
- scalability, data scale, traffic, complexity
- maintainability, engineer, operation

## Search Engine Workflow

-                        +---> Cache <------------------+
- Client ---> AppSver ---+---> Database --------+---> AppSver
-                        +---> FullTextIndex ---+
-                        +---> MessageQueue ---> AppSvr ---> FullTextIndex
