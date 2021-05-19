# RefinedCoding-D09-018
- GFS P11 - P15
- https://pdos.csail.mit.edu/6.824/papers/gfs.pdf

## 1. Introduction


## 2. Design
### 2.1 Assumption
### 2.2 Interface
### 2.3 Architecture
### 2.4 Single Master
- 客户端程序发送文件名和内容位移之前，为啥还要翻译呢，翻译前后的内容有什么区别啊？谢谢了

### 2.5 Chunk Size
- Linux上一个大文件有可能在不连续的磁盘空间么？
- 延迟分配的话，会不会因为磁盘碎片太多，造成文件块频繁复制移动啊

### 2.6 Metadata
- In-Memory Data Structures
- Chunk Locations
- Operation Logs

### 2.7 Consistency Model
- Guarantees by GFS
- Implication for Applications

## 3. Sysetm Interfactions
## 3.1 Leases & Mutation Order
## 3.2 Data Flow

## Experience
- GFS有哪些应用场景？作者就稍微题了一下

## Terminology
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




