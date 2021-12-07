# DDIA: pp 51-60

## Graph-like Data Models

#### Property Graph Model

Information storage:

- Vertex: id, out-edge, in-edge, properties (features)
- Edge: id, tail-vertex, head-vertex, edge-label, properties (features)

=> easily extend to accommodate changes in data structures



Cypher Query Language

- declarative query language for property graph for Neo4j

=> Compare with Graph queries in SQL

- In a graph query, the number of joins is not fixed in advance
  - In Cyper: use `() -[:WITHIN*0..]-> ()` rule: it may require traversing several upper layers to match the query. 
  - Use SQL (*recursive common table expressions*) to do similar task: require several joins and recursive traverse of the tables for the required data



#### Triple-Stores model and SPARQL

Information storage:

- *(subject, predicate, object)*

- Object can be two things
  - A value in a primitive datatype, such as a string or a number
  - Another vertex in the graph
- Require writing triples for vertex features as well. :(



Further application:

- RDF (Resource Description Framework) for *web of data*
  - Have not realized in practice
  - Subject, predicate, object: URIs



SPARQL query language:

- predates Cypher
- similar to Cypher



Graph Model vs Network Model (CODASYL)

- Graph Model:
  - Any vertex can have an edge to any other vertex
  - Can refer vertex through ID, can use index to find vertices with a particular value
  - Stored data is unordered => sort results when reading
  - Support high-level declarative query languages

- CODASYL:
  - Use a schema to restrict which record type can be nested within which other record type. ?
  - Read through access path
  - Ordered data => high overhead for insertion
  - Only imperative language supported

---

Next:

Datalog

- older than SPARQL and Cypher
- Cascalog is a Datalog implementation for querying large datasets in Hadoop.