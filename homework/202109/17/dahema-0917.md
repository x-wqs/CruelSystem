# chapter 2 Estimation

## power of 2

## Latency numbers
- 1 ns = 10^-9 seconds
- 1 μs= 10^-6 seconds = 1,000 ns
- 1 ms = 10^-3 seconds = 1,000 μs = 1,000,000 ns
- L1 cache: 1ns, L2 cache: 4ns
- memory reference: 100ns
- ssd random read: 16 us
- ssd read 1M bytes sequentitally 49 us
- round trip data cetner: 500 us
- disk seek: 2ms
- disk read 1M bytes sequentitally: 825us
- packet roundtrip CA to Netherlands: 150ms
- Memory is fast but the disk is slow.
- Avoid disk seeks if possible.
- Simple compression algorithms are fast.
- Compress data before sending it over the internet if possible.
- Data centers are usually in different regions, and it takes time to send data between them.

## Availability numbers
- p50, p90, p99, 
- SLA

## Example
### Assumptions
- 300 M MAU
- 50% DAU
- post 2 tweets per day
- 10% contain media
- data is stored for 5 years

### Estimations
- DAU = 150 M
- tweet QPS = 3.5 K
- peek QPS = 7 K
- tweet size
  - tweet_id: 64 bytes
  - test 140 bytes
  - media 1MB
- media storage 30 TB per day
- 55 PB for 5 years

## Tips
- Rounding and Approximation
- Write down your assumptions
- Label your units, like KB, MB.
- Commonly asked back-of-the-envelope estimations: read/write QPS, peek PQS, storage, cache, network bandwidth
