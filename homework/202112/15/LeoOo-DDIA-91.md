

## Data warehousing
Database used mainly for analytic workloads instead of transactional.

Business analysts are not supposed to run on transaction dbs because they might slow the system down.

Data warehousing is a readonly copy and for OLAP, generally is ETLed transaction db.
It can be optimized for analytical access patterns.

Star-schema: fact-table and dimension tables.
  - snowflake schema, further breakdown the dimensional table.

## Columnar storage.
Row-oriented: all values in the same row are stored next to each other.

Even if you are only interested in one column, the system has to load the entire row from the disk.

Column db: store each column together: only load the columns you need.
- Columns can be futher compressed, e.g. run-length encoding, bitmap.

Vectorized processing
- compress column data to fit CPU's L1 cache, certain operators like AND/OR can be applied to such compressed chunks for very fast execution.

### Sort order in column storage
