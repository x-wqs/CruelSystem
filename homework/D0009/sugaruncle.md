# Google File System P1-5 notes

## Design Space

1. Component failures are the morm rather than the exception.
2. Files are huge by traditional standards. Multi-GB files are commmon.
3. Most files are mutated by appending new data rather than overwriting existing data.
4. Co-designing the applications and the file system API benifits the overall system by increasing our flexibility.
用中文总结一下就是，机器是普通机器，很容易出问题；文件会很大；以“附加”代替“覆写”，最后，面向应用设计系统，使得设计更加灵活。

## Design Assumptions

1. The System is built from many in expensive commodity components that may fail.
   * 这个和上面第一条相同，所以处理错误的能力对GFS非常重要。
2. The System stores a modest number of larger files.
   * 百万级的100MB文件，适量的Multi-GB文件，（我对文件大小其实没什么概念）
3. large streaming reads and small random reads。
    * [流式读取](https://groups.google.com/g/cs501pku/c/rVblGSxAUWE?pli=1)貌似对应的是顺序读取，与随机读取相对应。
4. The workloads also have many large sequential writes that append data to files.
5. The system must efficiently implement well-defined semantics for multiple clients that concurrently append to the same file.
6. Hight sustained bandwidth is more import tant than low latency.
   * 高带宽比低延迟重要，不懂，可以反应慢一些，但一旦建立连接就要非常稳定？

## Interface

* create 
* delete
* open
* close
* read
* write

## Architecture

参照图一，一个master 用于管理数据调度，很多chunkservers 用于实际存储数据。
后面有提到，不会再master 上运行chunkserver 上的任务，避免master成为瓶颈。

## Consistency Model

* Guarantees by GFS / GFS
* Implication for Applications
