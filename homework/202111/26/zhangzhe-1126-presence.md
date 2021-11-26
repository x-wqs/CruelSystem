# LinkedIn's Real-time Presence Platform



- Play Framework
- Akka Actor Model



### Problem 1: How do we know when a member is connected to LinkedIn

Goal: 

- Distribute changes in a member's presence status to a potentially huge number of the member's connections

Solution: 

- Real-time delivery platform: publish/subscribe system.

  1. When Alice connects to the Platform, she **subscribe the topic** about Bob's presence status

  2. When Bob get online, Platform **publish an online event** on the topic about Bob's presence status.
  3. The Platform **sends the event** published on the topic **to its subscribers** (Alice).



### Problem 2: Jittery connections

Problem Description: 

- Unstable connection of mobile users leads to jittery presence status. 
- It would result in huge traffic to distribute these changes.

Solution:

- Use periodic heartbeats.
  1. When user opens the app, it results in a persist connection with the Real-time Platform
  2. Platform sends periodic heartbeats with the member with a duration $d$ seconds.
  3. If the user is in a bad network condition, drops connection and reconnects within $d$ seconds, Platform still deem him as online.
- Details of periodic heartbeats
  - Include: `member ID`, `LastHeartbeatTimestamp`
  - Platform checks whether there is an unexpired entry for the `member ID` in the distributed K/V store for presence.
    - If there is an unexpired entry, then updates the `LastHeartbeatTimestamp` to the latest value for this user.
    - If there is no entry or previous one has expired, 1) Conlude this user became online; 2) Publish an online event and distribute this event; 3) Add an entry to the store with an expiry duration which is slightly larger than $d$ seconds $(d+\epsilon)$ to keep the record alive at least the latest value.



### Problem 3: Detecting the absence of a heartbeat

Givne that Heartbeats expire in $(d+\epsilon)$ seconds:

Idea: Use a timer which fires in $(d+2\epsilon)$ s to check whether the heartbeat stored has expired.

- During the `process heartbeat` step, create `delayed trigger` if it doesn't yet exist; If it already exists, reset it to fire in another $d+2\epsilon$ s.
- When the delayed trigger for an online member fires, we check whether the member's heartbeat has expired in K/V store: If it has, publish an offline event.



Multiple servers case:

- to prevent duplication of the delayed triggers for a member on multiple nodes, do a best-effort sticky routing of a given member’s heartbeats to a given node.
  - Hash the member ID in the heartbeat request to route it to the same node. (Load balancer strategy used in d2 load balancer of LinkedIn)



Delayed triggers using Akka Actors

- Actor: Actors are objects which encapsulate state and the behavior defining what they should do when they receive certain messages. Each Actor has a mailbox and they communicate exclusively by exchanging messages. An Actor is assigned a lightweight thread when available, which reads the message from the mailbox and alters the state of the Actor based on the behavior defined for that message.

- Create one Actor per online member to act as the delayed trigger and treat a heartbeat like a message in its mailbox.
  1. When a `Heartbeat` message is received for a member, we create an Actor for the member if it doesn’t already exist and start an Akka Scheduler within the Actor to send a `Publish Offline` message to itself after d + 2ε seconds.
  2. When the scheduler fires, the Actor handles the `Publish Offline` message by publishing an offline event if the member’s entry has expired in the K/V store, as described before.



### Putting it all together

The Real-time Platform sends the change in Bob’s status to Alice and Alice can now see Bob as online. This process is really fast and it takes **less than 200ms at p99**.

Each node in the Presence Platform can handle approximately **1.8K QPS** of incoming requests (process heartbeat requests, GET API requests, etc.) and the platform itself is horizontally scalable.

