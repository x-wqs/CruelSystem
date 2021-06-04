# 20210511: Map Reduce P11 - P13  
    
## 6.1 Large Scale Indexing    
- a complete rewrite of indexing system for Google web search service  
- when was the GFS paper published?  
- simple code dropped from 3000 C++ lines to 700 MR lines  
- good performance for one change from a few months to a few days  
- easy operation  
  
## 7. Related Work  
- parallel prefix computation of N element array in LogN time on N Processor  
- how does fault tolerant work in Map Reduce?  
- bulk synchronous programmnig & MPI primitives  
- locality optimization inspiread from active disk techniques  
- backup task similar to the eager scheduling mechanism  
- sorting facility similar to NOW-Sort operation  
- mappers patition data to be sorted and send to reducer  
- River system no better than backup tasks  
- BAD-FS programming model for wide area network  
- TACC relies on re-execution for fault tolerance  
  
## 8.1 Conclusions  
- easy to use fo programmers without parallel/distributed system experance  
- hide parallelization, load balancing, fault tolerance & locality optimization  
- used by Google indexing, sorting, data mining & machine learning  
- implementation by thousands of machines  
  
## 8.2 Lessons  
- restircting programming model make parallelize/distribute computation and fault tolerant esay to use  
- locality optimization for scarce network  
- writing intermediate data to local disk to save network bandwidth  
- redundant execution helps slow machines, machine failures and data loss.  
  
## 8.3 References  
- https://www.youtube.com/watch?v=Rz8JCS9TfOQ  
- when does N equal to NP? N = 0 or P = 1  
- spend two nights to optimize light speed as light speed too slow  
  
## Terminology 术语    
significant, 重大  
to date, 迄今为止  
intervention, 干涉  
hiccups, 打  
dealt with, 处理  
associative function, 联想功能  
simplification, 简化  
distillation, 蒸馏  
primitives, 原语  
exploits, 漏洞利用  
transparent, 透明  
locality optimization, 局部性优化  
draws, 抽签  
inspiration, 灵感  
commodity processors, 商品加工商  
shortcomings, 缺点  
spirit, 精神  
uniformities, 均匀性  
non-uniformities, 不均匀性  
heterogeneous, 异质  
heterogeneous hardware, 异构硬件  
system perturbations, 系统扰动  
achieves, 达到  
finegrained tasks, 细粒度的任务  
redundant executions, 多余的执行  
fundamental similarities, 基本相似之处  
locality-aware scheduling, 位置感知调度  
congested network links, 拥塞的网络链接  
scarce, 稀缺  
redundant, 多余的  
  
  
## Notes    
  
    
    
    
    
    