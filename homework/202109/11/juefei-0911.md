## Design 1 million user application

### Data flow

browser -> domain name system (return ip)
browser(ip) -> web server (return html or json)

### Database

Usally, we want to split database since it make us to scale DB and service independently.

sql vs nosql

vertical vs horizontal scaling

### Load balancer

A load balancer evenly distributes incoming traffic among web servers

### Database application

Leader - Follower mode

advantages vs disadvantages

Advantages of database replication:
    • Better performance: In the master-slave model, all writes and updates happen in master
        nodes; whereas, read operations are distributed across slave nodes. This model improves
            performance because it allows more queries to be processed in parallel.
                • Reliability: If one of your database servers is destroyed by a natural disaster, such as a
                    typhoon or an earthquake, data is still preserved. You do not need to worry about data loss
                        because data is replicated across multiple locations.
                            • High availability: By replicating data across different locations, your website remains in
                                operation even if a database is offline as you can access data stored in another database
                                    server.
                                        In the previous section, we discussed how a load balancer helped to improve system
                                            availability. We ask the same question here: what if one of the databases goes offline? The
                                                architectural design discussed in Figure 1-5 can handle this case:
                                                    • If only one slave database is available and it goes offline, read operations will be directed
                                                        to the master database temporarily. As soon as the issue is found, a new slave database
                                                            will replace the old one. In case multiple slave databases are available, read operations are
                                                                redirected to other healthy slave databases. A new database server will replace the old one.
                                                                    • If the master database goes offline, a slave database will be promoted to be the new
                                                                        master. All the database operations will be temporarily executed on the new master
                                                                            database. A new slave database will replace the old one for data replication immediately.
                                                                                In production systems, promoting a new master is more complicated as the data in a slave
                                                                                    database might not be up to date. The missing data needs to be updated by running data
                                                                                        recovery scripts. Although some other replication methods like multi-masters and circular
                                                                                            replication could help, those setups are more complicated; and their discussions are
                                                                                                beyond the scope of this book. Interested readers should refer to the listed reference
                                                                                                    materials
