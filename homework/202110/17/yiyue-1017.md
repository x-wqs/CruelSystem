# Design a unique ID generator in distributed system

## Step 1 - problem & design scope

sample questions to ask the interviewer
- characteristics of unique IDs? Unique and sortable
- For each new record, ID increments by 1? increments by time but not necessarily only 1
- IDs only contain numeric values? yes
- ID length requirement? 64-bit
- Scale of the system? 10000 IDs per second

## Step 2 - High-level design and buy in

### multi-master replication
- increase the next ID by k
- k - the # of database servers in use
- Drawbacks
 - Hard to scale with multiple data centers
 - IDs do not go up with time across multiple servers
 - Not scale when a server is added / removed

### UUID
- Each server contains an ID generator
- Web server responsible for generating IDs independently
- pros
 - simple
 - easy to scale
- cons
 - IDs are 128 bits, but requirement is 64 bits
 - not go up with time
 - non-numeric

### Ticket server
- use a centralized `auto_increment` feature in a single databse server
- pros
 - numeric IDs
 - easy to implement
- cons
 - single point of failure

### Twitter snowflake approach

- Sign bit - 1 bit. Always be 0.
- Timestamp - 41 bits.
- Datacenter ID - 5 bits
- Machine ID - 5 bits
- Sequence number - 12 bits

## Step 3 - Design deep dive

- Datacenter IDs and machine IDs chosen at the startup time, fixed once the system is up running
