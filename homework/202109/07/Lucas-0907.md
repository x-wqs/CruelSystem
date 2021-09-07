# 三篇文章了解 TiDB 技术内幕 - 说计算
- https://pingcap.com/zh/blog/tidb-internal-2

## 需求
- 利用键值结构存储数据库表
- 表结构三部分：原信息，行信息，索引数据
- 每行记录，可以按行存储，可以按列存储
- OLTP更关注数据的写入，所以采用行存储
- 支持主索引，也支持二级索引
- 支持点查，比如 select name from users where id = 1;
- 支持范围查询，比如select name from user where age > 20 and age <30;
- 通过二级索引indexAgent查询年龄范围
- 需要支持 Unique Index和非Unique Index 两种索引

## 实现
- 键值数据库为每个表分配一个TableID，集群内唯一
- 每个索引分配一个IndexID，表内唯一
- 每一行分配一个RowID，整数型的主索引可以作为RowID，表内唯一
- 所有ID都是Int64
- 每行数据编码为
- Key: TablePrefix{TableID}_RecordDPrefixSep{RowID}
- Value: [Col1, Col2, Col3, Col4]
- 每个Index编码为
- Key: TablePrefix{TableID}_IndexDPrefixSep{IndexID}_IndexedColumnValue
- Key: RowID
- Unique Index 使用上述编码
- Non-Unique Index 可能对应多行的ColumnValue
- Key: TablePrefix{TableID}_IndexDPrefixSep{IndexID}_IndexedColumn_RowID
- Value: null
