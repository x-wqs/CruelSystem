# Reliable, scalable and maintainable applications

- Data intensive application vs compute-intensive
  - The amount of data, the complexity of data and the speed of change of the data.

- Reliability
  - The system should continue to work correctly, even when things go wrong.
    - Fault: things that could go wrong.
    - Being fault-tolerant and resilient.
    - But it is not about tolerating all kinds of faults.
    - Fault vs failure
      - Fault: one component of the system deviate from its spec.
      - Failure: the system as a whole stop to serve the user.
    - Counterintuitively, deliberately incuding/injecting faults is a good approach to test the fault-tolerance of the system.
    - In somce case, prevention is better than cure (because cure does not exist), e.g. security related.
    - But in this book, we focus on faults that can be cured.
  - Hardware error are random and independent.
    - Have backup and buy from different provider/batch.
  - Software error can be subtle yet catastrophic.
    - No silver bullet.
    - Thorough testing, process isolation, monitoring and analysis, self-diagnose.
  - Human errors are unavoidable.
    - Good abstraction.
    - Decouple the error prone component with critical component.
    - Provide sandbox to test.
    - Unit test and intergration test.
    - Be able to recover from human errors. Gradually roll out and rollback.
    -
    -
- Scalability
  - Being able to deal with growth with a reasonable cost.
- Maintainability
  - How friendly ths system is to work with different people.


