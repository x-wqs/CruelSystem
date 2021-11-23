## Database Design
- users(id, name, createdAt)
- workspaces(id, userId, shared, createdAt)
- files(id, name, type, path, checksum, workspaceId, latestVersion, lastModifiedAt)
- versions(id, number, fileId, deviceId, lastModifiedAt)
- blocks(id, versionId, order)
- device(id, user, lastLoginAt)

## High-level Design

-           +---> Notif <- MessageQueue --+
- Client ---+---> LB ------------------> AppSvr ---+---> MetadataCache ---> MetadataDB
-           |      |
-           +---> BlockServer ---> CloudStorage ---> ColdStorage

## Detail Design

### BlockServer
- delta sync, only modified blocks
- compression, gzip or bzip3
- TODO: [1,2,3,4], [5,6,7,8], [9] -> U1: [1,2,3,x],[4],[5,6,7,8],[9]

### Strong Consistency
- difference among clients is not acceptable
- cache and master are consistent

### Upload Workflow
- Upload Status: Client1 ---> AppSvr ---> MetadataDB ---> NotificationService
- Upload Content: Client1 ---> BlockServer ---> CloudStorage ---> AppSvr ---> MetadataDB ---> NotificationService

### Download Workflow
- TODO


### Sync Conflicts

## Teminology
- narrow down
- synchronization
- guru
