# [LinkedIn's Real-Time Presence Platform](https://engineering.linkedin.com/blog/2018/01/now-you-see-me--now-you-dont--linkedins-real-time-presence-platf)

This tech blog walk through how LinkedIn use the [play framework](https://www.playframework.com/) and the [Akka actor model](https://akka.io/) to track and distribute the presence of millions of users at any moment.

## Problem 1: How do we know a user is online?
We want to:
- Identify user's online status and keep track of the last active timestamp.
- Distribute changes of user's presence status to the user's connections.

The system leverages a real-time delivery platform built for instant messaging that is a publish/subscribe system that allows data to be streamed from the server to mobile or web clients over a persistent connection as events happen.

When Alice login via mobile or web, Alice build a persistent connection with the server and the server subscribe Alice to the topic of Bob's presence (Alice is connected with Bob). When Bob logged in, an event will be published to Alice.

> This is a working solution. But creating a topic for each user might be too expensive, millions of topic might be expensive to manage.

## Problem 2: Unstable/Jittery connection
Unstable connection would be due to lossy networks, and it can cause user's status change haphazardly (i.e. randomly bouncing online and offline), which could:
- leads to a bad experience to the viewer
- cause a lot of wasteful traffic and overhead throughout the entire system.

Solution: periodic heartbeats.
- The duration between heartbeats smoothen the fluctuations. As long as a heart beat is received in every d seconds, the user status will stay online.

### Processing the heart beat.
- To process the heartbeat with a userID, the system first check if there's an entry for the userID in a distributed KV store.
  - If no such entry, we conclude the user just went online, therefore an event is published to all connection of the user's.
  - An new entry is created with the expected expiration time and userID, also the last_online_timestamp is updated.
  - An existing entry is updated at every heartbeat.

## Problem 3: Detect absence of heartbeat
How to detect the absence of heartbeat? High-level idea is to create a 'delayed trigger' every time that we process a heartbeat. The trigger get reset when the next heartbeat come in, or it trigger the system to check the heartbeat of the user and publish an offline event. Any node can check the distributed KV store to check a user's presence.

### How to implement delayed trigger?
* With Akka actor (akka.io)

Millions of user go online and offline at any moment, the solution needs to be lightweight.

Akka is a toolkit for building highly concurrent, distributed, and resilient message-driven applications, and it works well with the Play Framework that we use at LinkedIn. Actors are objects which encapsulate state and the behavior defining what they should do when they receive certain messages. Each Actor has a mailbox and they communicate exclusively by exchanging messages. An Actor is assigned a lightweight thread when available, which reads the message from the mailbox and alters the state of the Actor based on the behavior defined for that message.

ThreadPool (e.g. 64 threads for 32 cores) => 2K Actors => Akka Scheduler => Publish to 2K Users.

> I'm wondering how the actor framework is implemented, to be so lightweight but powerful, how can it now holding a thread for the scheduler to work, what if there're so many actors in the system that some actors starve?
> Seems it is kind of like a command processor, the messages are like commands, upon receiving the command, the actor execute the command and alter its state.

### Pttting it together

Alice Online => HeartBeat => Akka actor => Publish online event to her connections

Alice missed a heatbeat in the next `d` seconds => Akka actor delayed scheduler triggered => Publish offline event to her connection

The Real-time Platform sends the change in Bob’s status to Alice and Alice can now see Bob as online. This process is really fast and it takes less than 200ms at p99.

This platform is built for massive scale. Each node in the Presence Platform can handle approximately 1.8K QPS of incoming requests (process heartbeat requests, GET API requests, etc.)

> 1.8K QPS is probably due to the system has to maintain long poll for each user which become the bottleneck of the system. 
