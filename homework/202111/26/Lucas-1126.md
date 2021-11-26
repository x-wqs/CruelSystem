### Offline Status
- PP will create or reset delayed trigger when receiving heartbeat
- If no heartbeat received in the interval, the trigger will be executed.
- PP will send an offline event to the topic on RP.

### Scalability
- a best-effort sticky routing
- d2 loader balancer

### Delayed Trigger
- 10K online/offline events / second
- Akka and Play Framework, a truly lightweight solution
- One Akka Actor per online user
- Each Akka actor has a mailbox and will act as the message received
- the message may heartbeat message
- Actor will create delayed trigger with Akka Scheduler
- Actor will publish an offline event to the topic when the scheduler fires

### Graceful Shutdown
- Actors are in memory
- increased delay for deployment and reboot

## Terminology
- jittery, 神经过敏的；战战兢兢的；紧张不安的
- lossy, 
- haphazardly
- constantly bounce
- a massive amount
- fluctuation, 波动，起伏
- dive into
- periodic heartbeat
- result in
- in turn
- drain, 劳累
