# 0000-system-youtube
- https://blog.csdn.net/xiaozaq/article/details/109378302

## Functional Requirements
- upload/view/share/search videos
- count views/likes/dislikes
- add/view comments
- client: web, app, TV
- streaming video

## Non-functional Requirements
- high reliability/availablity/scalability
- low latency

## Not Required
- recommendation, popular / hot videos
- channel, subscription, newsfeed
- watch later, favorites, playlist
- authentication, autorization, encryption

## Capacity Estimation
- 1.5B users, 800M DAU, view 5 videos / day
- QPS: 800M * 5 / 24 / 3600 = 4 * 10^9 / 10^5 = 40K videos / sec
- upload/view ratio = 1:200,  40K / 200 = 200 videos / sec
- video average: 720p * 5 minutes = 1280 * 720 * 3 * 30 frames / s * 60 * 5 = 30 * 10^6 * 10^3 = 30GB?
- video average: 720p * 5 minutes * 5 types = 2.5 GB
- *** video average: 720p * 5 minutes * 5 = 10M * 5 * 5 = 250 MB
- Write: 200 video * 250 MB = 50 GB / sec
- Upload: 50 GB / 5 types = 10 GBps
- Download: 10 GBps * 200 = 2TBps

## System API
- String upload(String token, String title, String desc, String[] tags, bytes[] contents)
	- return { status: 202, url: "xxx" }
- String search(token, query, location, count, pageToken)
	- return [ { id: 1, title: "abc", url: "xxx", thumbnail, date, views } ]
- Stream play(token, videoId, offset, codec, resolution)

## Database Design
- videos(id, title, desc, creatorId, size, thumbnail, url, likes, dislikes, views, createdAt)
- comments(id, videoId, userId, content, timestamp)
- users(id, type, lastName, firstName, email, addr, birthDate, details)
