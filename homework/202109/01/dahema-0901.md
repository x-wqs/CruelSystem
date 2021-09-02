# 秒杀系统

## 需求
- 瞬时大流量高并发 QPS: 平常 1000， 秒杀 100k
- 有限库存 不能超卖
- 黄牛恶意请求
- 固定时间开启
- 严格限购 

## 服务
- 秒杀服务
- 商品信息和库存服务
- 订单服务
- 支付服务

## 存储
- 商品信息表：product_id, name, desc, price
- 秒杀活动表：activity_id, product_id, price, quantity
- 库存信息表：stock_id, product_id, activity_id, stock, lock
- 订单信息表：order_id, product_id, activity_id, user_id, paid_status

## 秒杀操作
- 使用db：事务， update 语句自带锁
- redis：TODO
