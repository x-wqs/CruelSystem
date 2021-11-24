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

TODO

### Putting it all together

TODO