# DDIA: pp 121 - 130



### Thrift and Protocol Buffers

Change the datatype of data:

- Okay, but lose precision (32-bit int to 64-bit int, old code still reads as 32-bit int)



List:

- Protocol Buffers: Use marker `repeated`.  Each entry has same field tag inside the list. => New code read old data as a list of zero or one elements; old code read new data's last element.

- Thrift: has dedicated list datatype.



### Avro

- binary encoding format.
- No tag number, nothing to identify fields or their datatypes.
- Parse the binary data: use the order in the schema.

*Writer's schema*: Application encode the data using any schema it knows.

*Reader's schema*: the schema the application code is relying on. code may have been generated from that schema during the application’s build process.



*Writer's schema* and *Reader's schema* don't have to be the same.

- Avro library will resolves the difference.
- Okay for having different order of fields. 
- Writer has one field that reader doesn't read: ignore
- Reader read a field that writer doesn't write: use default value



Schema Evolution:

- New writer schema, Old read schema
- Requirement: Only add/remove fields with default values.



Default values: in Avro, if you want to allow a field to be *null*, you have to use a *union type*, e.g., `union {null, long, string} field`.



No *optinal*, *required* marker.



Change the datatype, field name, add a branch to a union type: backward compatible but not forward compatible.



Writer's schema

- Good if millions of records with the same schema.

- Reader can fetch the writer schema with version number if possible.
- Avro RPC protocol: reader and writer negotiate the schema version through network.



Dynamically generated schemas:

- In Avro: just generate a new Avro schema. Data export process don't care.
- In Thrift or Protocol Buffers: requires manually setting the field tag.



Avro is friendly to dynamically typed programming languages (JS, python, ruby).



### Summary of binary encodings

**Advantage** of binary encoding based on schemas:

1. more compact: they can omit the field names
2.  the schema is required for decoding, you can be sure that it is up to date 

3. Keeping a database of schemas allows you to check forward and backward com‐ patibility of schema changes
4. For statically typed pl users, the ability to generate code from the schema is useful, since it does type checking at compile time.



## Models of Dataflow

Cover:

Most common ways how data flows between processes: Databases, Service calls (REST, RPC), async message passing



### Through Databases

Encode: write data to the db

Decode: read data from the db



Old code might read the new data (with new fields), and write back with old fields => lost new fields' data.



*data outlives code*: in databse, the five-year-old data will still be there, in the original encoding, unless you have explicitly rewritten it since then.

*Migrating* data to a new schema is expensive.

- LinkedIn's document database Espresso use Avro for storage
- Fill the new fields of old data with *null*.

