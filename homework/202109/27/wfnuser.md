# Building Software Systems at Google 

## Google Web Search
index servers: 
- query -> (docid, score)
- sharded by docid

doc servers:
- (docid, query) -> (title, snippet)

cache servers:
- cache index and doc snippets
- 延迟优化明显

- simple batch indexing system

- Google data center

- Replicas... 最终副本的数量是的所有的index machine可以在内存里保存完整的索引

- In-Memory Indexing System
  - 每次访问的节点数上千
  - 可用性问题

- canary requests

- query serving system