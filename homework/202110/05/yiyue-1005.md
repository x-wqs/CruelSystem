## Step 2 - propose high-level design and get buy-in

### Place rate limiter
- Client-side - unreliable. Can easily forged by malicious actors. No control over client implmentation
- Server-side
- Middleware
  - client sends 3 requests to the server within one second
  - first two routed to API server
  - throttle the third
  - cloud microservices have API gateway to contain rate limiting

### algorithms

#### Token bucket
- token bucket - a container with pre-defined capacity. tokens are put in the bucket at preset rates periodically. Once bucket full, no more tokens added
- each request consumes one token
  - if enough tokens, take one for each request, request goes through
  - if not, drop request
- params
  - bucket size
  - refill rate
- # of buckets
  - different buckets for different API endpoints
  - if throttle based on IP address, each address - one bucket
- pros
  - easy
  - memory efficient
  - allow for burst of traffic for short period
- cons
  - challenging to tune parameters


#### leaking bucket
requests -> bucket full? ->(no)-> queue -> (processed at fixed rate) -> request go through
						 ->(yes, drop requests)

- params
  - bucket size
  - outflow rate - # requests processed at fixed rate
- pros
  - memory efficient
  - fixed rate - suitable when stable outflow rate needed
- cons
  - burst of traffic fills up the queue
  - hard to tune parameters

#### fixed window counter
- divide timeline into fix-sized time windows, assign a counter for each window
- request increments the counter by one
- once counter reaches pre-defined threshold, drop new requests unitl new time window starts
- pros
  - memory efficient
  - easy
  - resetting vailable quota at end of unit time window fits certain use cases
- cons
  - spikes in traffic at edges of window could cause more requests than allowed to go through