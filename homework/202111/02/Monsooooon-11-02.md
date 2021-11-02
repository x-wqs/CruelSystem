第一次学习System Design的Mock和记录笔记，暂时先用中文适应一下...

# 明确需求背景
设计一个 Donation/Fund Raiser的系统, 收集项目启动资金

明确设计场景时提出的问题：
1. 如何进行支付: 用户可以调用第三方API进行支付，暂时不用去管


# 确定Requirement
## Functional Req：

1. 支持用户注册 User Registeration -- <third party/local> -- id
用户是否可以选择使用FB/GG的账户进行注册，还是App自己去生成ID，不管怎么样都需要ID来标识出用户

2. 用户可以进行捐赠，系统记录捐赠记录 from XXX to XXX. 

3. 支持捐赠后发送email进行确认，或者分享到社交媒体上

## Non-Functional Req:

1. HA
2. Consistent
3. Concurrent update
4. Scalability

这里关于为什么并发更新会导致系统进入不一致状态，面试者没有解释清楚。

在Scalability上，需要对 总用户量，用户进行的交易数，交易保存的日期，每个交易占据的空间等等进行一些合理假设。然后估计出保存所有交易记录需要的存储大小。
由于容灾，预留存储等等问题，系统的容量需要进行调整。

随后可以估计出系统平均面临的QPS是多少

# 架构设计

几个关键的组成部分：
- Web/Mobile: 表示客户端
- LB: 进行负载均衡
- API Server: 执行业务逻辑，多个
- DB: 保存用户数据
- Cache: 加快DB访问（提到了Partiton和一致性hash但是感觉只是卖弄了一下，数据量这么小应该不需要分片）
- Payment API: 外部依赖

还提到了处理业务逻辑时保持一致性：分布式TXN，防止Payment API和DB更新有不一致

API-Gateway？ 这里把网关的功能也放在API Server里了。比较耦合

# API 设计

- sign up: 
  - 支持使用 username, email, phone number等等进行注册
  - 支持 third party注册
- login:
  - 这里没太听懂
- donate
  - userid, amount, cause

## 具体接口设计：Donate

主要application server内部要做什么，以及如何和支付API进行交互

考虑支付可能遇到的场景，以及每个场景下如何处理
1. Succ
2. Fail/Rejected
3. Expired

对应的场景
1. 如果支付成功了，需要同步更新DB里项目的余额，判断是否完成筹款
2. 如果支付失败，可以提示用户重试
3. 如果超时，xxxx （这里不清楚）

## 数据结构
这里主要是考虑内存
