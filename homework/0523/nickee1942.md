# The Design of aPractical System forFault-Tolerant Virtual Machines
## Deal with the failure
Use the heartbeat sys to detect if any of the VM ( primary or backup) is down. Then the healthy VM could take over to do the action with the external world.
Practical impl
## Start and restart VM
The server VM starts determined by the clustering service
## Manage the log
Need to avoid the primary log buffer being full. When it is full, there will be a delay to the external world. The solution is to make less latency between primary and backup by sharing the physical resource( CPU Ram, etc)
## Operation on fault-tolerant VMs
This section discusses the possibilities when some operations happen on the primary VM and this operation should not happen, simultaneously on the backup VM. Here the author uses the VMotion to solve this concern.
## Problems for disk IOs
The author discusses the implementation issues for the disk IOs, The implementation has a function to detect the IO races. Bounce buffers, which have the same size as the memory, are used to solve the race of memory access.
## Problems for Network IOs
This section is a discussion for improving the performance of the implementation of network IOs. Here the author mainly provides two optimizations. One is to improve the clustering to reduce the VM traps and interrupts. The other is trying to reduce the delay for transmitted packets.
### To be continued
