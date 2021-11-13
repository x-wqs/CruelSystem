design an auto complete search system

# clarification
- language cn or english or... 
- support top k or fixed size
- what's the data? data size? will data change? 
- fuzzy matching or not? 
- only prefix or any parts of provided data
- how to determine the result?  matching degeree and frequnency? 
- spell check? 
- allow special characters or alphabet only? 
- allow space? 
- real time display or other users input or not? 
- word length max average
- DAU  ~10 million

# requirements
## (non) functional
- fast response ~0.1 sec
- relevant by matching
- sorted by ranking func
- data changes as user inputs
- scalable
- HA

# capacity
QPS: 10 million * [10 queries per user] / 0.1 (million sec/day)  = 1100 QPS
Input size: 20 bytes per query (ASCII) 

**NOTICE**: every stroke is an query

search? input=c

search? input=ca

search? input=car

 therefore
- Actual QPS = 1100 * 20 = 22k QPS
- Peak QPS = 40k QPS
- Daily new data: 20% * 10 million * 10 queries * 20 bytes = 400 MB = 0.4 GB

# high level design
user input data is used to do actual searching and saved and aggregated as new dataset

## data gathering service
simplified model: a map of searched strings frequency
##  query service
same map

we can use a single relational db and SQL to get the top 5 matches

# dive deep
## trie
use trie to achieve better performance

- label strings and frequency in leaf node
- top 5: go down by prefix and search the subtree
- time complexity O(len(word))  + O(size(subtree))  + sort results
- cache result on every node to speed up

## data gathering
batch or stream?  or both? 

- input data is represented as analytical logs
- aggregate data in time window (Flink) 
- store data in data warehouse
- use big data tools to build weekly trie
- trie cache & trie db (mongodb) 

## data querying
- LB -> api servers -> trie cache -> trie db
- 01 switch of different ver of trie caches

## speed up
- use AJAX at frontend code to avoid whole page refreshing
- browser cache
- data sampling to build trie
- trie partition by prefix characters
