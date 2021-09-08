## 计算

- Question: how to save Table and execute SQL query on key-value model
- Data to store：table metadata，row, index
- Row: 行存和列存，行存 is faster to perform read and write, update, delete
- Index: 查询分为点查和range查询
- Select statement- 1) need to read one line of data fast, so every row should have an associated ID. 2) multiple lines might be read 3) read data by 点查和range

### 如何映射table -> key-value model

- 对每个表分配TableID, 每个索引分配 IndexID, 每一行分配RowID. TableID is unique in the union
- 每行数据

```
Key: tablePrefix{tableID}_recordPrefixSep{rowID}
Value: [col1, col2, col3, col4]
```

- Index数据

  - unique index
```
Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue
Value: rowID
```

  - non-unique index

```
Key: tablePrefix{tableID}_indexPrefixSep{indexID}_indexedColumnsValue_rowID
Value: rowID
```

### 元信息管理

- database/table unique ID could use as key in the format of `m_{tableID}`
- value is serialized metadata

### SQL query

e.g. `select count(*) from user where name="TiDB"`
- construct key range: all RowID are `[0, MaxInt64)` use this as range
- scan key range: read data
- filter data: for every row read, check if `name="TiDB`, return this row if true
- add all the data to count

flaws
- every row read costs RPC 
- not every row is needed and not needed to read
- the actual value is not of usage