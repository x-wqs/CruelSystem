# Estimation

## Power of two
- A byte = 8 bits
- An ASCII charactor = one byte

## Latency numbers
- L1 cache reference: 1ns
- Branch mispredict: 3ns
- L2 cach reference: 4ns
- Mutex lock/unlock: 17ns
- Send 2000 bytes over commodity network: 44ns
- Memory is fast. Disk is slow
- Avoid disk seeks
- compression algorithms are fast. Compress data if possible

# Availability numbers
- High availability - ability of a system to be continously operational for long period
- measured as percentage
- SLA (service level agreement) - agreement between service provider and customer - level of uptime service will deliver

## estimate Twitter QPS(Query per second) 
- DAU (daily active users) = monthly active users * % of users use daily
- Tweets QPS = DAU * (avg tweets per day) / 24 hour / 3600 seconds
- Peak QPS = 2 * QPS

# Tips
- rounding and approximation - no need to perform complicated math operations
- write down assumptions
- label your units - e.g. kb or mb?
- Commonly asked estimations - QPS, peak QPS, storage, cache, # of servers

# Framework

## 1 - understand the problem and establish design scope
- Specific features to build
- # of users the product has
- Scale up speed. Anticipated scales in 3 months, 6 months, a year
- technology stack. Existing services to leverage

Write down assumptions if any

E.g. news feed system
- Mobile / web app?
- Most important features?
- How is news feed sorted?
- # friends a user can have
- traffic volumn
- feed can contain images videos or text?

