
## How
The solution proposed in this task is based on Apache Kafka, one reason being that Kafka is already being used for similar pupose at LinkedIn.

The idea is that at each 'checkpoint' of request processing, i.e. certain stage of the processing flow that we are interested in, a event (msg) is generated with a unique id with time stamp. As the same request went through the pipeline, more events with the same id are generated and are pushed to Kafka. Lastly, when the event reaches the exit, a final event is generated.

### Process events
Once the events are generated, we need to process them, to make the system scalable, Samza, a stream processing system is used. First, a partitioner partitions the events by the id, so that all events for the same requests are handled by the same machine. Second, a joiner calculate the pairing events and send the result to the next step of the pipeline.

> Is Samza a map-reduce framework? what is a partitioner and joiner exaclty?

An in-memory is used cache to store the events, whenever an end-event reaches the system, the system find the corresponding start-event (if any) and calculate the latency, if there isn't such a matching event, then it will wait longer.

After the calculation, the latency is sent to the next step of the monitoring pipeline, where more meaningful analysis is performed, e.g. plotting, case studying, diagnosing.

### How Aeon is used
- Debugging a notification pushing bug. Used to find the bottleneck of the system.
- In the realtime platform, to monitor the system performance.
