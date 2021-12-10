# Storage and retrieval

Fundamentally, a database should do two things: 
- given the data, store the data.
- asked about the data, return the data.

## Index
- metadata that are kept on the side and acts as a signpost and helps you to locate the data (on the disk) you want.
- overhead on writes.
  - hard to beat simple appending to a file.
  - databases don't index everything by default, developer/user can choose indices manually.

### Hash index
- like a hashmap, but instead of k-v pair, we store key-offset pair so that we can quickly locate the value of a key in the file.
  - the value can be a complex object, e.g. nested json.
  - Bitcask (default storage engine) from Riak (subject to the constraint that all keys must reside in memory).
  - works great ** if value for each key is updated frequently **. e.g. number likes of a video.
- append-only writes
  - how to avoid using too much disk space?
    - break the log into segments when it reaches a certain size.
    - make subsequent writes to a new segment file.
  - compaction 
    - throw away duplicate keys in the log and only kep the most recenlty updated value of the key.
  - compaction and merging
    - compact => smaller files
    - compact merge multiple files
      - with a background thread
      - can serve read/write with old files
    - read from merged file
    - delete the old file.
  - each segment has its own in-memory hash index.


