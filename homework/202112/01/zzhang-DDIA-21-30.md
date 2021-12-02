# DDIA: pp21-30

Maintainability

- Simplicity:
  - remove *accidental* complexity
    - *accidental*: not inherent in the problem that the software solves as seen by the users but arises only from the implementation
  - clear abstraction

- Evolvability

  - Agile: adpat to changes fast

    

Summary:

- functional requirements: what it should do, such as allowing data to be stored, retrieved, searched, and processed in various ways
- non-functional requirements: security, reliability, compliance, scalability, compatibility, maintainability, etc.



## Data Models and Query languages

Focus: data model (relational model, document model, graph-based model)



### Relational Model vs Document Model

SQL

Relational model:

- Data is organized into *relations* (*tables* in SQL)
- each relation is an unordered collection of *tuples* (*rows* in SQL)

- Good for *transaction processing* and *batch processing*



NoSQL: Not Only SQL

- greater scalability: very large datasets or very large write throughput
- specialized query operations
- dynamic



impedance mismatch

- Require a translation layer between the objects in the application code and the database model of tables, rows, and columns.

- Object-relational mapping (ORM) frameworks: ActiveRecord, Hibernate

  

