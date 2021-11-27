ddia page 1-10

# data intensive apps
big application from many small building blocks
- database:Store data so that they, or another 
application, can find it again later
- cache:Remember the result of an 
expensive operation, to speed up reads
- search indexes:Allow users to search data by 
keyword or filter it in various ways
- stream processing:Send a message to another 
process, to be handled asynchronously
- batch processing:Periodically crunch a large 
amount of accumulated data

# reliability
- continuing to work correctly, even when things go
wrong.
- fault-tolerant/resilient
- failure: the system as awhole stops providing the required
service to the user
- hardware fault: constantly happens for large data center
- software fault: hard to detect and figure out
- human fault: operational mistakes

# scalability
- a system’s ability to cope with increased load.
- definition of load: varies common QPS; caxhe hit rate; read write ratio for db
- describe performance: throughput & latency
- use percentile (p90,  p99,  p999)  instead of avg
- SLA & SLO based on pXXX
- scaling up /vertical vs scaling out/horizontal

# maintainability
to be continued. 
