# DDIA: pp 131-140

## Dataflow Through Services: REST and RPC

*service-oriented architecture (SOA)* or *microservices architecture*: Decompose large application into smaller services with different functionality => one service make request to another if it requires different func or data from others.

- Design goal: make services **independently** depolyable and evolvable.



Web service:

- Use HTTP as the underlying protocol for talking to the service

Examples:

- Client applications make requests over HTTP.
- Inside datacenter, one service make requests to another service. *Middleware*: Software supports this kind of use case.
- Data exchange between different organization's backend systems: includes public APIs provided by online services.

Approaches to web services: *REST* and *SOAP*.

**REST**: design philosophy: simple data formats, using URLs for identifying resources, use HTTP features for cache control, authentication and content type negotiation. 

**SOAP**: XML-based protocol. The API of a SOAP web service is described using WSDL (Web Services Description Langugae). Too complex and WSDL is not human-readable. => Not so popular now.

 



**Remote procedure call (RPC)**:

- an idea to make a request to a remote network service look the same as calling a function within the same process.

Difference between network request and local functioncall:

- Network request is unpredictable. Request or response may be lost due to network problems or remote machine's problem.
- Network request might return without a result due to *timeout*. Hard to know whether the request got through or not.
- If response get lost and you retry the request, there might be duplication problem. => could add *deduplication* (*idempotence*) into protocol.

- Latency is widely variable
- All parameters need to be encoded into the sequence of bytes.
- Client and service might use different programming languages.

Current directions for RPC: handle these differences.

- Finagle (using Thrift) and Rest.li (using JSON over HTTP) use *futures* (*promises*) to encapsulate async actions that may fail.
- gRPC (using Protocol Buffers) supports *streams*, where a call consists of a series of requests and responses over time
- *service discovery*: allow a client to find out at which IP:port it can find a particular service.

Data encoding and evolutions for RPC:

- Reaonable to assume: all the servers will be updated first, and all the clients second
- RESTful APIs: use a version number in the URL or in the HTTP *Accept* header.



## Message-Passing Dataflow

Async message-passing systems:

- send message through a *message broker* (*message queue*, *message-oriented middleware*)

- One-way sending: the sender doesn't wait for the message to be delivered and receive a reply.

Broker Procedure:

1. One process sends a msg to a named queue or topic
2. the broker ensures that the msg is delivered to one or more *consumers* of or *subscribers* to that queue or topic.

* consumer can itself publish msg to another topic, e.g., reply queue that is consumed by the original sender.



### Actor model

- a programming model for concurrency in a single process.

- Each actor represents one client with local state. Actors can communicate with each other by sending and receiving async msgs. Each actor can be scheduled independetly.

Distributed actor framework:

- Messages might sent over network => may be lost.

Three popular distributed actor frameworks: => handle msg encoding

- *Akka*: Use Java's build-in serialization (by default), or Protocol Bufferers

- *Orleans*: custom data encoding format without rolling upgrade deployment support (by default). To deploy new version of application: set up new cluster, move traffic to the new one, shut down the odl one.
- *Erlang OPT*:  hard to make changes to record schemas. Rolling upgrades are possible but need to be planned carefully.

## Summary

Rolling upgrades:

- new versions of a service to be released without downtime and make depolyments less risky.

backward compatibility:

- new code can read old data

forward compatibility:

- old code can read new data

