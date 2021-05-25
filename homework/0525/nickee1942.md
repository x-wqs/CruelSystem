# Lecture 4: Primary-Backup Replication
## Fault-tolerant challenges
### 1. If a failure happens, will the primary fail? There are some possible solutions such as “split-brain”.
### 2. The Sync problem between primary and backup. Some detailed solutions have to make the sync correctly. For example, the order for applying changes is important; properly deal with the non determinism 
### 3. Failure over

## Two approaches
### 1. state transfer, transfer from primary to the backup
### 2. state machine replication. The primary sends operations to the backup
Both two approaches have the same core point. Two approaches deal with whatever happens in the primary, it should sync to the backup. This could make the backup identity the same as the primary

## What level of operations need to replicate
## VM fault-tolerant
## Transparent replication
## Overview of the Virtual machine fault tolerant system
The hardware is at the bottom of the structure. Above it is the VM fault-tolerant, which includes the logging channels. Then there are Linux systems run on the VM with applications. The client side sends a request to the hardware. Then level by level, from hardware to VM fault-tolerant, to the system, and finally to the application. It fetches the matched information from the application and sends back the information to the client.
