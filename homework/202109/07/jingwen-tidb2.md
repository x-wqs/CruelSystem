#### 三篇文章了解 TiDB 技术内幕 - 说计算

https://pingcap.com/zh/blog/tidb-internal-2

#### 需求：

1 一个table一般需要存储三种信息，1 元信息，2 表格中的行，3 支持引索的数据

2 因为主要业务是OLTP，需要快速读取修改删除一整行，所以对于row的信息选择了行存储（ps：列存储能更好地支持某一个维度上进行aggregation）

3 需要支持点查询和区间查询

4 需要支持primary index和secondary index

5 插入：需要将row写入kv，新建引索数据

6 更新：需要将row数据更新，更新引索数据

7 删除：需要将row删除，删除引索数据

8 读取一行数据：每个row有一个id

9 读取多行数据：

10 通过引索读取数据（点查询或者连续区间查询）





#### 实现：

在key的编码中加上tableID, indexPrefixID, indexedColumnsValue, rowID

查询语句的SQL运算：

1 构造key range（startKey，EndKey）

2 扫描Key Range，读取出数据

3 过滤数据（ where语句）

4 计算count

