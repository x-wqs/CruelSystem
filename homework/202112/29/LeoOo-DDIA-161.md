
Trigger-based replication
- trigger lets you register custom application code that get automatically executed whe na data change (write transaction).
  - trigger could use this to log the change into a separate table for external process to apply application logic and replicate the data change to another system.

Problems with replication lag
- fault tolerant, scalability and latency are the reasons for replication.
- read-scaling: increase the capacity for serving read-only requests by adding more followers. Async replicate.
  - Sync replicate would be more likely to fail.
  - Async cause outt of date info if the follower fall behind. Inconsistency.
  - Eventual consistency, if the writes pause and wait a while the followers will catch up and become consistent with leader.
  - 'Eventual' is vague, it can be half second or longer than tolerable if system is fully loaded or have a high network traffic.

Reading your own writes, read-after-write consistency, aka read-your-writes consistency
- Let user submit some data then view what they have submitted.
  - but no promise to other users.
- when read something the user may have modified, always read it from the leader, o/w, read it from a follower.
  - user's profile can only be modified by the user => read user's own profile from the leader, read other's profile from follower.
- track the time of the last update and read from the leader in the following minute after the write.
  - could also monitor the replication lag on followers and prevent query any follower that has more than 1 minute lag.
- client remember the timestamp of its most recent write - the system make sure to read from a replica that is updated at least until that timestamp.
- cross-device read-after-write consistency
  - remembering user's last update timestamp become more difficult.
