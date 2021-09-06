# Video streaming system - Youtube, Tiktok

## Functional Requirements 
- upload video
- watch video
- generate thumbnail
- share video
- like video
- comments
- log in, etc.

## Non Functional Requirements
- High availability
- High scalibility 
- Low latency 
- real time reommendation 

## Estimation
- MAU: 2B, DAU : 150M
- Upload rate: 100 video/s, assume each video is 5 mins, and size is 50 MB/min
  - storage: 100 * 5 * 50 = 25000 MB/s = 24.5 G/s
  - bandwidth: 100 * 5 * 60 * 166KB/s = 5 G/s, assume 166KB/s

## Service
- user service
- upload service
- encode service
- thumb service
- video service

## storage
- video table: video_id, hash_code, resolution, size, duration, metadata
- chunk table: chunk_id, video_id, start_time, end_time, path, resolution
- thumbnail table: thumb_id, video_id, path, size, type, memoent
- video table to chunk, thumbnail table is one to many
- user table: user_id, user_name, email, gender, etc.
- user video talbe: uer_id, video_id, create_time, metadata
- video table and user table is many to many.

## work flow
- client send upload requet to web server
- web server return encode server address
- client send chunks to encode server
- encode server save metadata to db
- encode server save chunks to cloud storage server
- web server insert record to user video table 

## Q&A
- client repeat upload video -> upload both, then show success frontend and deduplicate later
- how to store chunks and thumbnails -> seaweedFS or FastFS
  - the difference of small and big size files, chunks and thumbnails are small size.
  - thumbnails get higher read rate
- how to do low latency -> preload while watching
- cover thumbnail -> fetch from file server
- progress thumnail -> preload to local cache when user click on video

## Scale
- CDN, content delivery network
- sharding
  - user video table shard by user id, easy to search all videos for given user
  - video table, chunk table shard by video id, easy to fetch all chunks for given video
