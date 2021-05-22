## source
mit 6.824

https://www.google.com/url?q=https://youtu.be/6ETFk1-53qU&sa=D&source=editors&ust=1621657447198000&usg=AFQjCNGnMQpugFqtEMi_i6AxEpv1kEY6rA

## plan
- storage

- gfs

- consistence

## storage system
> building fault tolerant storage

> app = state less

> storage holds persistent state

## why hard
> high performance --> shared data across servers

> many servers --> constant faults

> fault tolerance --> replication

> replication --> potential inconsistences

> strong consistencies --> lower performance

## ideal consistency

> behave as if single system

> bad replications without protocal cause
 inconsistency
  
## gfs: a case study

### standard problems
> fault tolerance , performance, consistency

### special of gfs
> 1000+ machines

> one master
 
> in-consistencies

## key properties of gfs
> big: large data set

> fast: automatic sharing

> global: all clients see same

> fault tolerance: automic

## master
> file name --> array of chunk handles

> chunk handle --> version numbers(list of chunk servers, one primary and seconderies, lease time)

> log + checkpoints

## read a file
> client sends filenames to master

> master to client: chunk handles + list of chunk servers + version numbers

> client caches list

> client reads from closest chunkserver

> server checks version number, if ok, send data

## write
> client asks master chunk handles

> client use handles finding servers

> client push data to closest chunkserver, and then servers spread data one by one

consistency
> version number only increases when select a new primary

> when primary dies, it is called split brain when two primary

## how to get stronger consistency
> update all p + s or none

> google: build additional systems


## terms
constitute 组成

bizarre 怪异的

garbled 断章

vague 模糊的

syndrome 综合症
