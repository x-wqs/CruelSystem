# lecture 1 notes

why distributed systems?

- sharing: connect physical separated machines
- increase capacity by parallelism
- tolerate faults
- achieve security via isolation

paper oriented lecture

## main topics

abstraction; infra for apps:
- storage
- communication
- computation

## fault tolerance

- availablity
- recoverablity
-> relicated servers

## consistency

well-defined behavior 尤其对于replica很难

## performance

scalable output

  Scaling gets harder as N grows:
- Load im-balance, stragglers, slowest-of-N latency.
- Non-parallelizable code: initialization, interaction.
- Bottlenecks from shared resources, e.g. network.

## implementation

RPC
threads
concurrency control

重点在fault-tolerance/storage

## Mapreduce

see previous

MapReduce limits what apps can do:
- No interaction or state (other than via intermediate output).
- No iteration, no multi-stage pipelines.
- No real-time or streaming processing.

how to get good load balance?
- many more tasks than workers

Other failures/problems:
  * What if the coordinator gives two workers the same Map() task?
    perhaps the coordinator incorrectly thinks one worker died.
    it will tell Reduce workers about only one of them.
  * What if the coordinator gives two workers the same Reduce() task?
    they will both try to write the same output file on GFS!
    atomic GFS rename prevents mixing; one complete file will be visible.
  * What if a single worker is very slow -- a "straggler"?
    perhaps due to flakey hardware.
    coordinator starts a second copy of last few tasks.
  * What if a worker computes incorrect output, due to broken h/w or s/w?
    too bad! MR assumes "fail-stop" CPUs and software.
  * What if the coordinator crashes?

highly influencial: hadoop, spark...
