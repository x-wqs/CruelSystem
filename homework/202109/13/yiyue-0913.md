## Database replication
- master/slave relationship - original (master) the copies (slaves)
- master database only supports write operations
- slave database gets copies of the data from the master database, only supports read operations
- data modifying commands (insert, delete, update) sent to master database
- database replication advantages
  - better performance
    - writes and updates - master
    - read - slave nodes
    - more queries processed in parallel
  - reliability
    - one server destroyed, data still preserved
    - no data loss
  - high availability
    - website remains in operation even if one database offline
- if one database offline
  - only one slave database, goes offline. Read -> master
  - new slave database replaces the old one
- if master offline
  - a slave will be promoted to master
  - all operations will be executed on new master
  - new slave replace old one for data replication
  - in production system, the promotion will need to run data recovery scripts

## cache
- temporary storage area that stores result of expensive responses / frequently accessed data

## cache tier
- temporary data store layer, faster than database
- benefits - better system performance, reduce database workloads, scale the cache tier independently
- read-through cache
  - request -> web server -> checks if cache has response
  - if yes, data -> client
  - if not, queries database -> stores response in cache -> client
## considerations
- when to use
  - data read frequently but modified infrequently
  - not ideal for persisting data, no important data stored
- expiry policy
  - expired -> remove from cache
  - if too short - reload data from database too frequently
  - if too long - stale data
- consistency
  - data store and cache in sync
- mitigating failures
  - multiple cache servers across different data centers are recommended
  - overprovision the required memory - provides buffer as memory usage increases