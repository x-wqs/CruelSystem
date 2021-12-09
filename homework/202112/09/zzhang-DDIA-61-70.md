# DDIA: pp 61-70



### The Foundation: Datalog

Similar to triple-store model:

- Write with `predicate(subject, object)`.

Write rules to match data recursively.

- Rules can be combined and reused in different queries.



### Summary

"NoSQL" database:

- **Document database**: target use cases where data comes in self-contained docu‐ ments and relationships between one document and another are rare.
- **Graph database**:  target use cases where anything is potentially related to everything.



Other not-mentioned different kinds of database in practice

- Genome data => sequence similarity search => specialized genome database software like *GenBank*.

- Particle physics projects like Large Hadron Collider (LHC).
- Full-text search



Next: 

trade-offs that come into play when *implementing* the data models



## Storage and Retrieval

Start with two families of storage engines

- *log-structured* storage engines
- *page-oriented* storage engines such as B-trees













