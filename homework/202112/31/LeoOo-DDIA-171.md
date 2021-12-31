
## Handling write conflicts

Concurrent-update could lead to conflict in a multi-leader replication setting.
- With single-leader, the 2nd write will be blocked or it will wait for the frist write to complete.
- With multi-leader, both writes are successful and the conflict is only detected asyncly after some time. It might be too late to reolsve the conflict. Block and wait will cost the performance -- might as well just use single-leader replication.

### Conflict avoidance
- make sure all writes for a particular record go through the same leader.
  - frequently recommended approach.
- but still can be problematic when we have to change leaders.

### Converging toward a consistent state
- Single-leader: last write wins.
- Multi-leader: all replica must arrive at the same final value when all changes have been replicated.
  - give each write a unique ID (a timestamp, a random number, a hash, UUID), the highest ID win.
  - data loss prone
- Given each replica a unique ID and higher ID replica wins
- Somehow merge the values together.
- Record the conflict in an explicit data structure that preserves all information.

### Custom conflict resolution logic
Application builder provid logic to resolve conflicts, apply such logic
- on write: system call the logic as soon as it detect a conflict.
- on read: conflicting writes are stored and resolved next time the data is read.

### Automatic conflict resolution

- conflict-free replicated datatypes (CRDTs) are a family of data structures for sets, maps, etc. that can be concurrently edited by multiple users and automatically resolve conflicts in sensible ways. (two-way merge)
- mergeable persistent data structures track history explicitly and use a three-way merge.
- operational transformation: designed particularly for concurrent editing of an ordered list of items. Used by google docs.
