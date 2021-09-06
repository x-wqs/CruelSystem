### 5.5 Follower and candidate crashes

- `RequestVote` and `AppendEntries` will fail
- Retry indefinitely
- If RPC completion is done before server crashes, same RPC will be received after restart

### 5.6 Timing and availability

- Safety doesn't depend on timing
- availability depends on timing
- Timing requirement: broadcastTime << electionTimeout << MTBF
  - broadcastTime = avg(time to send RPC and receive response). Properties
  - electionTimout: we choose
  - MTBF = avg(time between failures). Properties

### 6 clusterm membership changes

- Safe configuration change mechanism: same term only one leader
- Two-phase approach
  - Switch to transitional configuration - join consensus - combines both old and new configurations
  - new configuration
- Issues
  - New servers added to the cluster won't be able to vote. And catchup happens before the 2-phase
  - cluster leader might return to follower state when committing Cnew
  - removed servers will time out and do elections many times
    - to resolve, a server doesn't respond within minimum election timeout when receiving RequestVote RPC

### 7 Log compaction

- Logs grow 
- Snapshotting to compaction
  - write current system state to a snapshot
  - previous log will be discarded
- Incremental approaches
  - select a region of data which has many deleted and overwritten objects
  - rewrite live objects compactly and free the region
- InstallSnapshot RPC to update far behind followers
  - snapshot conatins new information: follower discards entire log and use snapshot
  - snapshot describes prefix of its log: follower deletes entries covered by snapshot but retains entreis following the snapshot
- If only leader allows to send snapshot
  - waste network bandwith and slow process
  - leader implementation is more complex
- Decide when to snapshot
- Writing snapshot can't delay normal operations
