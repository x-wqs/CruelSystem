# Encoding and evolution

Formats for encoding data, JSON, Prot Buff, Thrift, AVro.

## Formats for encoding data
- in memory, objects, structs, list, array, etc. optimized to be manipulated by the CPU. (a pointer)
- on file or network, self-contained sequence of bytes, e.g. json, pointers won't help.
Need translation between the two scenarios, encoding and decoding. serialization and deserialization, marshalling and unmarshalling.

## Language specific
- Jave serializable, Python pickle, Java Kyro

## Language agnostic
- JSON, XML, CSV
- Binary encoding
  - avoid verbose encoding like json and xml, save space and network traffic.
  - Apache Thrift and ProtoBuf.
  - Code generation tool take a schema definition and produce class that implement the schema in different languages.
    - application then can use the code to en/decode.
  

## Schema evolution.
