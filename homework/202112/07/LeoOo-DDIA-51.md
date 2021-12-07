## Graph-like Data models
Document model fits well for one-to-many and relationships.

Graph model fits many to many relationships naturally.
- Social graph
- Roads or rail network
- Web graph

Graph model can also be used to store homogeneous data: objects with completely different types.

## Property graph (neo4j, etc)

Each vertex consists of
- A unique identifier
- A set of outgoing edges
- A set of incoming edges
- A collection of properties (k-v pairs)

Each edge contains:
- id
- u, v
- label to describe the relationship
- k-v pairs.

### The Cypher query language for property graph
- declarative

## Triple-stores and SPARQL
all information is stored in the form of very simple three-part statements: (subject, predicate, object). 

> I like this model better, it is quite simiple and elegant

the subject of a triple is equivalent to a vertex in a graph. The object can be
- a value in a primitive datatype
  - the predicate and object is a k,v pair represent a property
  - (lucy, age, 33)
- another vertex in the graph
  - predicate is an edge
  - (lucy, marriedTo, alain)

### SPARQL
