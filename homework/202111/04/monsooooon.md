# 11-04 

[视频链接](https://www.youtube.com/watch?v=XT0JHVZzkW0)

# 接口设计
## donate
### 接口被调用后的状态
1. Succ
2. Rejected/Retries(包含次数上限)
3. Payment Expires
### TXN如何记录
可以使用Relational Model, 也可以使用Document Model, 不过感觉Relational 更加适合一点?

Transactions

| TxnID | UID | Amount | Status | ...
这里应该还要有 Timestamp(start, end), Target Project ID

### 如何处理异常情况
#### 支付API超时
调用了支付的API,但是等待超时了,可能是因为支付API挂了,或者网络问题没有收到回复. 我们没法知道到底支付成没成功.
如何处理这种情况?

1. 首先, 应该retry对支付API的请求, 只要重试成功了就可以了 (没接触过支付场景,这里是猜的)
   1. 这里具体是怎么知道重试的是同一个TXN的?
2. 如果错误扣费的话一般会引起用户比较大的不满,可以先不管是否正式支付,一律告知用户支付成功,然后等支付API恢复了再进行确认.如果没有支付成功再次告知用户.


# FeedBack
1. Scalability
2. Message Queue -> Handle burst of incoming messages
3. State management -> Handle loss of in-memory data
4. 明确database, nosql or sql
5. request async flow vs sync flow: 如果支付api的scalability比较差怎么办(抗不了比较大的流量)? 同样可以使用 MQ 来限制请求速度, TODO: 需要考虑一致性的问题. 
   1. 使用asycn flow的场景: 下游服务的稳定性/qps限制相对比较低, 但是上游的请求可能有峰值, 防止block user的操作 
6. 请求重试->需要txnid, + backoff strategy, should charge exactly once

# Summary
1. 熟悉了一下SD的流程, 不同的层次进行设计
2. 对于一些常见的原则(分布式一致性, partition, replication) 都可以尽量体现在设计中
3. 印度口音太重又没有字幕,信息获取的量有点低, 可能要适应一下...
