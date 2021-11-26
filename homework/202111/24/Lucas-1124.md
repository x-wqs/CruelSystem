## Functional Requirements
- friends can see a user online after login
- friends can see a user offline after leaving

## Non-functional Requirements
- high reliability/availability/scalability
- low latency

## Capacity Estimation
- 500M DAU
- 10K status changes / second

## System API

## Database Design

## High-level Design

## Detail Design

### Realtime Platform / RP
- The user will setup WebSocket with RP after login
- The user will subscribe all topics for status of friends

### Presence Platform / PP
- PP will send status change to RP

### Online Status
- The user will send heartbeat periodically
- RP receives <MemberId, Timestamp>
- RP checks whether MemberId exists unexpired in KV store
- If no, it means the user comes online
- PP will publish an online event to the topic on RP
- PP will add <MemberId, Expiration> in KV store
