## Primary/Backup Replication

### Failures

- "fail-stop" failure
- logic bugs, config error
- Other failures (e.g. earthquake)

### Challenges

- Primary replication fails
- Keep primary and backup replication sync

### Replication approaches

- State transfer
  - Primary sends new state to backups
- Replicated state machine
  - Primary sends new operations to backups
  - level of operations to replicate: 1) application-level, 2) machine-level
  - VMware FT replicates machine-level state

### VMware FT

![Screen Shot 2021-05-26 at 10.59.21 PM](https://tva1.sinaimg.cn/large/008i3skNgy1gqwswysjlgj31jc0u0e81.jpg)

left: primary, middle: backup, right: storage server

Interrupt -> VM FT -> send Loggins to backup VM FT

if the network between primary and backup breaks, primary or backup will "go live", primary or backup will be the new primary depending on who sends flag to storage server first

### Goal: how to behave like a single machine

- divergence sources: non-deterministic (e.g. get the current time)

- Input packets: 
- Timer interrupts

#### Output Rule

before primary sends output, primary wait for backup to ack all previous logs