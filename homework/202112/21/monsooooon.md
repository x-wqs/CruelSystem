ddia

# message passing dataflow overview
- async message-passing: 
  - use intermediary called a message broker/queue, which stores the message temporarily
  - benefits:
    - buffer for recver to avoid overloaded
    - auto redeliver msg when recver crashes/network goes wrong
    - decoupling senders & recvers  
  - one way communication

# message brokers
- one process sends a message to a named queue or topic, 
- the broker ensures that the message is delivered to one or more consumers of that queue or topic.
- usually don't enforce any particular data model: use what you like for compatibility

# distributed actor framework
- __actor model__: a programming model for concurrency in a single process
  - rather than dealing directly with threads, logic is encapsulated in actors.
  - each actor typically represents one client or entity, it may have some local state
  - it communicates with other actors by sending and receiving asynchronous
messages.
  - message delivery is not guaranteed: in certain error scenarios, messages will be lost.
  - each actor processes only one message at a time, it doesn’t need to worry about threads, and each actor can be scheduled independently by the
framework.
- __distributed actor frameworks__: a programming model that is used to scale an application
across multiple nodes.
    - like actor model, sender and receiver do not have to be one the same node
    - message passing is transparent to actors
    - a distributed actor framework essentially integrates a message broker and the actor programming model into a single framework.

# summary
- codec approaches matter: will affect efficiency & software deployment options
- rolling upgrade: need to consider backward & forward compatibility
- encoding formats:
  - language specific encodings: usually bad
  - texutal encodings(JSON, XML): easy to debug, vague
  - binary schema-driven encodings(Thrift, Pb, Avro): compact, efficient, clear compatibility semantics, not directly human readable
  - data flow mode: via RPC/REST, via DB, via MQ
- conclusion: backward/forward compatibility and rolling upgrades are quite achievable.
