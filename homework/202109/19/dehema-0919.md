# Chapter 3 A framwork for system design interviews

## step 1 understand the problem and establish design scope
- features detailed requirment
- DAU, MAU or users metric
- scalibility
- current tech stack

## step 2 propose high-level design and buy-in
- Sketched out a high-level blueprint for the overall design, and ask feedback.
- draw diagrams for the key components, like the service, database, message queue, CDN, queues and etc
- give the APIs of each components, and database schema.
- do back-of-the-enveloper calculations to check the scale.


## step 3 design deep dive 
- dive to the components, interviewer interested.
- explain how the components work in details, and the optimization.
- write down any APIs, data structures, or schema
- show the ability of trade off, like compexility and scalibilty, or performance and realibility.
- trade off over SQL database and NOSQL database.

## wrap up
- system bottlenecks and potential improvements
- recap of the design
- error cases
- logging, monitoring and alerting
- maintenance, operational cost.
- fault tolerance
- support scalibility

## tips
- clarification and understand of the requirements
- communicate with interviewer while thinking
- suggest multiple approaches 
- work with interviewer as teammate
- always be prepare
- don't jump into solution without clarify the requirements
- if stuck, ask for clarification or helps

## time allocation
- clarify and understand the question and estimation: 10 mins
- high level design: 15 mins
- devie deep: 20 mins
- wrap: 5 mins
