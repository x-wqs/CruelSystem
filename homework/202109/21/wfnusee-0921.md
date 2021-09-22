# Chapter2: BACK-OF-THE-ENVELOPE ESTIMATION

## Power of two
data is stored by bits...

## Latency numbers every programmer should know
L1 cache reference: 1 ns
Branch mispredict: 3 ns
L2 cache reference: 4 ns
Mutex lock/unlock: 17 ns
Main memory reference: 100 ns

Send 2,000 bytes over commodity network: 44 ns
SSD random read: 16,000 ns
Read 1,000,000 bytes from memory: 3,000 ns
Round trip in the same data center: 500,000 ns

Disk seek: 2 ms
Read 1,000,000 bytes from dist: 825,000 ns

## Availability numbers
SLA: 99.999...

## Estimate Twitter QPS and storage requirements

Tweets QPS = DAU * (2 tweets / 24 hour / 3600 seconds)
           = (300 million * 50%) * (2 tweets / 24 hour / 3600 seconds)
           = 7000

tweet_id: 64 bytes
text: 140 bytes
media: 1 MB

5-year media storage = 150 million * 2 * 10% * 1 MB * 365 * 5 = ~ 55 PB

## Tips

- Rounding & Approximation
- Write down your assumption
- Label the units
- QPS, peak QPS, storage, cache, number of servers
