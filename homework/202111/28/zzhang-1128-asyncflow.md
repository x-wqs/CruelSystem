# Latency Insights for Asynchronous One-Way Flows

总结：之前monitor只能看单个request的delay，为了看整个work flow的delay，在task前后加上带相同key的start event和end event，算一下这两个event的时间差即可。



Asynchronous delay monitoring: 

**Issues:**

- A request is sent to a system and begins a chain of events moves through several other systems before completing.
- This request may return immediately and then continue processing asynchronously.
- The current monitors only show the delay between the sending time and the received time of immediate response instead of the whole flow.



**Main idea: **

- Emit tracking events to Kafka to power a latency monitoring solution



**Example:** 

Latency within LinkedIn of sending a **push notification** for a message to a member's phone.

- Start event [StartTrackingEvent]: when a request to create a message is received by the messaging platform.
- End event [EndTrackingEvent]: when our push notification platform publishes a push notification to Apple Push Notification Service (APNS) or Google Cloud Messaging (GCM).





Use Samaza to **join events**: connect two events like *StartTrackingEvent* and *EndTrackingEvent*

* Samza: Stream processing framework, consum events and perform computation on events.

* Use two jobs:

  * **Partitioner**: partitions the events to ensure that the events to track as a single push notification end up on the same process
  * **Joiner**: calculates the latency of those matched events.

  

### Partitoner

- Two related events *StartTrackingEvent* and *EndTrackingEvent* share the same key, e.g. pushId.
- Use this key to partition events with the same pushId to the same message queue, so that they ultimately end up being processed by the same Joiner.



### Joiner

- Every time it consumes a new EndTrackingEvent, it checks to see if there is a matching *StartTrackingEvent* entry in the cache. 
  - If it finds one, it calculates the latency by subtracting the *StartTrackingEvent* emission time from the *EndTrackingEvent* emission time.  
  - It then takes that latency and emits it to InGraphs for real-time monitoring
    - InGraph: a latency monitor

Improvement:

- Use a timed expiry to automatically remove cache entries after a few minutes



### Use case

- Push notification
- Real-time platform: typing indicators, seen receipts, and message delivery 

