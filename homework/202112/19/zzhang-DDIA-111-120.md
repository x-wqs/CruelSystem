# DDIA: pp 111- 120

## Encoding and Evolution

Evolvability: systems easy to adpt to change.



Requirement: 

- *Backward compatibility*: Newer code can read data written by older code.
- *Forward compatibility*: Older code can read data written by newer code.



Focus:

Formats of encoding data: JSON, XML, Protocol Buffers, Thrift, and Avro.

- Handle schema changes, support system with old&new data,code coexisting.

- Used ofr data storage and for communication: In web: REST, RPC, message-passing systems (actors, message queues).



### Formats for Encoding Data

Encoding: The translation from the in-memory representation to a byte sequence

Decoding: reverse



Language's built in encoding:

- bad
  - Tied to particular language
  - Security problem because decoding process needs to be able to instantiate arbitrary classes
  - No forward and backward compatibility
  - Not efficient



JSON, XML, CSV

- Ambiguilty: XML, CSV cannot distinguish between a number and a string; JSON cannot distinguish integer and floating-point numbers and doesn't specify the precision => Return twice in Twitter ID's case: JSON number and decimal string

- Binary string problem: JSON and XML don't support binary string.
  - Encode the binary data as text using Base64 => data size increases 33%

- CSV doesn't have schema => hard to interpret common or newline character inside one row.



Encode JSON directly to binary sequence

- need to include the key string, field type, length etc.
- only save a few bytes comparing to the original textual one.



### Thrift and Protocol Buffers

- Require a schema for any data that is encoded.



**Thrift**: 2 binary encoding formats

- *BinaryProtocol*: -> use field tag to replace the key name.
- *CompactProtocol*:-> pack the field type and tag number into a single byte. Use variable-length integers (-64~63, one byte; -8192~8191, two bytes, ...).



**Protocol Buffer**:

- Similar to Thrift's compactProtocol.



Schema evolution: schema need to change over time.

In Trift and protocol Buffer:

- Change field name doesn't affect encoding (they use field tag).

- Add new field: give a new tag number. => Cannot set this new field 'required'.
- Delete a field: Never use the same tag number again. Can only delete 'optinal' field.





















