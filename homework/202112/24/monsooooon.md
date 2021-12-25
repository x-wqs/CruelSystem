ref: http://highscalability.com/blog/2013/7/8/the-architecture-twitter-uses-to-deal-with-150m-active-users.html


twitter overview (in 2013)
# current scale
- 150M global active users (DAU MAU?)
- 300 QPS to generate timelines (Peak? Avg?)
- [Firehose](https://www.pubnub.com/blog/what-is-a-data-firehose-api/) 22MB/s
- 400 million tweets/day
- 5 sec to broadcast a celebrity's tweet to her followers

# important pts
- use API, not just a website
- 300K QPS for reading, 600 QPS for writing
- outliers that have huge follower lists are common
  - more and more
  - these are VIP
- implicit social contract calculation via following relationship & clicks
- need to have monitoring & debugging facilities, like metrics/logs/traces

# the challenge
- how to handle heavy-read reqs pattern? use write based fanout approach
- do a lot of processing when tweets arrive to figure out where tweets should go. This makes read time access fast and easy. Don’t do any computation on reads.

# push or pull mode
- mode of timeline: 
  - A __user timeline__ is all the tweets a particular user has sent. 
  - A __home timeline__ is a temporal merge of all the user timelines of the people are you are following. 
- pull based timeline generation
  - you ask for the timeline from REST API call, you get it
- push based
  - when writting, the tweet goes to every followers' cache
- timeline: array of tweetIDs, use multi-get from other storages (T-bird) to build real timeline contents
- for active user(logged-in within 30 days), their timelines are stored in fast cache (like Redis cluster) with 800 entries long
- read filtering happens on the edge (just before responding to API req)

