design ytb

# problem clarification
there will be many complex features for ytb, we need to ask the interviewer what are the most important features'
## functional
| questions | clarification |
| :---: | :---: |
| central func| upload vid and watch vid|
| client types| mobile app, web pages, tv|
| DAU| 5 million|
| average watching time per user| 30 min|
| support i18n| yes|
| vid resolutions| 1080p, 720p, 480p, 360p|
| encrypted communication| yes|
| file size | <1 Gb |
| use existing cloud services | yes | 

## non-functional
- fast uploading speed
- smooth streaming
- multi vid quality
- low infra cost
- HA, scalability
- multi clients

# back of envelope estimation
important thing: 
- what aspects should we consider? __qps, storage, network__ 
- what additional assumptions should we make?
  - how many uploads / day: 10% user 1 vid
  - storage cost
  - how many vid watches / user / day
  - avg vid size: 300 Mb
  - avg storage space / server
- daily new vid storage: 5 million * 10% * 300 mb = 150 mb * million = 150 TB
- CDN cost: __very high__
  - charged by data transmission, how about storage?
  - 0.02$ per GB
  - streaming cost: 5 million * 5 vid * 300 mb * $0.02 = $150000 / day
  - uploading cost: 150 TB * $0.02 = 150000 GB * $0.02 = $3000 / day

# high level design
decomposed by main functionalities
- vid uploading flow
- vid streaming flow
- vid info updating

we should use __BLOB storage__ service for vid (original or transcoded) storing and __CDN__ for fast streaming  

## components
- User Clients
- API servers & LB: handle basic user reqs such as login/sign-up/user info/vid meta data management, URL generator, and so on
- Blob storage for original and trancoded vids
- Trancoding servers
- vid/user MetaData DB & Cache
- Completion Queue & Handler(Event Consumer)
- CDN

## uploading flow
1. vid is uploaded to original blob storage
2. vid is processed by transcoding servers
3. trancoding finishes
   - send a completion events to queue
   - store transcoded vid to blob storage
4. completion handler consumes event and update MetaDB & Cache
5. API server send notification to user

## update metadata
user -> load balancer -> API servers -> metaDB & cache

## streaming flow
user -> CDN

# design deep dive
## transcoding
- why transcoding
  1. more compatible for diff devices & bandwidth 
  2. save storage space
   
- pipeline & tasks  
  - different user requirements, watermark, thumbnail images, HD vids
  - vid / audio / metadata are 3 diff parts and need seperate encoding work

## DAG
to be cont.
