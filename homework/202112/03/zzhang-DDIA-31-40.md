# DDIA: pp31-40

Document-oriented databases: 

MongoDB, RethinkDB, CouchDB, Espresso

- better locality



### Many-to-One and Many-to-Many Relationships

*Normalization*: remove duplicating values that could be stored in just one place.

- Example: record city ID instead of city name: easy to update

- Many-to-one relationship: Above city ID example. Hard to 'join' in document based db => ship to application code

- One-to-many: tree structure, one person one document for all information

- Many-to-Many: One person one document with reference of city, organization, recommendation reference etc.

  

#### Handle 'join' problem:

Two most prominent solution: relational model (popular SQL), network model (obscurity)

Network model: 

- Fetch records follow an access path (chain of links).
- Changing and maintaining access path is complex. 

Relational model

- A relation is simply a collection of tuples (rows)
- query optimizer

Comparison to document db

- document db follows similar route of relation model. Use *document reference* similar to *foreign key*



### Relational vs Document DB

Document DB:

- schema flexible (*schema-on-read*): good for those data with heterogenou or dynamic structure
- better performance due to locality
- similar application data structure (resume)

Relational DB:

- better support for joins, and many-to-one and many-to-many relationships



Highly interconnected data => use graph model





