- WebServer: maintain user connection, auth and limit rate
- PostService: Persist Post data into DB and cache
- Newsfeed Service: store posts/feeds
- NewsFeed Cache: store news feed IDs
- Retrieve Newsfeed
	- request newsfeed, like /rest/feed?before=1234&count=20
	- LB distribute traffic to WebServer
	- WebServer call Newsfeed Service
	- Newsfeed Service gets post ID list from Newsfeed Cache
	- Newsfeed Service gets user info, post content from user/post cache/DB
	- The full hydrated newsfeed is returned in JSON format.

## Detail Design

- Fanout Service
- write, push, pre-compute, real time, fast, Cons: too many friend, hotkey/celebrity problem, inactive user
- read, pull, on demand, suitable for inactive user and celebrity, Cons: slow
- Hybrid, push for majority
- FanoutWorker read <PostID, FriendID>

## Cache Architecture
- NewsFeed:
- Content: HostCache, Normal
- Social: Follower, Following
- Action: Like, Reply, Others
- Counter: Like, Reply

## Terminology
- reverse chronological
- build behind the scence
- avoid pitfall
- crucial
- overload
- slim
- the full hydrated news feed
- compilation
- scrollable
