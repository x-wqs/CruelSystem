## Implementation of Replication logs

**Statement-based replication**

too many problems, used in MySQL before v5.1

**Write-ahead log (WAL) shipping**

an append-only sequence of bytes containing all writes to the database

describes the data on a low level and closely coupled to the storage engine

**Logical (row-based) log replication**

usually a sequence of records describing wirtes to databases tables at the granularity of a row

be decouple from storage engine, thus ensures backward compatibility, and also easier for external apps to parse

used by binlog of MySQL (when configured)

**Trigger-based replication**

when you need move replication up to the application layer

flexible, but more overheads and error-prone



## Problems of Replication Lag

In **Leaders and Followers** replication model, if writing in a sync manner, a single node failure would make the entire system down, if in a async manner, there's problem called **replication lag**

### Reading Your Own Writes

**read-after-write-consistency:** if the user reload the page, they'll see the update immediately

### Monotonic Read

It's possible for user to see things moving backward in time

**eventrual consistency** < **monotonic read** < **strong consistency**

one way to achieve this is to make sure each user alaways reads the same node

### Consistent Prefix Reads

one solution is to make sure that casually related writes are written to the same partition

### Solutions for Replication Lag

**transactions:** a way for a database to provide stronger guarantees so that the apps can be simpler