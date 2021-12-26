<a name="XYCS3"></a>

previous reading progress (P141-P151): https://github.com/refinedcoding/CruelSystem/blob/main/homework/202112/23/monsooooon.md
this file: P151-155

# overview of replication
> Replication means keeping a copy of the same data on multiple machines that are connected via a network.

<a name="qMvYH"></a>
## why replication?

- reduce data access latency
- fault tolerance/HA
- scale out for higher r/w rates/throughput
<a name="Wjkj0"></a>
# single leader based approach
<a name="Pd58g"></a>
## overview

- data is stored in multiple nodes. 1 -> leader, others -> replicas, followers
- req handling process:
   - client send req to leader ONLY
   - leader save it, and replicate it to all other replicas for updating
   - client reads from all nodes (both leader & replicas)
- this process is adopted by many db, since it's easy (mysql, postgresql, mongodb...)
<a name="jaNoc"></a>
## sync or async replication

- replication latency btw leader & replicas
- approaches
   - **sync**:  the leader waits until followers have confirmed that they received the write before reporting success to the user, and before making the write visible to other clients.
   - **async**: the leader sends the message, but doesn’t wait for a response from the follower.
   - **semi-sync**: the leader waits for at least N succ responses from all followers,  and then reporting success to the user



<a name="O0Q1N"></a>
# chain replication
[https://www.youtube.com/watch?v=IXHzbCuADt0&t=3459s](https://www.youtube.com/watch?v=IXHzbCuADt0&t=3459s)<br />![image.png](https://cdn.nlark.com/yuque/0/2021/png/25679250/1640509254930-c9561236-a756-4a1b-8068-f8234a6fd09c.png#clientId=uf9e3ac37-306f-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=328&id=u14abe8be&margin=%5Bobject%20Object%5D&name=image.png&originHeight=656&originWidth=1462&originalType=binary&ratio=1&rotation=0&showTitle=false&size=825562&status=done&style=none&taskId=uc7e0d5f3-3d64-462e-92f1-ab1ae6a0809&title=&width=731)

- orgnize nodes/servers into a chain
- one head & one tail node
- write
   - client send writes to the head to update data
   - head update local storge & pass write to next node, until the tail
   - tail node responds succ to client
- read
   - client only send read req to tail
   - tail node responds to client
- consistency: tail handle all reads & writes, so strong consistency
- failure recovery: 
   - all visible status change only happens at the tail and record by all other nodes
   - if head fails, next node becomes head
   - if tail fails, its pre-nodes takes over
   - if middle nodes fails, re-connect linked list
- compares to Raft:
   - raft leader has large write load (send req to all followers)
   - CR: only pass write req to next node
   - CR: suffers from slow write speed, O(n) with #replicas
- configuration manager (CM):
   - monitor servers/nodes liveness
   - if CM find some nodes looks dead, it change every nodes' conf
   - CM uses Raft, Paxos, ZK to ensure fault tolerance
