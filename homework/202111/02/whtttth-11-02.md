# Design donation/fund raiser app

## Functional Requirement

* Support individual user profile and track their donations

* User able to donate to multiple specific cause/event

* Support payment from user to cause through the system

* (optional) notification after donation, share on social media.

## Non Functional Requirement

* Highly available, should not be down anytime

* Consistency ensured, because payment envolved, neet to 
support concurrent payment from different users

* Scalability, handle user growth, data replication etc.

* Fault tolerance, tailover handling

## Back of the Envolop Calculation

* 10M users ---- 5 read/day 0.5 write/day 

* 500 read/s 50 write/s, the system is not read/write heavy, focus more on consistency and fault tolerance

* Storage 500K per user --- 5TB storage for single copy, consider partition/sharding for data

##  Data model

### User Metadata

key: userID\
value: { name, email, DOB, location, etc.}

### Event Metadata

key: eventID\
value: { name, desription}

### Event Payment history

TBD (could be denormalized for performance improvement)

### User Payment history

TBD (could be denormalized for performance improvement)

### SQL vs NoSQL 

TBD, mostly focusing on payment part

##  API definition

* sign up
* login
* donate


## High Level Design Component

* Client: Present user GUI, profile, provide apis etc.
* LB: Load balancing traffic to replicated application server from client
* API server: Handles and route user's api traffic including donate, update, signup, general queries etc.
* Payment server: handle payment request
* database: store user metadata, event/cause metadata, payment history
* cache: to be enabled in multiple places for performance improvement, for example database access on read path, user profile/history etc.

## High Level Design Block Diagram

* TBD




