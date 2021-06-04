# Lecture1 Notes

# Why we need Distributed System?

* connect physically seperated machines
* increase capacity
    * mapreduce
    * zoom: zoom sessions going on at the same time and zoom.com has to support it all. 
    It requires a lot of computers to basically increase the capacity so you can support all those in parallel zoom sessions
* tolerate faults 
* archieve security via isolate

## Historical context

* local area networks(1980s)
    * No Large Scale internet applications. DNS + Email + [AFS (Andrew File System)](https://en.wikipedia.org/wiki/Andrew_File_System)
* data centers, big websites (1990s)
    * web search, shopping leading to 
        * huge datasets(reverse index of webs)
        * lots of users
    * new solution is needed!(mapreduce...)
* cloud computing(2000s)
* current state (activate!)

## Challenges

* many concurrent parts
* must deal with partial failure
* tricky to realize the performance benefits

## Why take this course?

* interesting: hard problems but powerful solutions
* used in real word
* active research area:  many open problems
* hands-on: unique style of programming, builds up another type of skill of programming that you might not have done

## Course Structure

* Lectures about big ideas, papers, labs
* Readings: research papers as case studies
    * please read papers before class
* Labs: 
    * mapreduce
    * replication using raft
    * repulicated k/v service
    * sharded k/v service
    * optional project with group test cases are public

## Focus: infrastructure 

* storage
* computation
* communication 6.829
    * rpc
* abstractions
    * when we build a storage system we want our basically the stupid storage system more or less behave like you know a single machine sequential uh storage server like your regular file system on your laptop


## main topics

* fault tolerance
    * availability
    * recoverability(logging and transaction)

* consistency
    * Example: does get() return the value of the last put()
    * There are more than one consistency.

* performance
    * throughput
    * latency

Achieving these sort of free things at the same time uh it turns out to be extremely difficult and in fact what people do in practice is they make different trade-offs.

## MapReduce
## Motivation

    Google has lots of web data(terabytes of data). And google needs to do inverse index.

    They need a system which can handle parallelism as well as friendly to nonexpert.

## Approach
* map + reduce-> User writes sequential codes
* The framework deals with all the distributeness
