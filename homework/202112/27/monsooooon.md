Chapter 5, Replication
p155-p161

# chain replication review
![diagram](https://liqul.github.io/blog/assets/chain_replication.svg)
- a sync way to replicate data
- adopted in Azure Storage
- pros: 
  - strong consistency
  - easy to build
- cons:
  - heavy load on tail
  - need a reliable master
- extension: CRAQ for read heavy scenario  
![diagram](http://3.bp.blogspot.com/-9bBdliYUy74/TWAyS3oNuWI/AAAAAAAAA6A/R3gsAOof5RI/s400/fignew.png)

# setting up new followers
how to ensure new followers have up-to-date data?
- by locking db & migrate: cause downtime, not acceptable
- by taking consistent snapshot + sync delta modifications

# handling node outages
- follower failure: 
  - catch-up recovery: as described above
  - based on follower local log & leader log diff
- leader failure:
  - failover: promote new leader & reconfig the cluster
  - detecting leader failure: no foolproof way
  - choosing a new leader: consensus algo
  - reconfig followers & clients: need to avoid split brain
- possible issues when failover
  - new leader may not be up-to-date in async replication
  - dangerous to discard user input data when log conflicts 
  - timeout for detecting leader failure should be appropriate
- very complicated fields, need more discussion
  
# implementation of replication logs
## statement-based replication
- SQL level statement
- need to handle non-deterministic func: RAND, NOW
- need to guarantee statements order (may have dependencies)
- side effects (UDF, triggers...) 

## WAL-based replication
- WAL: works for LST-Tree & B-Tree
- descibe data changes at disk level
- hard for diff software/versions

## row-based log replication
- btw SQL & WAL level, record table row changes
- logical log:
  - insertion: new row values
  - deletion: PK for row to delete
  - updating: PK & new row values
- friendly to backward compatibility for newer version software
- easier to parse for external softwares (__CDC, change data capture__)
  - dataware house for OLAP purpose
  - custom caches & indexes
