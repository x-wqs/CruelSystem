## Rest and RPC
Servers expose API over the network, the APIs are called service.

A server can also be a client to another services, e.g. a web server is a client of DB server.

Service-oriented-architecure (SOA), aka, microservices. vs. Monolithic services.

Services are like databases, they allow clients to submit and query data.

REST and SOAP for web services
- REST is a design philosophy, not a protocol. RESTful APIs, GET/PUT/POST/DELETE

Remote procedure call (RPC)
- try to make a request to a remote network service look the same as calling a function or method in your programming language, within the same process. (location transparency)

RPC is fundamentally flawed?
- local func call is predictable, fail/success
- returns or raise an exception
- idempotence if you retry a failed network request
- local running time is roughly the same when run multiple times.

Direction of RPC
- binary encoding RPC have better performance than JSON over REST.
- RESTful API is good for experimenting
** The main focus of RPC frameworks is on requests between services owned by the same organi‐ zation, typically within the same datacenter.**
