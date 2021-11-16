# Alex Xu Ch 14 Design youtube

## Functional Requirement

* Ability to upload videos fast

* Ability to change video quality

* Support multiple clients including mobile apps, web browser, smartTV

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should able to generate newsfeed in real time for each user

* Fault tolerance, tailover handling

## Back of the Envolop Calculation

* 5M daily active users 

* each user watch 5 videos per day

* 10% users upload 1 video perday with average size of 300MB

* daily storage requires 150TB


## High Level Design

### Client

We need to have multiple client applications to support browsing from computer, mobile phone, and smartTV

### CDN

Videos needs to be dilivered through CDN, such that video could be streamd to clients from different locations

### API Servers

Service all requests besides video streaming, including feed generaaion, generating video upload URL, user signup etc.

## System Components

* user: clients watching video streams from different devices

* each user watch 5 videos per day

* Load balancer: A load balancer evenly distributes requests among API servers.

* API servers: All user requests go through API servers except video streaming.

* Metadata DB: storeing user and video metadata

* Metadata cache: performance improvement for metadata query

* Original blob storage: A blob storage system is used to store original videos. 

* Transcoding servers: Video transcoding is also called video encoding. It is the process of
converting a video format to other formats (MPEG, HLS, etc), which provide the best
video streams possible for different devices and bandwidth capabilities.

* Transcoded storage: It is a blob storage that stores transcoded video files.

* CDN

* Completion queue: A message queue that stores information about video transcoding
completion events.

* Completion handler: This consists of a list of workers that pull event data from the
completion queue and update metadata cache and database.

