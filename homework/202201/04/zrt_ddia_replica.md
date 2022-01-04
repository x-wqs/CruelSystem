ddia -p191



DDIA 第五章：复制(Replica)

<https://vonng.gitbooks.io/ddia-cn/content/ch5.html>

* Motivation: 减少延迟、提高可用性、提高吞吐
* 难点：处理Replicas之间的数据变更
* 三种：single leader， multi leader， leaderless



## 无主复制

放弃主库概念，允许replicas直接接受client写入。

* 一种解决节点故障的方式是，写入时成功写入到k个服务器认为写入成功，读取时获取最新的数据版本。
* 节点重连
  * 客户端对陈旧的值重新写入
  * anti-entropy 一个后台进程，寻找差异并同步
  * quorum： w+r>n即可
* 运维：监控旧数据的出现频率
* 并发写入：
  * 只保留最后写入的数据：需要能够确定写入的时序
  * 按照分布式中的happened before的因果序来判断
    * vector clock
  * 合并多个同时写入的值