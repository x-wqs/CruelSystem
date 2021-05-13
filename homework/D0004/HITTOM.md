## plan
- what is distributed system
- historical context
- course structure
- main topic
- map reduce

## what is distributed system
- multiple network
- cooperating computers
### four main reasons for using distributed system
- connect physically seperated machines
- increase capacity through parallism
- tolerate fault, high avalibility
- achieve security via isolate
## historical context
- local area networks(1980s) DNS + email + AFS
- data centers, big websites(early 1990s) websearch, shopping
- cloud computing
- current state: active
## challenge
- many concurrent parts
- must deal with partical failure
- tricky to realize the performance benefit
## why take 6.824
- interesting: hard problems and powerful solutions
- used in real word
- active area of research
- hands on experience

## course structure
- lectures: basically focus on big ideas
- papers: case study for a paticular big idea(strong recommand read before class)
- labs: 
> 1) mapreduce
  2) replication using raft
  3) repulicated k/v service
  4) sharded k/v service
  5) optional project with group
  test cases are public

- exams: half semester and final

## focus: infrastructure
- storage
- complutaiton
- communication 6.829
- rpc
- abstractions

## main topics
- fault tolerance:
> 1) avalibility replication
  2) recoverablity loggin/transaction, durable storage

- consistency：
  does get() return the value of the last put()

- performance
>  1) throughput
   2) tail latency

- implementation

## map reduce
### why learningmap reduce
- illustration
- influential
- lab1
multi-hours computations of terabytes data

- goal: 
easy for non-experts writing distribute applications

- application: 
> 1) map/reduce sequential code
  2) mr framework deals with distributions


## abstract view
when all mappers done, reduces contacts to each mapper to know their data locations

no commucation happens in mapper(gfs to local disk)

commucations happen in reducer(mapper disk to reducer machine cross network, out to gfs)

mapreduce library do sorting, then hands it to reduce function

## fault tolerance
- workers fail:
> cordinator reruns map/reduce
  mapper can run twice(functional and deterministic)
  reducer can run twice(just rename)

- cordinator cann't fail

- slow workers: 
  backup, tackle swagglers
