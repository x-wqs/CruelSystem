# Design Google Drive

## Requirement clarification

### Functional features
- Add files.
- Download files.
- Sync files cross devices.
- File revision.
- Share files.
- Sharing notification.
- How about authentication and authorization?

### Non-functional features
- Reliability
- Fast sync speed
- Avoid unncecessary network traffic.
- Scalability
- High availability

### Back of envelope estimation
- 50M users with 10M DAU.
- 10G free space.
- Average 2 files upload per day.
  - 500KB per file.
- 1:1 read-write ratio.
- Total space needed: 50M*10G = 500PB
- QPS for upload: 10M*2/(24\*3600) = 240
- PeakQPS = QPS*2 = 480

## Propose high-level design

### Single server setup
#### Components
- A web server to upload and download.
- A database to keep track of metadata.
- A storage server to store the files.
  - Files of each user are stored in a different folder.
  - How to support (de-)encryption?
  -
#### API
- Simple upload
- Resumable upload
  - For large files that could be interrupted by user or accident.
    - How?
    - Backend need to monitor the upload progress.
      - If user is requesting to resume a previously interrupted upload. Then the server will send back parameters regarding the offset of file to continue uploading, etc.
- Download
  - Path as parameter.
- File revision
  - Path of file
* User authentication and HTTPs is needed for all APIs. *

### Scale!
#### Storage
- Option 1: Shard the data
  - Add more storage servers, shard by user_id, e.g. user_id%4
- Option 2: Amazon S3
#### Web server
- Add load balancer and more web server.
  - to improve the throughput.
  - load balancer can distribute traffic evenly.
  - make the system more reliable and avoid single point failure.
- Use Amazon S3 with geo-replica to make the system more reliable.
- Setup database replica and sharding.

#### Handle conflicts
- Resolution, first version completes processing wins. Later version from different user needs to be resolved manually.

### High-level design
- Block servers: as a pre-processing before upload to and download from S3. It split file for uploading and joins blocks for downloading.
  - Who is responsible to splitting or joining?
- Cloud storage: S3, etc.
- Cold storage: even cheaper servers that stores files that are accessed infrequently (cold.).
- Load balancer
- API servers
- Metadata database
  - stores metadata of users, file, blocks, versions, etc.
- Metadata cache
  - for fast retrivial of hot metadata without accessing the database.
- Notification service
  - The publisher and subscriber service that allows data to be transferred from notification service to user.
- Offline backup queue
  - A queue that make sure user is always up to date on the files.
  - A design choice, maybe user don't want to always be up-to-date, but only on the files that he choose.
  - This make the service stateful, not sure it is a good idea.
  - **But what if user do want to be up-to-date, is there a better way to achieve it?**

## Deep dive

### Block servers (in charge of splitting and joining large files for download and upload. And encrypting.)
* Delta sync, only transmit the modified blocks of the file to save network traffic.
* Compression, compress and decompress before and after transmitting.
The cloud storage stores the splitted and encrypted blocks.

### Strong consistency
* Consistency in mem caches.
  * Mem caches are eventual consistency by default, i.e. different replicas might have different data. To achieve strong consistency, we need to make sure:
    * Data in cache replica and the master is consistent.
      * Consistency between slave-caches and master cache.
    * Invalidate caches on database write to ensure cache and database hold the same value.
      * Consistency between cache and database.
  * Use RMDB for ACID (atomicity, consistency, isolation, durability).

### Metadata DB
#### Schema
- User
  - id, email, username, joined_at
- Device
  - id,push_id, last_logged_in_at
- Block
  - id, file_version_id, block_order
- Workspace (the root dir of the user, why do we need this?)
  - id, owner_id
- Share (NEW)
  - from_user_id, to_user_id, file_id
- File
  - id, file_name, relative_path, checksum, latest_version, workspace_id, last_modified, created_at.
- FileVersion
  - id, file_id, device_id, version_number, last_modified.
    - Note, last_modified field here is necessary and not duplicated.
    - It keeps info of previous version which becomes readonly after a new version is generated.

### Upload flow
Two things to accomplish:
- Upload file to cloud storage
  - Need to (create) mofidy meta data first before uploading.
  - Maintain an upload status (uploading, uploaded) in the meta db.
- Notify client when the file is uploading and uploaded.

### Download flow
Two things to accomplish:
- Users auto-update the changed files.
  - If client is online while a file is changed by another client, it will be notified and pull the data.
  - If client is offline when a file is changed, it pulls the change when it gets online.
- Download
  - First get a list of changed files from API server.
  - For each file, request the blocks from block server.
  - Block server download from block server.
  - Client download the blocks and build the file.
    - Why not let client access the block server?
      - Block server need to decrypt the blocks.
  - What about client decrypt the file?
    - Less secure to store the keys on client side?
      - Security is kind of beyond the scope of this design.

### Notification service
Allows data to be transferred to client as events happen.
* Long polling. Used by dropbox.
* Web Socket: persistent connection between client and server. bi-directional communication.
Long polling is preferred here because we don't need bi-directional communication (not like chatting app).
And the notification is less frequent.
Each client need to establish a long poll connection to the notification service.


### Save storage space
* De-dup data blocks. Hash.
* Limit the number of revisions.
* Overwrite old versions if exceed the limit.
* Move infrequent data to cold storage. (cheaper, e.g., S3 glacier).


### Failure handling
* Load balancer failure: Secondary load balancer become active and take over the traffic. Monitor each other by heatbeat.
* Block server failure: other servers pick up unfinished job.
  * How? Who is managing the unfinished and resend jobs?
* Replicate, slave, etc.
