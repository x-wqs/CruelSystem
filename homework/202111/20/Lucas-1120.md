# 0000-system-streaming-million-likes
- https://www.youtube.com/watch?v=yqc3PPmHvrA
- https://www.infoq.com/presentations/linkedin-play-akka-distributed-systems/

## Functional Requirement
- click like
- broadcast likes

## Non-functional Requirement
- high reliablity/availability/sacalibity
- low latency

## Capacity Estimatioin
- Upload: 100 likes / s
- Download: 

## System API

## Database Design
- Subscription KVStore
- <VideoID, [Frontend1, Frontend2]>

## High-level Design

- Client1 ---> Frontend1 ---> DispatcherN -> SubscriptionStore
-                                 |
- Client2 <--- AkkaActor <--- AkkaSupervisorActor

- Client: WebSocket, Bi-direction

## Detail Design
- Persistent Connection between Client and Frontend
- Connection Management: Akka Actor
- Multiple Video: Subscription with KVStore
- 10K Concurrent Views: more frontends
- Create 100 likes/s: multiple dispatcher
- Distribution of 1M likes/s: 

## Question
- 这个看上去好复杂，Dispatcher和Frontend之间为啥不用消息队列啊，每个Video一个Topic？谢谢了

## Teminology

