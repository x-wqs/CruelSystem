    leture4 Primary-Backup Replication:
    2 approach to handle failure. 
        1. State transfer 
        2. replicated state machine.
    leture4 is talk about replicated state machine.
    关于避免fail-stop 和 incorrect status:
        replication 不能帮你避免程序bug和配置错误等。
        可以帮你避免机器意外损坏。

    keep primay backup in sync
        在primay和 buckup上做同样的state transfer
        no need send data to backup, just send operation. That will make the backup less expensive
        in machine level replication. we can use VM. 
    VMFT:
        operation will be send to a logging channel to backup machine.
        need repair.
