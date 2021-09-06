6.824: Distributed Systems Engineering

## What is a distributed system?
  multiple computers

## Why do people build distributed systems?
  > to increase capacity via parallelism: horizontal scaling    
  > to tolerate faults via replication: make copies   
  > to place computing physically close to external entities: locality    
  > to achieve security via isolation: ???    

downsides: 
  > concurrent parts, complex interactions   
  > partial failure   
  > tricky to realize performance potential   

##  Main topics
heavy on fault-tolerance/storage; touches on communication & computation    

> Lab 1: MapReduce   
> Lab 2: replication for fault-tolerance using Raft   
> Lab 3: fault-tolerant key/value store   
> Lab 4: sharded key/value store   

The big goal: abstractions that hide the complexity of distribution.   

### Topic: fault tolerance
  > We often want:   
    1. Availability: less down time   
    2. Recoverability     
  > Big idea: replicated servers.   
    1. If one server crashes, use the other(s).    
    2. Very hard to get right: server marked as crashed might still serve requests   
  > Labs 1, 2 and 3   

### Topic: consistency
   > Read after write    
   > "Replica" servers should be kept identical.    


### Topic: performance
  > The goal: scalable throughput   
    Nx servers -> Nx total throughput via parallel CPU, disk, net.   
  > Scaling gets harder as N grows:
    1. im-balance load, stragglers, outliers
    2. Non-parallelizable code: initialization interaction   
    3. Bottlenecks from shared resources: network   
  > Some performance problems aren't easily solved by scaling   
    1. e.g. smaller latency for a single user request   
    2. e.g. all users want to update the same data   
  > requires better design vs more computers   
  > Lab 4

### Topic: Fault-tolerance, consistency, and performance are enemies.   
  > Strong fault tolerance requires communication   
    e.g., send data to backup    
  > Strong consistency requires communication,   
    e.g. Get() must check for a recent Put().   
  > Many designs provide only weak consistency, to gain speed.   
    e.g. Get() does **not** yield the latest Put()!   
    **trade-off**.
  > Many design points in the consistency/performance spectrum!
