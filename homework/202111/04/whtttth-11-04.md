# Design donation/fund raiser app Part 2

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

### Transaction

key: transactionId\
value: {amount, userId, eventId, status, timestamp}


### Event Payment history

list of transactions for each event(could be denormalized for performance improvement)\
key: eventId\
value: {transactionId1, transactionId2}

### User Payment history

list of transactions for each user(could be denormalized for performance improvement)\
key: userId\
value: {transactionId1, transactionId2}

### SQL vs NoSQL 

For metadata either choice is fine as we dont have strong relation between each individual record, users and event could by uniquely identified by unique IDs so a general k-v store is sufficient.


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

## Potential Components

* UniqueId generator, act as service generating uIds for transactions instead of using databases


## General Comment

* Start with general component when first start high level design, do not need to include specific optimization like cache in the begining, then add these optimization later into details
* think about scalability, and what is actually feasible for the current problem
* Besides technical component, think about customer/business requirement as well
* Think about user experience and how it looks like based on the current design, then consider what would happen when component fails and how we should handle it
* Be concious about the end to end user experience look like, and do not just focus on the backend component
* After the design, validate agains the initial assumptions




