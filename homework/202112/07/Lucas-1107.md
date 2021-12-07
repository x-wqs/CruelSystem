## 文档模型比较
- 无模式 schemaless，隐式模式，数据库不强制执行，读时模式，schema-on-read
- 写时模式，scheme-on-write，静态编译，类型检查，数据迁移 migration
- MySQL alert table 复制整个表，一段时间宕机，可以设置为null，读取时填充
- 关系数据库提供文档支持

## 查询数据的局部性
- Google Spaner，允许模式声明一个 表的行交错嵌套在父表内
- Oracle 多表索引集群表，multi-table index cluster tables
- Bigtable, Cassandra, HBase 列簇 columu family

## 查询语言
- 声明式，SQL，关系代数，不在乎实现，查询优化器决定使用索引和链接方法，隐藏数据库引擎，适合并行执行，CSS，XSL
- 命令式，IMS，COSASYL，特性顺序执行，JavaScript，DOM 文档对象模型

## MapReduce
- MongoDB, CouchDB 有限支持，介于命令式和声明式之间
