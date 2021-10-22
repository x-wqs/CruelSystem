### Timestamp

- Use 41 bits to represent timestamp

### Sequence number
- 12 bits
- 0 unless more than one ID is generated in a millisecond on the same server

## Wrap up

- Addtional talking points
 - clock synchronization
  - we assume ID generation servers have the same clock
  - Might not be true when server running on multiple cores / multi-machine
  - solution - Network Time Protocol
 - section length tuning
 - high availability

# Design a URL shortener

## understand the problem and establish design scope

- questions to clarify
 - example of how url shortener work
 - traffic volume - 100 million URLs generated per day
 - How long is the shortened URL - as short as possible
 - characters allowed in the URL
 - delete / update URL? - cannot
- use cases
 - URL shortening - given a long URL -> return a short URL
 - URL redirecting - given a short URL -> redirect to original
 - High availability, scalability, and fault tolerance
- assumptions
 - write operation - 100 million URLs generated per day
 - write operation per second - 1160
 - read operation - assuming ratio of read operation to write is 10:1, read operation per second 11600
 - Assume run for 10 years - support 365 billion records
 - Average URL length - 10
 - Storage requirement for 10 years - 365TB

## Propose high-level design and get buy-in

### API endpoints
- facilitate the communication between clients and servers
- REST-style
- two API endpoints
 - URL shortening
 ```java
 POST api/v1/data/shorten
 ```
 parameter: `{longUrl:longURLString}`
 return: shortURL