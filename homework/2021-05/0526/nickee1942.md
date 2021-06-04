# Lecture 4: Primary-Backup Replication (2nd part)
## How the system gonna work.
## Ex. Two machines are running
The goal is running like a single machine.
Non deterministic is a problem. As the responses probably return at different times. So the order is non-deterministic. 
### The system has disallowed the multicore.
### Output role
To pretend the primary and the backup don’t sync, the primary will not send the output to the client-side. It waits until the back up received and notice the primary through the logging channel.
### Deal with the failure
Use the heartbeat sys to detect if any of the VM ( primary or backup) is down. Then the healthy VM could take over to do the action with the external world.
Practical impl
## Start and restart VM
The server VM starts determined by the clustering service
## Manage the log
Need to avoid the primary log buffer being full. When it is full, there will be a delay to the external world. The solution is to make less latency between primary and backup by sharing the physical resource( CPU Ram, etc)
## Operation on fault-tolerant VMs
This section discusses the possibilities when some operations happen on the primary VM and this operation should not happen, simultaneously on the backup VM. Here the author uses the VMotion to solve this concern.
## Problems
For the disk IOs, it has a way to detect the IO races. Bounce buffers is one of the solutions that their memory size is the same. And for improving the performance of the network IOs. It could use two optimizations, improve the clustering to reduce the VM traps and interrupts, or trying to reduce the delay for transmitted packets.
