# Primary/backup replications

plan:

- failures
- challenge
- 2 approaches
  - state transfer
  - replicated state machine
- case: VM FT

## Failures

- fail-stop failures(failure will stop computer)
- logic bugs(eg:divided by 0), configuration error, malicious
- others(earthquake)

## Challenge

- has primary failed?(network disconnected, split-brain system)
- keep primary/backup in sync(apply changes in order, non-determinism)
- failover

## 2 approaches

### state transfer

state changes(result) -> backup

expensive when lot of changes generated

### replicated state machine(RSM)

operation -> backup

general approach

make all operation deterministic
non-deterministic operations are not allowed

level of operations to replicate:

- application-level operations(file append, write)
- machine-level operation (register)
  transparent! use virtual machines

### hybrid approach

use RSM to sync, when backup failed, use state transfer to clone

## VM-FT

exploit virtualization

support of multi-processor might use state transfer

VM-FT capture hardware interrupt

when logging channel suspend, primary and backup use storage to perform test-and set operation ot check if it can be upgraded

when there is no replica, the primary will block client request until backup finished

## Others

goal: behave like a single machine

divergence sources(non-deterministic)

input packets
time interrupts

### multi-core

condition race is hard to sync to the backup(it has been solved now?)

### non-deterministic

do some pre-work and stick the result to the registers

### performance

FT is similar to non-FT

primary have to wait the backup, so the bandwidth is reduced
