# MapReduce Lab QA

## 类似设计
- 同步请求
- 用休眠代替条件变量
- 使用通道

## 设计终点
- 协调器不能做太多工作
- 不能发送冗余请求
- 条件打印DPrintf in raft/utils.go
- 颜色打印
- 重定向到文件

## MapReduce问题
- Combiner in Map
- Sorting in Reduce
- 中间文件为什么存在本地而不在GFS
- 协调器如何退出，使用Exit任务
- 失败任务为什么使用超时重试，而不是论文中的备份任务？
- RPC调用使用指针而不是对象？RPC如何传输地址？

## 实现RPC数据结构

## 协调器的RPC处理

## Worker的循环去获取任务和处理返回

## 实现Map任务

## 实现Reduce任务

## 协调器循环去处理请求和分配任务

