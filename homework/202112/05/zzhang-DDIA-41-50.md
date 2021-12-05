# DDIA: pp 41-50



### Relational Versus Document Databases Today

#### Data locality for queries

Document based database

- document stored as a single continuous string

- fetch the whole document when needed

  

Locality in

- Bigtable data model: column-family concept

- Oracle: multi-table index cluster tables
- Spanner: a table's rows should be interleaved within a parent table. ?



### Query Languages for Data

**Declarative**: SQL, relational algebra, CSS, XSL

- pecify the pattern of the data you want
- database system’s query optimizer decides which indexes and which join methods to use, and in which order to execute various parts of the query.



**Imperative**: IMS and CODASYL

- tells the computer to perform certain operations in a certain order



### MapReduce Querying

Neither declarative nor imperative.

- map, reduce

- Mongoldb example:

  ```javascript
  db.observations.mapReduce(
    	function map(){
        // handle and generate key and value
        emit(keystuff, valuestuff);
      }, 
      function reduce(key,values){
        return Array.sum(values);
      }, 
      {query: {
        keyname: "value", 
        out: "outputname"}})
  ```
- Read-only

- Higher-level query languages like SQL can be implemented as a pipeline of MapReduce operations



MongoDB 2.2 support declarative query language: **aggregation pipeline**

```javascript
db.observations.aggregate([
  { $match: { key: "value"} }, // search data
  { $group: {
    	_id: { // key
        year: { $year: "$obstimestamp" },
    		month: { $month: "$obstimestamp" }
      },
   		total: { $sum: "$num" } // value
  } }
]);
```



## Graph-Like Data Models

appropriate when many-to-many relationships are very common in data



two objects:

- Vertices (nodes, entities)
- Edges (relationships, arcs)



Can support heterogenous vertices and edges data: different meaning, different feature dimension in one graph.



Next focus:

- Property graph model: Neo4j, Titan, InfiniteGraph

- Triple-store model: Datomic, AllegroGraph

Three declarative query language:

- Cypher
- SPARQL
- Datalog

