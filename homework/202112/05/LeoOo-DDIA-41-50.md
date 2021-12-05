
Data locality for queries
- NoSQL stores all releveant information for an item in the same document, which provides datalocality for queries.
  - it only applies if you need a large part of the document at the same time.
  - not the case if you only need a small part of it.
  - also need to rewrite the entire document upon update.
 
 Query language for Data
 - Declarative: SQL is declarative: you don't need to describe how to get a job done, you need to describe what you need.
 - Imperative: other languages like python

MapReduce Querying
- a programming model for proessing large amounts of data.
- map:
  - breakdown a large input and generate a collection of a k-v pair for the next step, i.e. reduce.
    - emit(key, value)
- reduce:
  - compute all the emitted k-v pairs
  - for all k-v pairs with the same key, the reduce function is called once.
  - the result can be save to a file or somewhere else.
- Graph [TODO]
