# DDIA: pp 91-100



### Transaction Processing or Analytics

Comparison between *Transaction processing systems (OLTP)* and *Analytic system (OLAP)*

| Property               | OLTP - transaction                                | OLAP -analytic                            |
| ---------------------- | ------------------------------------------------- | ----------------------------------------- |
| Read Pattern           | Small number of records per query, fetched by key | Aggregate over lagre number of records    |
| Write Pattern          | Random-access, low latency writes from user input | Bulk import (ETL) or event stream         |
| Primarily used by      | End user/customer, via web application            | Internal analyst, for decision support    |
| Representation of data | Latest state of data (current point in time)      | History of events that happened over time |
| Dataset size           | Gigabytes to terabytes                            | Terabytes to petabytes                    |



Data warehouse: run the analytics

- read-only copy of data
- ETL (extract-transform-load): extract data from OLTP database, transform, clean, load to data warehouse

- Examples: 
  - Commercial: Teradata, Vertica, SAP HANA, ParAccel
  - Open-sourced: Apache Hive, Spark SQL, Cloudera Impala, Facebook Presto, Apache Tajo, Apache Drill. Some borrows ideas from Google's Dremel.
- 通常有很多column
- Data models: 
  - *Star schema* (*dimensional modeling*)
    - 中心是*fact table*: 记录transaction event等，包含 attributes 以及 foreign key references  to other tables (*dimension tables*)
    - Preferred because they are simpler
  - *snowflake schema* 
    - dimensions are further broken down into subdimensions.



### Column-Oriented Storage

- Observation: 一行数据有很多column，但是实际query的时候只需要其中几个entry而已
- Column-Oriented:  
  - don’t store all the values from one row together, but store all the values from each *column* together instead。读的时候再去各个column数据里取

- 优势：

  - 读入的data少了，节约带宽

  - 按列处理，chunk 小，合理运用CPU cycle

    

#### Column Compression

*bitmap encoding*

- 基于observation: 一列里distinct value相对较少

bitmap encoding 和one-hot vector差不多，把所有distinct 的值另外存起来，每行数据就只需要用0-1表示有没有某个值就可以了

- => 一些AND OR操作可以直接在 bit vector上做 (*vectorized processing*)，非常节约时间



### Sort Order

-  可以按insertion顺序 (append) 来，也可以sorted like SSTables and use that as an indexing mechanism。
- 如果要sort得整行都sort，可以按照attribute的使用频率来决定sort 依据的attribute
- sort了之后，做compression更容易，因为primary sort column上很可能一串数据都有相同的value

- Several different sort orders => 存多分数据，每个有不同的primary sort column



