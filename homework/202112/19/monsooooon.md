ddia 
# column oriented storage overview
- how to store large fact table? may be 100 cols & PB size
- every query only accesses small num of cols (~5), ignore others
- may calc agg. (sum/max/count) on some cols
- __IDEA__: do not load all rows, load only queried cols
- Parquet: column storage format for non-relational data, from Dremel
- every cols has records with the same order

# column compression
- compress data to reduce disk IO, but increase CPU
- bitmap encodings
  - small distinct data, map binary val -> real val
  - `for col in (typeA, typeB, typeC)`, use OR to calc union
- __column families__:
  - from bigtable, used in Cassandra & HBase
  - store all columns from a row together, along with a row key
  - usually do not use column compression, thus mostly row-oriented.
- __WHY COMPRESSION__:
  - fast loading from disk->mem->CPU
  - small enough to fit into CPU cache(L1, or higher)
  - making use of SIMD (vectorized processing)

# sort order storing policy
- how to store rows? random(append-only) or sorted (like SST)
- sort all rows based on __sort key__ (1 or multiple)
- friendly for raneg queries, such as time range
- possible scenario: sort by time first, then sort by prod_id
- cols with sorted key can be easily compressed
- use __redundant__ data to provide multi sorted order
  - like secondary index in MySQL
  - increase write overhead, only for heavy read cases

# writing to column-oriented storage
- B Tree approach: not possible with compressed data
  - > if you wanted to insert a row in the middle of a sorted table, you would most likely have to rewrite all the column files.
- LSM-Tree approach: in-mem store & on-disk sst
  - > All writes first go to an in-memory store, where they are added to a sorted structure ... When enough writes have accumulated, they are merged with the column files on disk and written to new files in bulk.

# data cubes & materialized views
- materialized aggregations: pre-calc & pre-updated agg. funcs
  - COUNT, SUM, MAX, MIN, etc.
  - like a cache for frequent query patterns
  - increase write overhead, reduce read latency, useful for OLAP
  - data cube: pre-agg on multi-colunns
    - like col1 & col2, form two diamensional tbl
    - given one diamension, easily calc agg on another diamension
    - `select sum(price) from fact_tbl where col1 = val1`
    - `select sum(price) from fact_tbl where col2 = val2`
    - with more cols(diamensions), need more agg. results
    - not flexible for filtering out part of rows

# encoding & evolution overview
how to handle evolution of data models? like new fields
- forced schema: relational db
- schemaless/schema-on-read: non-relational, like documented based db
- newer and older ver of data format may coexist:
  - rolling udpates for internal sys
  - app code: user may not install latest ver
- __data format compatibility__ is required
  - backward: Newer code can read data written by older code.
  - forward: Older code can read data written by newer code.

# data encoding formats
- encoding/decoding: in mem data structure <-> on-disk files
- common std. codec format:
  - textual: JSON, XML, CSV
  - binary variants: pb, thrift
## textual encoding:
- problems with textual encoding:
  - ambiguity for numbers (int? long? float? double?)
  - poor support for binary strings, may use base64, size↑ by 33%
- popularity is WAY more important than any other shortcomings
## binary encoding:
- suitable for in-org data exchanges
- binary version of JSON: MessagePack, BSON ...

# thrift & pb
- binary format, use IDL to describe format
- use code gen tools to produce code for r/w binary data
- use Tag Number to 
- field tags & schema evolution

# mode of dataflows overview
- with codec, we will explore ways how data flows btw processes
  - via database
  - via service calls
  - via async message passing

# dataflow through databases
- important to have backward compatibility, because software will upgrade
- forward compatibility may also be required, because old service instance may read data writen by new ones
- how to handle new tags/fields? 
  - unknown for old code, should stay intact when r/w new data
- data outlives code: code changes but data stay unchanged in db
  - migration data? possible but costly
  - adding null value to odd data record when reading

# dataflow through services
- REST or RPC
- case of web:
  - client: browsers/mobile devices/PC, servers: web application servers
  - basic aspects: HTTP, URLs, SSL/TLS, HTML 
  - use GET/POST to fetch/update data
- web services: API implemented via HTTP
## REST: a style for HTTP API, not a protocal
  - use URL to locate resources
  - use HTTP method to perform actions
  - How to describe and documente RESTful API? OpenAPI(Swagger)

## RPC: try to make request to other service via 'local' method
- location transparency: local func -> network -> remote func
- drawbacks:
  - local func call is stabler than remote calls
  - idempotence: remote calls may fail, but succ or fail is not clear
  - remote calls can be very slow and need timeout for giving up
  - need marshalling for data with mem pointers
- possible to build RPC Lib on top of REST. __REST doesn't try to hide the fact that other services is remote, but RPC does.__
- RPC framework: for making RPC calls, gRPC, Thrift...

Page 135
To Be Cont.
