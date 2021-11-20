# 0000-system-dropbox
- Google Drive, Dropbox, Microsoft OneDrive, Apple iCloud

## Functional Requirement
- upload/download/sync/share files
- see file revisions
- send notification for a file change
- optional: encryption

## Non-functional Requirement
- high reliability/availability/scalability
- low latency
- bandwidth

## Capacity Estimatioin
- 50M registered users, 10M DAU
- Storage 50M * 10G = 500 PB
- upload 2 files / day, 500KB / file, 1:1 read to write ration
- Upload: 10M * 2 / 24 / 3600 = 200 QPS, Peak 400
- 
## System API
- int upload(token, data, resumable)
- bytes[] download(token, path)
- JSON listRevision(token, path)

## Database Design

## High-level Design

-           +---> Notif <---+
- Client ---+---> LB ---> AppSvr ---+---> MetadataCache ---> MetadataDB
-           |      |
-           +---> BlockServer ---> CloudStorage ---> ClodStorage

## Detail Design

## Teminology
- narrow down
- synchronization
