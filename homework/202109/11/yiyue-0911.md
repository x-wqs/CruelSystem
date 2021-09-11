### single server
- web server <- user(web browser / mobile app) <-> DNS
- request flow
  - user -> domabin names -> DNS
  - DNS -> ip address -> user
  - user -> ip -> web server
  - server -> HTML -> user
- traffic source
  - web application: server-side languages to handle business logic + client-side languages for presentation
  - mobile application: HTTP protocol

### database
- web server <-> database
- databases
  - relational databases - represent and store data in tables and rows - can perform join operations across different database tables
  - Non-Relational databases (NoSQL databases) - join operations not supported
    - key-value stores
    - graph stores
    - column stores
    - document stores
  - Non-relational databases suitable for
    - low latency application
    - unstructured data
    - only serialize & deserialize data (JSON, XML, YAML, etc)
    - massive amount of data

### vertical scaling vs horizontal scaling
- vertical scaling means adding more power (CPU, RAM, etc) to server
- horizontal scaling means adding more servers
- Traffic is low - vertical scaling
  - pro: simplicity
  - cons
    - hard limit. Impossible to add unlimited CPU and memory to a single server
    - No failover and redundancy. If one server down, website down
- Horizontal scaling - large scale applications

### Load balancer
- Evenly distributes incoming traffic among web servers
- user -> load balancer -> servers
- users connect to the public IP of the load balancer directly. Web servers are unreachable by clients
- Private IPs are used for communication between servers - unreachable over internet. 
- If one server goes offline, traffic will be routed to the other server. A new healthy web server will be added to the server pool
- If website traffic grows rapidly, add more servers to the server pool.