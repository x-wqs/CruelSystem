# 0000-alex-system.md

# Ch2 - 容量估算

## 需求
- 估算系统容量和性能需求
- 2的次方，延迟数字，可用性数字
- 1K, 1M, 1G = 10^9, 1T = 10^12, 1P = 10^15
- 从加州发送网络数据包到荷兰并收到确认，150ms
- 内存块，磁盘慢，避免磁盘寻址
- 简单压缩非常快，网络发送数据尽量先压缩
- 数据中心跨区域，传输成本很高
- Amazon/Google/Microsoft, set SLA at 99.9%, 8.77h/y

## 估算 Twitter QPS和存储需求
- 1M active users / month
- 50% of users use it daily
- a user posts 2 tweets / day
- 10% of tweets 带有多媒体
- 数据存储5年
- DAU = 300M * 50% = 150M
- QPS = 150M * 2 / 24 / 3600 = 300M / 10^5 = 3K
- Peek QPS = 6K
- Tweet Size: id 64 bytes, text 140 bytes, media 1MB
- Media Size: 150M * 2 * 10% *1 M = 30TB / day
- Total Storage: 30T * 365 * 5 = 30T * 2K = 60PB

## 提示
- 容量估算只是过场，如何解决问题比结果更重要
- 学会估算，写下自己的假设，记得写单位
- 常见指标，[Peek] QPS, storage, cache, #server

## 术语
back-of-the-envelope，粗略的计算
carry out，实现
