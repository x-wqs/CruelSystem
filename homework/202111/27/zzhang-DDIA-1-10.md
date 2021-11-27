# DDIA pp1-10: Reliability

Focus:

- Reliability

- Scalability

- Maintainability



#### Hardware Faults

**Example**: hard disk crash, RAM become faulty, power grid has a blackout

**Characteristic**: very frequent. Hard disk's mean time to failure is about 10 - 50 years.

**Response**:

1. Add redundancy: appropriate for small cluster.
2. Software fault-tolerance techneique: A single-server system requires planned downtime. However, a system that can tolerate machien failure can be patched one node at a time, without downtime of the entrie system.



#### Software Errors

**Example**: software bug (leap second), runaway process that uses up some shared resources, the service that the system depends on slows down, casading failures (small fault in one component triggers a fault in another components).

**Characteristic**: correlated, hard to find.

**Response**:

- Carefully thinking about assumptions (MQ's incoming message number, outgoing message number)
- testing, process isolation, allow processes to crash and restart
- Measure, monitor and analyze the system behavior.



#### Human Errors

**Example**: Configuration erros caused by operators

**Characteristic**: more common than hardware faults

**Response**:

- Well-designed abstractions, APIs, and admin interface.
- sandbox
- Thorough test
- Fast roll back, roll out new code gradually, provide tools to recompute data
- Detailed and clear monitoring
- management





