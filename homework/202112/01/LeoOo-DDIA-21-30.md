# Data models vs query languages
Relational model vs document model
- data is organized into relations and each relation is an unordered collection of tuples.
- data is normalized, no duplicates are stored, only need to modify one place to change.
- mature, query engine and optimizers are highly optimized.
- ACID
- fit situations that require the most strict consistency level.
Document model
- data is stored as a document, e.g. in a JSON format.
- better locality, no need to query different table for related data.
- flexible, no fixed schema.
- better scalability, write efficient.
- because the check of data is performed when read.
- query optimizers are not as optimized as in RMDB
- not all SQL operators are supported.
- data is not normalized, there can be multiple copy of the same value, may need to modify several places to change.
- but does not fit in scenarios that requires strong consistency because certain constraints cannot be enforced like RMDB.

