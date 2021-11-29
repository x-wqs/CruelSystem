# DDIA pp11-20: Scalability, Maintainability

## Scalability

- The ability to cope with increased load: more concurrent users, larger volumes of data



### Load

Define *Load*:

- requests per second to a web server
- ratio of reads to writes in a database
- number of simultaneously active users in a chat room
- hit rate on a cache
- ...



Example: twitter 2012

- Post tweet: 4.6k requests/sec on average, over 12k reuqests/sec at peak

- Home timeline: view tweets posted by the people they follow. 300k requests/sec

Challenge of twitter: 

- fan-out
  -  Fomal definition: number of requests to other services that we need to make in order to serve one incoming request
  - Concrete example: broadcast one user's new post to all followers. Insert post directly to followers' homelines will cause write heavy problem for hot users.
    - Solution: Insert new post to user's timeline for those normal users; For hot users, followers fetch new posts and merge with homeline when it is read.



#### Describe performance

Batch processing system (e.g., hadoop): throughput

Online system: distribution of response time

- average
- percentile: p50 (median), p95, p99, p999
  - Amazon: p99.9

Percentile can be used in *service level objectives (SLOs)* and *service level agreements (SLAs)*, which defines the expected performance and avilability of a service

- Example: SLA: median response time of less than 200 ms and a 99th percentile under 1s, and up at least 99.9% of time. Customers can demand a refund if the SLA is not met.

Tail latency amplification:

- Some slow backend request cause the slow response of end-user request



#### Test of performance

Example: Queuing delay

- problem: head-of-line blocking: some slow request block the remaining fast requests
- test attention: keep sending requests independently of the response time to fit better to the real case.

Visualize performance

- Calculate good approximation of percentiles at minimal CPU and memory cost: forward decay, t-digest, HdrHistogram
- Aggreagting response time data => add the histograms.



#### Scaling

- vertical scaling, scaling up
- horizonal scaling, scaling out, shared nothing

*elastic*: one system can automatically add computing resources when they detect a load increase. Useful when the load is highly unpredictable.



🌟 An architecture that scales well for a particular application is built around assumptions of which operations will be common and which will be rare -- the load parameters. => make appropriate assumptions first.



## Maintainability

*legacy system*: full of bugs, outdated, ...

Three design principles:

- Operability: easy for operation team
- Simplicity: easy to understand for new engineers
- Evolvability: easy to make changes



Tips to improve **operability**:

1. Provide visibility into the runtime behavior and internals of the system
2. Provide good support for automation and integration with standard tools
3. Avoid dependency on individual machines
4. Provide good documentation and an easy-to-understand operational model
5. Provide good default behavior but also give admin the freedom to override defaults when needed
6. Self-healing where appropriate but also give admin manual control over the system state when needed
7. Exhibit predictable behavior, minimize suprise.



