# Alex Xu Ch 12 Design a Chat System

## Functional Requirement

* Should support both 1-on-1 and group chat

* Support client access including mobile app, web and login from multiple client

* Push notifications for receiver

* Text length should be limited to 100K characters

* Should have an online indicator showing if the user is online, could support text only

* (optional) consider encryption if necessary

* Chat history should be stored permanently

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should able to generate newsfeed in real time for each user

* Fault tolerance, tailover handling

## Back of the Envolop Calculation

* 50M daily active users ---- 17500 qps

##  Data model

### User Metadata

key: userID\
value: { name, email, DOB, location, etc.}

### Message Metadata (1 on 1)

key: messageID\
value: { sender, receiver, text, creation_time, etc.}

### Message Metadata (Group Chat)

key: messageID, channelId\
value: { sender, text, creation_time, etc.}


### SQL vs NoSQL 

For metadata either choice is fine as we dont have strong relation between each individual record, users and event could by uniquely identified by unique IDs so a general k-v store is sufficient. 

MySQL generally provides ayto_increment which is handy for ID generation, but we can still use an indipendent ID generator


## High Level Design Component

* Client: Present user GUI, profile, provide apis etc.
* LB: Load balancing traffic to replicated application server from client
* service discovery: recommend the best chat server for a client based on the criteria like geographical location, server capacity
* Chat servers: facilitate message sending/receiving.
* Presence servers: manage online/offline status.
* API servers: handle everything including user login, signup, change profile, etc.
* Notification servers: send push notifications.
* key-value store: store chat history


## Long polling/ WebSocket

WebSocket connection is initiated by the client. It is bi-directional and persistent. It is widely used for sending async updates from server to client and bidirectional messages

Long polling, a client holds the connection until the new update is available or timeout, this allows client to receive updates immediately whenever there is an udpate from the serverside. However it has the problem of sender and receiver may not connect to the same chat server.

## Online Presence

The status will be registered once a user login and will be shown as part of user's state/profile. In general the status could be updated when user login and registered by service discovery, and it will be cleared when user logs out.

If there is a user disconnection, we can set up a heartbeat from online client to periodically send event to presence servers, if there is no update within a certain window the user will be considred disconnected and status will be set to offline.



