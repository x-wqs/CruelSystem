# The Design of aPractical System forFault-Tolerant Virtual Machines
## Design alternatives
### Shard or non-shared
By default is shared, which means the primary and the backup shared the same disk. It is expensive and the writes need to communicate with the external world. The non-shared disk does not have the delay according to the output role. However, the 2 copies from different shared disks need to be synced explicitly when the first fault tolerance is enabled.
### Execute reads from backup VM disk
By default, the backup VM won’t read from the virtual disk. If change the design to let the backup reads the disk, it could reduce the traffic on the logging channel. The trade-off is it may slow down the execution. From the author’s test result, backup VM reads the disk only apply for logging channel bandwidth is limited
## Performance
The author did a bunch of experiments to investigate the performance. The performance includes basic performance, different network benchmarks performance
## Related work
The author also described other related work in this area. They are mainly from Bressoud and Schneider, who are also the main references the author take from
## Conclusion
A complete system in VMware vSphere is designed and implemented based on replicate the execution of primary VM via the backup VM. And this design gets an excellent performance.
