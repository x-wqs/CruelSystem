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
