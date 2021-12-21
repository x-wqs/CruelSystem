
## Schema evolution
Schema inevitably need to change over time. Thrift and ProBuf need to be backward and forward compatible. (with tag numbers.)

### Avro
Another binary encoding format. Use schema to specify the structure of data being encoded. Avro IDL is for human editing, JSON is for machine.No tag numbers.

The binary data can only be decoded correctly if the codeing reading the data is using the exact same schema as the code that wrote the data.

Reader's schema and writer's schema, if reads an unexpected field, ignore it. if an expected field is missing, fill with a default value.

Only add or remove a field that has a default value.

Use cases
- large file with lots of records, writer include writer's schema once at the beginning of the file.
- database with individually written reocrds, include a version number at the beginning of every record (kind of like MVCC from Spanner?) Reader read the version number then fetch the writer's schema from db to decode.
- sending records over network, Avro RPC.

Avro can do `dynamiclly generated` schema, e.g. dump a file from RDBMS using the db table schema. Regenerate the schema if db shema changes, fields are id'ed by name, so they can still have a partial match, while with Thrift or ProBuff, you'll neeed to manually update the mapping.

Why binary encoding?
- more compact, no field names.
- schema is requited for decoding, it has to be up to date to work properly. other format might be inconsistent.
- type checking at compile time.
- schema allows forward/backward compatibility checking.
