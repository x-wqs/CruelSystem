## Detail Design
- Trie Data Structure
	- Empty root node
	- Find a prefix: O(P), P: length of a prefix, N: total number of nodes in a trie
	- Traverse substree: O(C), C: number of descendant of a given node
	- Sort descendant and get top k, O(C * LogK)
	- Type "tr", return { tree: 10, true:35, try: 29 }
	- Total time: O(P + C + C * LogK)
	- TODO: merge node if only one branch
- Optimized Trie
	- Limit the max length of a prefix, max 50, not index long sentence, O(1)
	- Cache top search queries in each node, trading space for time
	- Find the prefix node: O(1)
	- Return Top K, O(1)
- Data Gathering Service
	- Build Trie from analytics or logging
	- Option 1: update Trie in real-time (on every query), Twitter
	- Option 2: Schedule updating, Google
	
- AnalyticsLogs ---> Aggregator --> AggregatedData ---> Workers ---> WeeklyUpdate ---> TrieDB ---> TrieCache

- AnalyticsLogs: append only, <query, timestamp>
- Aggregator: aggregate in a shotr time for real-time application like Twitter, or weekly for others
- AggregatedData: <query, weekStartTime, frequency>
- Workers: asynchronous job at regular intervals
- TrieCache: distributed cache system, weekly snapshot of TrieDB
- TrieDB: persistent storage
	- DocumentDB: serialized
	- KV Store: <Prefix, TopQueries>, distributed KV Store
	
- App/Client ---> LB ---> QueryService ---> FilterLayer ---> TrieCache ---> TrieDB

- QueryService
	- light-weight AJAX
	- replenish data back to cache when offline or out of memory
	- cache in browser with TTL like Google, cache-control: private, max-age=3600
	- data sampling, log every N queries
	
- Trie Operations:
	- Create: by aggregated Analytics Logs
	- Update: weekly
	- Delete: hateful, violent, dangerous
	
# Sacalability / Sharding
	- Option 1: range based, on the first character, or multi-levels, but uneven distribution
	- Option 2: shard by frequency or server capacity
	- Option 3: hash prefix

## Fault Tolerance
- Update Secondary when using Primary
	
## Enhancement
- send request when stopping for 50ms
- start sending request after 3 character typed
- local browser cache
- push cache to CDN
- treading / real-time search, complicated, more weight for query, stream system
- TODO: personalization	
