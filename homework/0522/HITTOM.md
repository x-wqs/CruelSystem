## source


> https://pdos.csail.mit.edu/6.824/papers/vm-ft.pdf 1~5

> The Design of a
Practical System for
Fault-Tolerant Virtual Machinesmit 6.824

https://www.google.com/url?q=https://youtu.be/6ETFk1-53qU&sa=D&source=editors&ust=1621657447198000&usg=AFQjCNGnMQpugFqtEMi_i6AxEpv1kEY6rA

## plan
- storage

- gfs

- consistence

## abstoract
> describe basic design, alternate design choices, implementation detail of our practical system.

## introduction
> common approach:
secondaries keep indentical with primary nearly on the time, shipping all state of CPU, memory, I/O devices cost large bandwidth

> our approach:
model servers as deterministic state machines, use extra coordination to ensure primary and backup kept in sync. extra informatige system
> building fault tolerant storage

> app = state less

> storage holds persistent state

## why hard
> high performance --> shared data across servers

> many servers --> constant faults

> fault tolerance --> replication

> replication --> potential incons is far less than total state.

> extra data is needed for non-deterministic operations, hypervisor capture all these extra informations.

> deterministic reply and FT support only in uni-processor, multi-processor still work in progress, as access to shtences

> strong consistencies --> lower performance

## ideal consistency

> behave as if single system

> bad replications without protocal cause
 inconsistency
  
## gfs: a case study

### standared memory is a non-deterministic operation

> only deal with fail-stop failures, avoid incorrect externally visible action before server failures detected

## basic ft design
> 

## terms
> hypervisor 管理程序

> uni-processor 单处理器problems
> fault tolerance , performance, consistency

### special of gfs
> 1000+ machines

> one master
 
> in-consistencies

## key properties of gfs
> big: large data set

> fast 






## terms

