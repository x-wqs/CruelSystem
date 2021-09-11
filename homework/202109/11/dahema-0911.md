# SCALE FROM ZERO TO MILLIONS OF USERS

## single server setup
- user: web browser or mobile app
- user call DNS to get IP address for specific website(domain names)
- user send HTTP request to web server by IP address
- the web server returns HTML pages or JSON response for rendering

## database
- SQL database, relational database management system (RDBMS). 
  - Relational databases represent and store data in tables and rows. 
  - support SQL language, like JOIN, WHERE, GROUP BY, ...
  - MYSQL, PostgreSQL
  - pros: structure data, mulitiple indexes, mature tech.
  - cons: scalability
- NoSQL(Non-Relational databases)
  - key-value stroes, graph stores, column stores, document stores
  - pros: scalability, replica, low latency, unstructured data, massive amount of data.
  - cons: not support SQL, TOOD more

##  Vertical vs horizontal scaling
- vertical scaling: add more power to servers
  - cost effective
  - less complex process communication and maintenance
  - less need for software changes
  - upgrade limitations, had hard limit for add CPU or memory
  - single point failure, no failover and redunancy
  - Higher possibility for downtime
- horizontal scaling: add more servers
  - scaling is easier from a hardware perpective
  - fewer periods of downtime
  - increased resilience and fault tolerance
  - increased performance
  - more complicated of maintenance and operation
  - increase initial costs
- which one to use
  - cost: Initial hardware costs for horizontal upgrades are higher.
  - Topographic distribution
  - reliability
  - Upgradability and flexibilit
  - Performance and complexity

## load balancer
- LB sit between client and server accepting incoming traffic
  - user and webserver
  - webserver and internal platform layer(like appliation servers or cache)
  - internal platform layer and DB
- LB distributing the traffic across multiple backend servers using various algorithms.
  - Most available CPU and memory
  - Least Connection Method 
  - Least Response Time Method
  - Least Bandwidth Method
  - Round Robin Method
  - Weighted Round Robin Method
  - IP Hash
- LB benefits
  - faster response as doing the load balance
  - increases the availability and fault tolerance
- double LB to avoid signle point failure

# Database replica
- master database generally only supports write operations.
- slave database gets copies of data from the master db and only supports read operations.
- benefits
  - Better performance
  - Reliability
  - High availability
- db go offline
  - only one slave db and it is offline, read ops go to master, if other slave dbs is alive, send to other slave dbs, replce the bad one
  - if master go offline, a slave will be promted to new master.
