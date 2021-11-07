# Design Newsfeed

## Functional Requirement

* Newsfeed is generated based on the posts from people, pages, and groups that a user follows

* User could have many friend or follow large number of pages/groups

* Feed contains images, videos, text

* Allow appending new posts for active users when they arrive

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should able to generate newsfeed in real time for each user

* When a new post is created it should appear in user's feed soon

* Fault tolerance, tailover handling

## Back of the Envolop Calculation

* 300M daily active users ---- 17500 qps

* 500 posts for each users feend and average size is 1KB, then storage requirement is 500KB/user---150TB for total sorage, such that for all active users they need to be stored in memeory and we need multiple mahcines to store all of them 

##  Data model

### User Metadata

key: userID\
value: { name, email, DOB, location, etc.}

### Entity Metadata (group, page)

key: entityID\
value: { name, type, desription, creationdate, email, etc.}

### Follow

key: userId, friend/entityId

### FeedItem

key: itemId
value: { userUd, Content, mediaId, creationDate, etc.}


### SQL vs NoSQL 

For metadata either choice is fine as we dont have strong relation between each individual record, users and event could by uniquely identified by unique IDs so a general k-v store is sufficient. 

Media resources could be stored in blob storage for persistent storage and could be easily referenced by an ID, the ID will be stored in the metadata along with other attributes of the feedItem


##  API definition

* getUserFeed(api_dev_key, user_id, since_id, count, max_id, exclude_replies)


## High Level Design Component

* Client: Present user GUI, profile, provide apis etc.
* LB: Load balancing traffic to replicated application server from client
* API server: Handles and route user's api traffic including donate, update, signup, general queries etc.
* web server: maintain connection with user and transfer data in both directions
* Metadata database: store metadata of users, pages and groups
* Posts database and cache: To store metadata about posts and their contents.
* Video and photo storage, and cache: Blob storage, to store all the media included in the posts.
* Newsfeed generation service: To gather and rank all the relevant posts for a user to generate newsfeed and store in the cache. This service will also receive live updates and will add these newer feed items to any user’s timeline.
* Feed notification service: To notify the user that there are newer items available for their newsfeed.

## Feed generation: Pull vs Push

"Pull" model or Fan-out-on-load: This method involves keeping all the recent feed data in memory so that users can pull it from the server whenever they need it. Clients can pull the feed data on a regular basis or manually whenever they need it. Possible problems with this approach are a) New data might not be shown to the users until they issue a pull request, b) It’s hard to find the right pull cadence, as most of the time pull requests will result in an empty response if there is no new data, causing waste of resources.

"Push" model or Fan-out-on-write: For a push system, once a user has published a post, we can immediately push this post to all the followers. The advantage is that when fetching feed you don’t need to go through your friend’s list and get feeds for each of them. It significantly reduces read operations. To efficiently handle this, users have to maintain a Long Poll request with the server for receiving the updates. A possible problem with this approach is that when a user has millions of followers (a celebrity-user) the server has to push updates to a lot of people.



