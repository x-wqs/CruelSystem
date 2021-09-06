## source


> https://pdos.csail.mit.edu/6.824/papers/vm-ft.pdf 1~10

> The Design of a
Practical System for
Fault-Tolerant Virtual Machines


## abstract
> describe basic design, alternate design choices, implementation detail of our practical system.

## introduction
> common approach:
secondaries keep indentical with primary nearly on the time, shipping all state of CPU, memory, I/O devices cost large bandwidth

> our approach:
model servers as deterministic state machines, use extra coordination to ensure primary and backup kept in sync. extra informatige system
> building fault tolerant storage

> extra data is needed for non-deterministic operations, hypervisor capture all these extra informations.

> deterministic reply and FT support only in uni-processor, multi-processor still work in progress, as access to standared memory is a non-deterministic operation

> only deal with fail-stop failures, avoid incorrect externally visible action before server failures detected

## basic ft design
> all network/keyboard/mouse inputs come to primary VM, then primary send to backups via network connection know as logging channel. outputs of backup are dropped by hypervisor, only primary's output returned to clients.

> detect failure by heartbeating and monitoring traffic on logging channel

### deterministic replay implementation
#### challenge
- capture all input and non-determinism including undefined behaviour of x86 microprocessor

- correctly applying non-determinism to backup machines

- don't degrade performance

> deterministic replay records inputs and all possible non-determinism in a stream written to log file. During replay, event is delivered at the same point in instruction stream

### FT protocol
> output to external world is delayed until primary recieve acknowledgment that backup has recieved log entries about the output operation

### detecting and responding to failure
> UDP heartbeating between servers and monitoring traffic log

> to avoid split-brain, make use of shared storage. When primary/backup wants to go live, it does atomic test-and-set operation on shared storage

## practical implementation of FT
### starting and restarting FT VM
> clone rather than migrate.

> system choose best server location of new backup

### managing the logging channel
> backup stop excution when channel empty, don't affect client

> primary stop excution when channel full, affect client

> have mechanism to slow primary when backup get too far behind

### operation on FT VM
> when primary power off, backup should not attemp to go live

### implementation issues for disk io
> simultaneous same memory page disk operation is non-determined. Solved by force IO race(rare) sequential

> disk operations across memory race with application's access to memory. Solved by bounce buffers(page protected is trap?)

> when primary tear down, new primary don't know pending IOs of backup during go-live process. Solved by just re-issue(no problem even they have been completed successfully?)

### implementation issues for network IO
> disable non-determined asynchronous network optimizations

> reduce delay for transmitted packages. done by sending and recieving log entries without thread context switch

## design alternatives
### non-shared disk
> disk considered internal part of VM: more cost of sync, no communication when split brain without third part server

### execute backup reading
> less bandwidth cost and more time cost, following with some sync problem

## performance evaluation
### basic performace
> performance of FT/non-Ft > 0.94

### network benchmark
#### challenge
> high speed have high log and replay rate

> heavy log traffic

> output rule delays sending packet, high delay causes less bandwidth(such as TCP)

## related work
### previous work
#### pros
> demonstrated feasibility of keeping backup in sync

#### cons
> impose notion of epoch as they cann't sync efficiently

> require primary wait for backup recieve and acknowledge all previous log entries(actually only output must be delayed)
## terms
> hypervisor 管理程序

> uni-processor 单处理器

> conjunction 关联

> epoch 时段

> disverge 分散

> analogous 类似的

> compensate 补偿

> susceptible 易受影响的

> disrupt 打乱

> VMotion: VMware 开发出的一项独特技术，它将服务器、存储和网络设备完全虚拟化，使得正在运行的整个虚拟机能够在瞬间从一台服务器移到另一台服务器上

> quiesce 停顿

> subtle 微妙的

> idempotent 幂等

> emulate 仿真

> idle 空闲

> feasibility 可行性

> impose 强加

> derive 派生
