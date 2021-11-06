# 0000-system-design-news-feeds
- include status update, photos, videos, links, activity and likes
- from people, pages, groups you followed
- Facebook news feeds, Instagram feed, Twitter timeline etc

## Functional Requirements
- a user may has many friends or follow a large number of pages and groups
- newsfeed will be generated based on posts from people, pages, groups that a user follows
- feeds may contain text, image and videos
- user can refresh and view newsfeeds
- newfeeds are sorted by reverse chrononlogical order
- a user can receive notification when there's news feeds update
- both mobile app and web pages

## Non-funtional Requirements
- it's better to generate newsfeed in real-time
- maxmium latency seen by end user would be 2 secondso
- a post should not take more than 5 second before assumed

## Capacity Estimation
- 300M daily active user, see newsfeed 5 times per day
- QPS: 1.5B / 24 / 3600 ~= 1.5 * 10^9 / 10^5 = 15K read
- Memory: 300M * 500 * 1K = 150TB = 100G * 1500 machines

## System API
- int post(token, content)
- List<Long> getNewsfeed(token, count, beforePostId, afterPostId)

## Database Design
- users(id, lastName, firstName, email, birth, createAt, lastLogin)
- access_tokens(id, token, userId)
- entities(id, name, type, desc)
- follows(id, friendOrEntityId, type)
- posts(id, entityType, entityId, lat, lon, content, likes)
- medias(id, type, link)
- post_medias(id, postId, mediaId)
- TODO-PG 表继承

## High Level Design
- Publish Feed / Fanout
	- write to cache and DB, populated to friends's new feed
	- by analogy 类推
	- Push / fanout on write, the right pull cadence 节奏
	- Pull / fanout on load,
- Build News Feed
	- Pull
		- retrieve IDs of all users/groups/pages a user follows
		- retrieve the latest, most popular and relavant posts
		- rank these posts with based on the relavance, likes, commets, shares, update time
		- sort the feed in the cache and return top 20
		- return next 20 if scroll down
		- periodically rank for the new posts and notify
		- high backlogs 高积压
	- Offline
		- dedicated server to generate continuously
		- not compiled on load 符合负载，regular basis 定期
		- LRU, LinkedHashMap
		- keep 20 posts * 10 page in memory
		- only generate for active user or with login pattern
