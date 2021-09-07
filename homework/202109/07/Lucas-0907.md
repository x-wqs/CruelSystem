# 三篇文章了解 TiDB 技术内幕 - 说计算
- https://pingcap.com/zh/blog/tidb-internal-2

## 利用键值结构存储数据库表
- 表结构三部分：原信息，行信息，索引数据
- 每行记录，可以按行存储，可以按列存储
- OLTP更关注数据的写入，所以采用行存储
- 支持主索引，也支持二级索引
- 支持点查，比如 select name from users where id = 1;
- 支持范围查询，比如select name from user where age > 20 and age <30;
- 通过二级索引indexAgent查询年龄范围
- 需要支持 Unique Index和非Unique Index 两种索引
