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

### Delayed trigger with Akka actor
Millions of user go online and offline at any moment, the solution needs to be lightweight.
Akka is a toolkit for building highly concurrent, distributed, and resilient message-driven applications, and it works well with the Play Framework that we use at LinkedIn. Actors are objects which encapsulate state and the behavior defining what they should do when they receive certain messages. Each Actor has a mailbox and they communicate exclusively by exchanging messages. An Actor is assigned a lightweight thread when available, which reads the message from the mailbox and alters the state of the Actor based on the behavior defined for that message.

Since Actors are so lightweight, there can be millions of Actors in the system, each with their own state and behavior. A relatively small number of threads proportional to the number of cores on the system can serve the Actors to execute their behavior when they receive these messages. *As each Actor receives a message, a thread is assigned to execute its behavior for that message* and then, it is free to be assigned to the next Actor.

We can basically create one Actor per online member to act as the delayed trigger for that member and treat a heartbeat like a message in its mailbox.

When a Heartbeat message is received for a member, we create an Actor for the member if it doesn’t already exist and start an Akka Scheduler within the Actor to send a Publish Offline message to itself after d + 2ε seconds.

When the scheduler fires, the Actor handles the Publish Offline message by publishing an offline event if the member’s entry has expired in the K/V store, as described before.

> I'm wondering how the actor framework is implemented, to be so lightweight but powerful, how can it now holding a thread for the scheduler to work, what if there're so many actors in the system that some actors starve?
> Seems it is kind of like a command processor, the messages are like commands, upon receiving the command, the actor execute the command and alter its state.

### Putting it together
When Alice is interested in Bob’s presence status, she queries the Presence Platform through its GET API to get Bob’s current presence status (offline) and his last heartbeat timestamp. At the same time, she subscribes to changes in Bob’s presence status through the Real-time Platform.

When Bob comes online, the Presence Platform detects that and publishes an online event to Bob’s presence status topic on the Real-time Platform.

The Real-time Platform sends the change in Bob’s status to Alice and Alice can now see Bob as online. This process is really fast and it takes less than 200ms at p99.

This platform is built for massive scale. Each node in the Presence Platform can handle approximately 1.8K QPS of incoming requests (process heartbeat requests, GET API requests, etc.)

> 1.8K QPS is not that impressive.Even on the slow side. It would take 1000 nodes to handle 1M QPS?


