### don'ts
- jump to a solution without clarification
- get into too much detail on a single component in the beginning. Give high-level design
- ask for hints if stuck
- think in silence

## Time allocation
- Understand the problem & establish design scope 3 - 10 minutes
- Propose desgin 10 - 15 mins
- Design deep dive 10 - 25 mins
- wrap 3 - 5 mins

# Design a rate limiter

- In network system, rate limiter controls the rate of traffic sent by a client or a service
- HTTP - a rate limiter limits the # of client requests allowed to be sent over a specified period. If API request count exceeds the threshold, all the excess calls are blocked
- Benefits of API rate limiter
  - Prevent resource starvation caused by Denial of Service(DoS) attack
  - reduce cost
  - prevent servers from being overloaded

## Step 1 - understand the problem and establish design scope
Questions to ask
 - What type of rate limiter? Client-side or server-side
 - Throttle based on IP, user ID or?
 - Built for a startup or a large company with a large user base
 - Work in a distrubted environment?
 - A separate service or implemented in application code
Requirements
- Accurately limit excessive requests
- low latency - should not slow down HTTP response time
- Use little memory
- Distributed rate limiting - across multiple servers or processes
- Exception handling
- high fault tolerance
