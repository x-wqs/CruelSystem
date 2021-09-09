# TiDB 计算

## 需求
- 支持 Primary Index，Secondary Index
- select, 查询方式：点查和range查询
- Insert，将 Row 写入 KV，并且建立好索引数据
- Update，将 Row 更新的同时，更新索引数据（如果有必要）
- Delete，在删除 Row 的同时，将索引也删除

## 方案
- 每个表分配一个 TableID，集群唯一
- 每一个索引都会分配一个 IndexID，表内唯一
- 每一行分配一个 RowID（如果表有整数型的 Primary Key，那么会用 Primary Key 的值当做 RowID）， 表内唯一
- ID 都是 int64 类型
- 每行数据编码为
  - Key: TablePrefix{TableID}_RecordDPrefixSep{RowID}
  - Value: [Col1, Col2, Col3, Col4]
- Unique Index编码为
  - Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue
  - Value: rowID
- 非Unique Index编码为
  - Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue_rowID
  - Value: null

## 元信息管理
- Database/Table都被分配了一个唯一的ID，在编码为Key-Value时，这个ID都会编码到Key中，加上m_前缀
- 有一个专门的 Key-Value 存储当前 Schema 信息的版本

## SQL on KV 架构
- TiKV Cluster 主要作用是作为 KV 引擎存储数据

## SQL 运算
- Select count(*) from user where name="TiDB"
  - 构造key range：根据Row key 编码原则，构造出一个查询区间[start, end]
  - 扫描key range，根据上面的查询区间，读取TiKV中的数据
  - 过滤数据，对于读到的每一行数据，计算 name="TiDB"这个表达式，如果为真，则向上返回这一行，否则丢弃这一行数据
  - 计算 Count：对符合要求的每一行，累计到 Count 值上面
- 缺点
  - 每一行读取都需要RPC开销
  - 不需要把每一行读出来

## 分布式 SQL 运算
- 将计算尽量靠近存储节点，以避免大量的 RPC 调用
- 将Filter下推到存储节点进行计算，只需要返回有效的行，避免无意义的网络传输
- 将聚合函数、GroupBy也下推到存储节点，进行预聚合，每个节点只需要返回一个Count值即可
- 再由 tidb-server 将 Count 值 Sum 起来

## SQL 层架构
- SQL 请求直接或者通过 Load Balancer 发送到 tidb-server
- tidb-server 解析 MySQL Protocol Packet
- tidb-server 做语法解析、查询计划制定和优化、执行查询计划获取和处理数据
- 数据全部存储在 TiKV 集群中
- tidb-server 需要将查询结果返回给用户
