## 设计秒杀系统
- https://www.cnblogs.com/54chensongxia/p/13609148.html
- https://gongfukangee.github.io/2019/06/09/SecondsKill/
- https://www.infoq.cn/article/ypqschsrdsk8bv5nhny4
- https://segmentfault.com/a/1190000020970562

## 功能需求
- 商品有限，供不应求，防止超卖
- 防止恶意刷单
- 指定时间开放
- 订单长时间未支付，释放商品，补充到库存
- 保证先来流量抢到商品

## 非功能性需求
- 并发请求高
- 业务量大，不影响其他系统
- 防止攻击
- 防止恶意操作，不停点击秒杀按钮

## 应用接口
- int createTempOrder(userId, productId)
- int cleanTempOrder()
- int payOrder(userId, orderId)

## 数据库设计
- products(id key, name, stock, sold, version)
- orders(id key, productId, createTime)

## 架构设计
- TODO

## 模块设计
- TODO

## 秒杀流程
- TODO

## 其他方面
- 可用性
- 扩展性
- 安全性
- 限流
- 监测
