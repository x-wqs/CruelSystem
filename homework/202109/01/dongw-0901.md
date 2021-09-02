秒杀系统
Characteristics
1.very huge instant concurrency traffic
 2. inventory is limited
 3.the purchasing process is easy
 4. countdown before the start of the purchasing

Main problems to be solved:
1. how to handle the instant huge traffic:
 the qps for databsase ,service is limited, e.g single database can handle 1000 qps. we need to estimate the concurrency traffic
 2. inventory for the goods is limited, we shouldn't sell more than the inventory
 we have to ensure we will not sell more than provided and we shouldn't sell less than provided.
 3. stablility and high availabilty:
 the system should be reliable and stable when the event starts and it shouldn't affect other services.
 
Details:
1)high traffic:
filter out invalid requests in different layers.
CDN -> frontend read layer -> backend write layer -> DB, the requests should be reduced as the layer goes. We try to filter 
out those invalid requests at each layer. Specifically, we could 
a) static page technology: cache the video, audio of the goods which is very consuming to cdns, the users can get the those resouces from their nearby cdn. 
b) cache as much as the information of the goods to reduce the internet traffic, like redis which can handle 10,000ish qps
c) Asynchronous: mainly refers to the backend and db side
d) Peak load shifting: use message queue to create the order


