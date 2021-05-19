##source
https://pdos.csail.mit.edu/6.824/papers/gfs.pdf  

page: 1~5

## abstract
gfs is a scalable distributed file system for large distributed data-intensive applications, reflect a marked depature from some earlier file system assumtions. met storage needs and fault tolerant.

## introduction
radicall different points:

- fault tolerance:

> component failures are norm rather than exception.
> constant monitoring, error detection, fault tolerance, automatic recovery are integral to the system.

- huge

> multi-GB files are common.

- append data

> most files are appended new data rather than overwitten. Once written, files are only read, and often only sequetially.

- flexibility

> co-designing the applications and file system API

## design
### assumptions

- built from many inexpensive commodity components that ofen fail

- many large files

- large streaming reads and small random reads

- many large, sequential writes appending to files

- must efficently deal with concurrently append by multiple clients

- high sustained bandwidth more important than low lantency

### interface
usual interface(create, delete, open, close, read, write), snapshot and record append.

### architecture
one master with metadata, multiple chunkservers with data accessed by multiple clients. 

Files are divided into fixed-size chrunks. chrunk is identified by an immutable and global unique 64 bit. Each chunk is replicated on three chunkservers defaultly.

client and chunkserver don't cache because of huge stream.(but cache metadata).

### single master
client request to master, cache location of replicas and index, send request to most likely closest replica in subsequent requests.

### chunk size
64MB

- pros:

> reduce contact between clients and master

> perform many operations on a given chunk reduce network overhead by keep persistent TCP connection

> reduce size of metadata in on master, allowing keep metadata in memory

- cons:

> hot spots caused by small file ???
> (potentially solved by reading data between clients)

### metadata
three major types:

- file and chunk namespaces
> persistent by operation log stored by master local disk and replications on remote machines

- mapping from files to chunk
> same to namespaces

- location of each chunk's replicas
> not persistent, ask chunkserver about chunk at master startup and chunkserver new join

#### in-memory data structures
> less than 64 bytes of metadata for each 64 MB chunk

> file namespace data also require less than 64 bytes per file becase storing file names compactly using prefix compression.

#### chunk locations
> chunkserver has the final word over chunk states, so don't maintain in master

#### operation log
> new checkpoint can be created without delaying incoming mutations. Master switch to new log and creates new checkpoint in seperate thread

> recovery only needs latest complete checkpoint and subsequent log files. older checkpoints and log files can be freely deleted.

### consistency model
#### gurantees by GFS
> after a sequence fo successful matations, mutated file region is guaranteed to be defined and contain data written by last mutation. By:

> a) applying mutations to all replicas in the same order

> b) using chunk version numbers to detect any replica that has become stale 

> A chunk is lost irreversibly only if all its replicas are lost before GFS can react, typically within minutes.

#### implications for applications
> mutate files by appending rather than overwriting

> use checksum when concurrently append to a file


## terms
anticipated 预期的

radically 根本地

mutate 改变

onerous 繁重的

sustained 持续的

stringent 严格

respectively 分别地

flaky 片状

orphaned 孤儿

migration 移民

sophisticated 复杂的

bottlenect 瓶颈

fragmentation 碎片化

stagger 错开

eternally 永恒的

catastroph 灾难

exclusively 只

mingle 混合

dwarf 矮人

purge 清除

permanent 永恒的

semantic 语义学
