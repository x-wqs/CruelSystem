## Detail Design

### Video Storage
- distributed file storage system
- replication for read-heavy
- upload to primary and sync to seconday

### Thumbnail Storage
- KV store, <VideoID:Timestamp, content>
- cache in memory, LRU

### Upload Video
- split to chunk, support resuming upload
- GOP alignment, Group of Picture
- upload centers close to users
- Safty: DRM, AES, Watermark
- TODO-S3: pre-signed upload URL

### Encode Video
- bit rates, formats, at most 4K
- notify uploader once encoding done

## Deduplicate Video
- waste video storage, cache, network, energe
- video matching algorithm
- skip low resolution or subpart of existing videos

## Stream Video
- MEPG-DASH, Moving Picture Experts Group, Dynamic Adaptive Streming over HTTP
- Apple HLS, HTTP Live Streaming
- Microsoft Smooth Streaming
- Adobe HDS, HTTP Dynamic Streaming

## Handle Errors
- upload failure
- malformated video
- encoding failure

## Reliability / LoadBalancer
- consistent hash
- elastic scaling
- replace down servers and distribute load
- minimize replication when server down

## Scalabilty / MetadataDB
- sharding with UserID
	- Pros: fast to find videos of some channel
	- Cons: too heavy for a hot user on a server, not evenly ditributed
- shareding with VideoID
	- cache metadata for hot videos

## Performance / Cache / CDN
- geographically distributed
- MetadataDB LRU, 20% hot
- CDN: distribute hot video geographically

## Enhancement

## Terminology
- segregate,分离
- deliberate,深思熟虑的
- stale，陈旧的
- widespread
- drawback,缺点，出口退税
- intertier
- complex technologies underneath the simplicity
- Youtube is enormous
- get buy-in,得到支持
- bunch of workers, 一群
- continuous
- contiguous
