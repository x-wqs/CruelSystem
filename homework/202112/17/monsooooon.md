ddia ch3 part2
# oltp or olap
- oltp: commercial transaction, frequent crud op and low latency
  - txn: no relationship with ACID txn
  - lookup small number of records
  - interactive app for end users
- olap: data analytics
  - large range row scan
  - use aggregate funcs to calc statistics (avg, count, max, min), not raw data
  - used for business intelligence / analytics
  - no clear cut with oltp
  - database -> data warehouse / data lake
- ETL(Extract-Transform-Load)
  - from oltp to olap (read-only)
  - continuous streams of updates / periodic dump

# olap engine overview
- use sql and relational models (mostly)
- has graph vis tools for analytic results
- open-source: Hive, Spark SQL, Presto, Google's Dremel

# olap schema
- stars
  - fact table: each row -> an event that occurred at a particular time
  - dimension table: has foreign key with fact table: who, when, where, why, how, more info
  - the fact table is in the middle, surrounded by its dimension tables; the connections to these tables are like the rays of a star.
- snowflakes
  - more tables, break diamension table into more
  - more normalized, less space, higher query cost
  - less used in olap
- VERY wide tables in OLAP: 100+ columns

# column oriented storage
- treat columns as store unit, not rows
- useful for wide tables
- columns can be compressed to save space & indexed to accelerate queries
-
