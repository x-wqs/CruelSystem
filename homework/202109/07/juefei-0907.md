### TiDB

insert: 将Row写入KV，并建立索引
update: 更新Row的同时，更新数据索引
delete: 删除Row的同时，删除Row的同时，删除数据索引

TiDB引擎：全局有序

表：TableID 索引：IndexID 每一行：RowID
TableID集群唯一，IndexID/RowID在表内唯一

数据存储模式：
```
Key: tablePrefix{tableID}_recordPrefixSep{rowID}
Value: [col1, col2, col3, col4]
```

索引存储模式:
```
Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue
Value: rowID
```
其中xxPrefix都是字符串常量：

```
var(
    tablePrefix     = []byte{'t'}
    recordPrefixSep = []byte("_r")
    indexPrefixSep  = []byte("_i")
)
```

### SQL运算

将SQL查询映射为对KV的查询，再通过KV接口获取对应的数据，最后执行各种计算。
