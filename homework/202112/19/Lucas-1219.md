CH4 数据编码
- 向后兼容，backward compatibility，新代码读旧数据，对数据而言
- 向前兼容，forward compatibility，旧代码读新数据
- 缓存中使用数据结构，写入文件需要编码
- 编码 Encoding，序列化 Serialization，排列 Marshell
- JSON, XML, 二进制编码
- Thrift Facebook，接口定义语言，IDL，自带代码生成工具
- Thrift 两种不同的编码方式，BinaryProtocol，CompactProtocol
- TODO: BinaryProtocol 和 CompactProtocol的区别
- ProtocolBuffer 类似于 BinaryProtocol
- TODO-21: Thrift 不适合 Hadoop，因为采用Apache Avro，没有标签页码？