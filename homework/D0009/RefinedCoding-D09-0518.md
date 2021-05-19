# RefinedCoding-D09-018
- GFS P1-P5
- https://pdos.csail.mit.edu/6.824/papers/gfs.pdf

## 1. Introduction


## 2. Design Overview 概要设计

### 2.1 Assumption 前提假设

### 2.2 Interface 接口

### 2.3 Architecture 架构

### 2.4 Single Master 单主节点
- 客户端程序发送文件名和内容位移之前，为啥还要翻译呢，翻译前后的内容有什么区别啊？谢谢了

### 2.5 Chunk Size 文件块大小
- Linux上一个大文件有可能在不连续的磁盘空间么？
- 延迟分配的话，会不会因为磁盘碎片太多，造成文件块频繁复制移动
- 你是说MapReduce的input.txt一般很大，比如100G,然后再GFS上被分成好多100MB的Chunk，
- 然后就可以用1000个节点来并行处理1000个Mapper job ?
- 现在有一个100G的大日志文件和1000个MapReduce的节点，怎么充分利用所有节点并行处理这个大文件呢
- MapReduce过程详解及其性能优化
- https://www.cnblogs.com/felixzh/p/8604188.html
- 使用mapReduce分析HDFS中大文件只起一个map的问题？
- https://www.aboutyun.com/thread-19021-1-1.html
- 总结：Mapper会去切分大文件

### 2.6 Metadata 元数据
- In-Memory Data Structures 内存数据接口
- Chunk Locations 文件块位置
- Operation Logs 行为日志

### 2.7 Consistency Model 一致性模型
- Guarantees by GFS / GFS约束
- Implication for Applications 应用蕴含

## 3. Sysetm Interfactions 系统交互

## 3.1 Leases & Mutation Order 租约和突变次序

## 3.2 Data Flow 数据流

## Terminology 术语
contiguous region	,	连续区域
Successive	,	连续
promptly	,	及时
a routine basis	,	例行基础
constantly	,	不断地
a modest number of	,	适量的
tolerate	,	容忍
alluded	,	暗示
Implication	,	含义
Mutation	,	突变
observations	,	观察
anticipated	,	预期的
reflect	,	反映
assumptions	,	假设
reexamined	,	重新检查
radically	,	根本
consists of	,	由组成
inexpensive commodity parts	,	廉价的商品零件
virtually guarantee	,	几乎保证
Therefore	,	所以
fault tolerance	,	容错
be integral to	,	不可或缺的
comprising billions of objects	,	包含数十亿个对象
unwieldy	,	笨重
mutated	,	变异的
may constitute large repositories	,	可能构成大型存储库
sequentially	,	依序
continuously	,	连续地
archival data	,	档案数据
intermediate results	,	中间结果
simultaneously	,	同时地
atomicity guarantees	,	原子性保证
loses its appeal	,	失去吸引力
consistency model	,	一致性模型
imposing an onerous burden	,	施加沉重的负担
append concurrently	,	同时追加
heavily accessed	,	大量访问
on a continuous basis	,	连续地
challenges 	,	挑战
opportunities	,	机会


