## Step 2 - propose high-lvel design and get buy-in

Develop a high-level design and reach an agreement with interviewer
- Come up with an initial blueprint for the design. Ask for feedback.
- Draw box diagrams with key components - clients(mobile/web), APIs, web servers, data stores, cache, CDN, message queue
- Evaluate by calculations if blueprint fits the constraints

Go through a few use cases.
API endpoints and database schema could be designed at simple problem.

E.g. News feed system

- Feed publishing
  - user publishes a post
  - data is written into cache / database
  - post will be populated into friends' news feed
- news feed building - aggregating friends' posts 

## Step 3 - design deep dive
DNS <- User (web/mobile) -> (v1/me/feed) -> load balancer -> web servers -> post service -> post cache -> post DB
																		 -> notfication service
																		 -> Fanout service -> (get friend ids) -> graph DB
																		 				   -> (get friends data) -> user cache -> user db
																		 				   -> message queue -> fanout workers -> news feed cache

## Step 4 - wrap up
- identify the systems bottlenecks and discuss potential improvements
- a recap of the design
- error cases (server failure, network loss)
- operation issues
	- how to monitor metrics and error logs
	- how to roll out the system
- how to hande the scale curve
- propose refinements

### Dos
- ask for clarification
- understand the requirements
- no right / best answer
- think out loud
- suggest multiple approaches if possible
- design the most critical components first
