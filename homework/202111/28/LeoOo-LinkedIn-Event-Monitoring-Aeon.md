# [Samza Aeon: Latency Insights for Asynchronous One-Way Flows](https://engineering.linkedin.com/blog/2018/04/samza-aeon--latency-insights-for-asynchronous-one-way-flows)

## Why
Monitor the latency of request processing is usually an easy task. This is because there's two clear events that needs to be captured in order to calculate the latency:
- the time that the request reaches the server
- the time the last byte of response leave the server.

Being able to monitor the latency of processing each request is key to understand the status of the system and finding potential performance bottleneck.

However, the monitoring task is not so trival if an event triggers a flow of events, especially such events flow from one system to another system. E.g. user send a request, then a push notification is sent.

## How
The solution proposed in this task is based on Apache Kafka, one reason being that Kafka is already being used for similar pupose at LinkedIn.

The idea is that at each 'checkpoint' of request processing, i.e. certain stage of the processing flow that we are interested in, a event (msg) is generated with a unique id with time stamp. As the same request went through the pipeline, more events with the same id are generated and are pushed to Kafka. Lastly, when the event reaches the exit, a final event is generated.
