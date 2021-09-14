# SCALE FROM ZERO TO MILLIONS OF USERS

## Cache
- Caches take advantage of the locality of reference principle 
   - recently requested data is likely to be requested again
- advantages:
  - better system performance
  - reduce database workloads
  - scale the cache tier independently
- cache Invalidation
  - write through cache: data is written into the cache and the corresponding database simultaneously
  - Write around cache: data is written directly to permanent storage, bypassing the cache
  - Write back cache: data is written to cache alone, and completion is immediately confirmed to the client. The write to the permanent storage is done after specified intervals or under certain conditions
- when to use: data is read frequently but modified infrequently
- SPOF 
- Eviction Policy
  - First In First Out (FIFO)
  - Last In First Out (LIFO)
  - Least Recently Used (LRU)
  - Most Recently Used (MRU)
  - Least Frequently Used (LFU)
  - Random Replacement (RR)
- Hot spot:
  - Can identify by log or monitor system, like high CPU
  - Maybe a few keys cause the issue, then need to spread them evenly
  - In extreme cases, a single hot cache key create the hot spot, then a)Remap a separate set of cache nodes, b) Add a second layer of smaller caches in front of main nodes to act as a buffer.  This approach gives you more flexibility, but introduces additional latency into your caching tier
  - add read replica for Redis
- 缓存穿透 cache penetration, 频繁访问不存在的key，
  - 不存在的key value设成默认值或者空值，但是TTL不能太长
  - 使用bloom filter 把所有存在的key 放进去，bloom filter 可以判断书籍一定不存在或者可能存在
- 缓存击穿 cache breakdown， 大量同时访问某一个过期key
  - 通过加锁，只让一个线程去访问db
  - 永不过期，异步定期更新key 值
- 缓存雪崩 cache avalanche，大量key 同时过期
  - Some how randomize TTL
  - Redis replica, not work for memcached 
  - 服务降级或者熔断

## CDN
- kind of cache that comes into play for sites serving large amounts of static media
- In a typical CDN setup, a request will first ask the CDN for a piece of static media; the CDN will serve that content if it has it locally available. If it isn’t available, the CDN will query the back-end servers for the file, cache it locally, and serve it to the requesting user
- If the system we are building is not large enough to have its own CDN, we can ease a future transition by serving the static media off a separate subdomain (e.g., static.yourservice.com) using a lightweight HTTP server like Nginx, and cut-over the DNS from your servers to a CDN later
- consider use of CDN
  - cost
  - Setting an appropriate cache expiry

## Web tier
### stateless web tier
- scaling the web tier horizontally
- move state(for instance user session data) out of the web tier
- store session data in the persistent storage such as relational database or NoSQL
- Each web server in the cluster can access state data from databases

### Stateful architecture
- stateful server remembers client data (state) from one request to the next
- stateless server keeps no state information
- issue is same client's request must be routed to the same server
  - can be done with sticky sessions in most load balancers
  - adding or removing servers is difficult with this approach

### Stateless architecture
- requests from users can be sent to any web servers
- servers fetch state data from a shared data store
