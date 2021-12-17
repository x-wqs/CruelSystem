# DDIA: pp 101-110



Materialized Aggregates:

- Caches some of the counts or sums that queries use most often 

=> Materialized view: actual copy of the query results, written to disk.

- good for read-heavy scenario

- expensive for writing

*Data cube* (*OLAP cube*): special case of a materialized view

- Precomputed sum/counts of one dimension
- Not much flexibility



## Summary

- Comparison between *transaction processing* (PLTP) storage engines and *analytics* processing (OLAP) storage engines.



OLTP:

- Log-structured: lsm-trees, sstables, etc
  - Turn random-access writes into sequential writes on disk => higher write throughput
- Update-in-place: b-trees



OLAP:

- Column-oriented storage
- Precomputation

