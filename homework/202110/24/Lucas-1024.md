# Design TinyURL
- SDE Skills, https://www.youtube.com/watch?v=IfzTeMtgk7s

## Functional Requirements
- user gives a URL and get a short URL
- short URL will redirect to the original URL
- URL will not expire
- no authentication but reCAPTCHA, or for script use
- CAPTCHA: Completely Automated Public Turing test to tell Computers and Humans Apart
1. URL shortening
2. URL redirecting
3. High Availability, sacalability and fault tolerance

## Non-functional Requirements
- Scalability
- Availability
- Performance / Low Latency

## Capacity Estimation
- Traffic: 100M URLs per day
- Write: 100M / 24 / 3600 ~= 1K writes / second
- Read: read/write ration = 10, 10 QPS
- Store: 100M * 365 days * 5 years ~= 100M * 2K = 200B records
- Storage: 200B * 100 bytes = 200TB

## API Design
- POST 	/api/v1/url/shorten?longURL=xxx
- GET	/shortURL, e.g. https://tinyurl.com/aBcd5eF
- HTTP 301 Redirect permanently, and the browser cache the response and will not send traffic to TinyURL
- can reduce the load
- HTTP 302 Redirect temporarily, and the browser will keep sending request to TinyURL
- can analyse URL usage

## Sequence Diagram
Client                       TinyURL
 |								|
 |------- visit short URL ----->|
 |<- HTTP 301, location:longURL-|
 |------- visit long URL ------>|
 
## Database
- urls(id, shortURL, longURL)
- hash length: (26*2+2)^7 > 200B

## Hash Function - Hash + Collision
- First 7 characters of CRC232/MD5/SHA-1
- Keep appending pre-defined string and hash until no collision
- Expensive to query database for collision
- bloom filter, space efficient

## Hash Function - Base 62 Connersion
- 11157 = 2*62^2 + 55^62 + 59 = "2TX"
- Depends on a distributed unique ID generator
- No need to handle collision
- Easy to figure out the next available short URL
1. Input long URL
2. Check if the long URL in the database
3. Return short URL if existing
4. Generator unique IDif not existing
5. Convert ID to short URL with Base62 conversion
6. Create mapping in database

## Cache
- more read than write
- put mapping in cache with TTL

## Enhancement
- rate limiter
- scalability with application server
- Database sharding and replication
- Analytics for when, where, how often

## 提示
- 1 day = 24 * 3600 s ~= 25 * 4000 d = 100K seconds
- 5 years = 5 * 365 d ~= 5 * 400 d = 2K days
- 10^3	6	9	12	15
- KB	MB	GB	TB	PB
- K		M 	B

## 参考
- REST API
- 布隆过滤器

## 术语
tackle 应付
intentionally
well-crafted
clarificatioin
click the alias
traffic volume
envelope estimation
walk through
assumption
get buy-in
permanently
elinimate
probalilistic

