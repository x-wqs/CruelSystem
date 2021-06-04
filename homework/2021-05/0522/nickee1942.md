# The Design of aPractical System forFault-Tolerant Virtual Machines
## Background intro
Introduce the primary and backup is a common approach to implementing fault-tolerant servers. The backup is a replica of the primary. When the primary fails, the backup could take over its action to avoid the failure on the client-side response. This is also the core design method of this paper. The author introduces some bottlenecks to implement the fault-tolerant system, such as ensure deterministic execution of physical servers is difficult. Then the virtual machine idea comes in as an approach. 
## Solution
## Basic fault tolerance design
The author implemented the fault-tolerant virtual machines using the primary and backup approach on the VMware vSphere 4.0. Which provides fault tolerance for any x86 operating systems.
## Deterministic replay impl
Use the VMware deterministic replay to solve three big challenges which impact the replicating execution of VMs running on any operating system
## Fault tolerantce protocol
To make sure that when the backup takes over the primary without missing any logs. An output rule has been used. The primary VM waits until the backup received and acknowledged the log entry to send out messages to the external world(client-side).
## Deal with the failure
Use the heartbeat sys to detect if any of the VM ( primary or backup) is down. Then the healthy VM could take over to do the action with the external world.
# Practical impl
##Start and restart VM
The server VM starts determined by the clustering service
## Manage the log
Need to avoid the primary log buffer being full. When it is full, there will be a delay to the external world. The solution is to make less latency between primary and backup by sharing the physical resource( CPU Ram, etc)
### To be continued
