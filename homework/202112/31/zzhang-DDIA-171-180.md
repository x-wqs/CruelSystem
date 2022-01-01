# DDIA: pp 171-180

##### Collaborative editing

- When one user edits a document, the changes write to their local replica
- Async replicate to the server and other users

No-editing-conflicts guarantee: use lock => equivalent to single-leader replication with transactions on the leader. Might be slow.

Fast editing: => need to handle write conflicts



#### Handling write conflicts

Single-leader: writes in a sequntial order. The last write determines the final value

Multi-leader: no defined ordering of writes. => the database resolve the conflict in a *convergent* way.

- Give each write a unique ID, pick the write with the highest ID as the *winner*. If ID is related to timestamp => *last write win (LWW)* => data loss problem
- Give each replica a unique ID and let write from higher-numbered replica always take precedence over writes from lower-numbered replica => also data loss problem

- Merge the values together
- Record the conflict that preserves all information



##### Custom conflict resolution logic

- *on write*: As soon as sys detect a conflict in the log of replicated changes, it calls the *conflict handler*.
- *on read*: When a conflict is detected, all the conflicting writes are stored. When read, return to user or resolve the conflict and write the result back to the db.



### Automatic Conflict Resolution

- *Conflict-free* replicated datatypes (CRDTs): use data structures that can be concurrently edited by multiple users
- *Mergeable persistent data structures*: tack history explicity, like Git
- *Operational transformation*: for concurrent editting of an ordered list of items, such as the list of characters that constitue a text document. application such as Google Docs. 



Conflict examples:

- One unique item is assigned by more than one suer
- Multiple users change the value of the same field.



## Multi-leader Replication Topologies

Replication topology in multi-leader replication:

- Examples: 1) circular topology (MySQL); 2) star topology; 3) All-to-all topology (Most general)



To prevent infinite replication loop:

- Each node is given a unique identifier; in the replication log, each write is tagged with the identifiers of all the nodes it passed through; When a node receives a data change that is tagged with its own identifier, that data change is ignored.



Problems with circular and star topology: 

- if one node fails, changing the topo is costly.c



Problems with all-to-all topology:

- Causality: the update depends on the prior insert, we need to make sure that all nodes process the insert first, and then the update. => attach a timestamp is not sufficient, since closkcs cannot be trusted to be sufficiently in sync to correctly order these events.
- Use *version vectors*. 



##### Leaderless Replication

- Allow any replica to directly accept writes from clients
- *Dynamo-stype* db



##### Writing to the database when a node is down

Leaderless replication: failover does not exist.

- Write is successful when part of replicas have acknowledge the write.
- When read, read reqs are sent to several nodes in parallel.
  - *Read repair*: when a client read from several nodes in parallel, it detect any stale responses. If it gets different version of data, write newer value back to staled replicas. This method works well for values that are frequently read
  - *Anti-entropy process*: there is a background process that constantly looks for differences in the data between replicas and copies any missing data from one replica to another. Doesn't copu writes in any particular order.
    - For values that are rarely read, anti-entropy process is a must.



Quorums reads and writes:

- if there are n replicas, every write must be confirmed by *w* nodes to be considered successful, and we must query at least *r* nodes for each read. As long as *w+r>n*, we can get the up-to-date value when reading.



Typical choice: n is odd (typically 3 or 5). Set w = r = (n+1)/2 (rounded up). 

