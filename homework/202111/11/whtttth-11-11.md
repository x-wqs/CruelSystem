# Alex Xu Ch 13 Design a Search Autocomplete System

## Functional Requirement

* Fast response time: when user types a query, auto complete must fast enough to present suggestions

* Autocomplete suggestions should be relevant to the search term

* Results returned by the system must be sorted by popularity or other ranking models

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should able to generate newsfeed in real time for each user

* Fault tolerance, tailover handling

## Back of the Envolop Calculation

* 10M daily active users 

* each user performs 10 search per day

* on average 20 requests sent for each query

* ~24000 QPS, kpeak 48000 QPS


## Trie Data Structure

We can use trie data structure to store the overall result and perform fast update/query in the system.

Trie is a tree like data structure that stores strings and is designed for string retrieval and search.

* A trie is a tree-like data structure.
* The root represents an empty string.
* Each node stores a character and has 26 children, one for each possible character. To
save space, we do not draw empty links.
* Each tree node represents a single word or a prefix string.

In each node we store the relative character in the string and also the frequency it has been queried to this point. 

To get the top k most seached quries at each node we can tracerse the subtree and perform a sort to get the top k result.

To speed up the process we can also maintain the top queries at each node so we dont need to search everytime, but that requires more storage for the data.


## Data Gatehring Service

When a user performs the querie, the action will also be collected to update the data. 

### Analytics Logs

First we store the raw data in the analytic logs, which will be append only and not indexed, so the service could capture huge traffic without impact much on the performance.

### Aggregators

We need to aggregate data so it can be easily processed by our system. 
We can choose to aggregate data in a shorter time interval as real-time results are important. We can also choose to aggregate data asynchronously with less frequency.

### Aggregated Data

In our service the aggregated data will contain updated fruquencies for different strings and be applied to the trie data structure.