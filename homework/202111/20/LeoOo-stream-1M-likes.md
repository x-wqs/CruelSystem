
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


