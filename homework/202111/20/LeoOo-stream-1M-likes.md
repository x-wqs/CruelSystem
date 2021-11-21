
# Stream 1 million likes.

LinkedIn

[From InfoQ](https://www.youtube.com/watch?v=yqc3PPmHvrA)

> Principle: start small and add simple layers to your architecture.

## Challenge: interactions of live videos. How are the likes/comments get distributed to users?

### Challenge 1: the delivery pipe.

viewer 1 -> likes -> likes backend (publisher) -> real-time delivery -> persistent connection -> viewer 2

Persistent connection: HTTP Long Poll with Server Sent Events

The eventsource interface can handle HTTP requests from user and send response back in a JSON format:
```
data: {'like' object}
data: {'comment' object}
```

### Challenge 2: connection management
How to manage multiple connections from different users?

`Akka` is a toolkit for building concurrent, distributed message-driven applications.

Akka actor is an object with 'state' and 'behavior'
Each actor has a message queue, which holds events to be published. One actor can be used to manage a lot of connections, each conenction is assigned a conenction id, and actor publish events to each connection.

Actors are managed by a supervisor actor (i.e. a hierarchical structure), the supervisor send likes to the actor and actors publish the events to the connections (broadcasting).

### Challenge 3: Multiple live videos (whos watching who?)
Subscription, whenever a client is watching a video, it subscript to the live video, in-memory subscription table (in-memory).
* (connection_id, video_id)
* Supervisor then lookup the table to find which client is watching which video.

### Challeng 4: More than a single machine can handle
10K concurrent viewers.

Add more machine (front-end nodes) and a real-time dispatcher. Dispatches between frontend nodes.
Dispatcher maintain the map between video and front-nodes.
* (node_id, video_id)
Dispatcher dispatches the event to the nodes that are responsible to certain videos.

The dispatcher is the bottleneck of the system.

Add more dispatcher node.

Put the (video, node) to a standalone key-value store available to all dispatcher nodes.

Each frontend-node can subscribe to any dispatcher node.

?? What about the part that process the likes, seems out of the scope of the talk?

### Challenge 5: Distribute 1M likes/s.
