#### sliding window log algorithm
- keeps track of request timestamps. Timestamp data kept in cache
- new request comes, remove all the outdated timestamps - older than the start of the current window
- add timestamp of the new request to log
- log size <= allowed count, accept, else, reject

- pros
  - accurate
- cons
  - consumes a lot of memory. rejected request's timestamp still stored

#### sliding window counter algorithm

# of requests in the rolling window = requests in current window + requests in previous window * overlap % of rolling window and previous window

- pros
  - smooths out spikes in traffic
  - memory efficient
- cons
  - only works for not-so-strict back window

### high-level architecture

- A counter track # requests sent from same user / IP / etc
- counter > limit, request disallowed

Where to store counters
- database - bad, slowness of disk access
- in-memory cache - fast and supports time-based expeiration strategy

## Step 3 - design deep dive

### rate limiting rules

- rules written in configuration files and saved on disk

### exceeding rate limit

- a request is rate limited
- APIs return HTTP response 429 (too many requests)
- depending on the use cases, may enqueue some rate-limited requests

### rate limiter headers

HTTP headers
`X-Ratelimit-Remaining` : remaining # of allowed requests within window
`X-Ratelimit-Limit` : # calls can make per window
`X-Ratelimit-Retry-After`: # of seconds to wait until can make request again
