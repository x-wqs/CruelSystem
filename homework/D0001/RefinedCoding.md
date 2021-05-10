# 20210510

## Map Reduce P1 - P5
- 应用场景
  - crawled documents
  - web request logs
  - inverted index
  - graph structure of web documents
  - stat page access per host or day
  - distributed grep
  - reverse web link graph
  - term vector per host, 主机术语向量？
  - distributed sort
- Example: Count Word Frequency  
  ```java
  Map<String, Integer> merge(String docName, String content) {  
    for word in content {  
      ans.add(word, ans.getOrDefault(word) + 1)  
    }  
  }  
  Map<String, Integer> reduce(String word, List<Integer> counts) {  
    return sum(counts);  
  }  
  ```

## Workflow
  - master: total / 64MB blocks = splits
  - job consists of a set of tasks
  - M splits
  - R pieces, hash(key) mod R
  
## Questions
1. In 2.1, reduce函数里面String key参数做什么用的？
2. In 3.1.4, The locations of these buffered pairs on
the local disk are passed back to the master, who
is responsible for forwarding these locations to the
reduce workers. Can master and reducer read the location on the local disk of mapper?
3. In 3.1.5, The sorting is needed because typically
many different keys map to the same reduce task. If
the amount of intermediate data is too large to fit in
memory, an external sort is used. Does only one reducer or all of reducers sort intermediate data? Give an example or try in lab.
4. In 3.1, The sorting is needed because typically
many different keys map to the same reduce task. If
the amount of intermediate data is too large to fit in
memory, an external sort is used. How to show total frequency for one word?
5. In 3.3, Completed map tasks are re-executed on a failure because their output is stored on the local disk(s) of the
failed machine and is therefore inaccessible. How does reducer read the intermediate data from the local disk on mapper?
6. In 3.4, shall we split the original input to M splits or by 64MB blocks? 
A: M splits, 64MB is for GFS storage.
7. In 3.5, Why does master keeps O(M * R) states in memory?


## Terminology 术语
commodity,商品,廉价的？商用机？  
conceptually,概念上   
hundreds of thousands of,
conspire,共谋,协力  
obscure,昏暗的,朦胧的,晦涩的,不清楚的,隐蔽的,不著名的,无名的  
obscure,使…模糊不清,掩盖,隐藏,使难理解  
reaction,反应,反作用  
fault-tolerance,容错系统  
tolerance,宽容,容忍,耐力,公差  
inspired by,受启发  
primitive,原始的,远古的,简单的,粗糙的  
primitive,原始人  
intermediate,中间的,过渡的,中级的,中等的  
derived,衍生的  
appropriately,适当地  
tailored,量身定制  
refinements,细化  
invocation,调用  
emit,发射  
drawnfrom,取自  
term vector ,术语向量  
partitioning facilities,隔断设施  
considerably,相当  
bisection,二等分  
unreliable,不可靠  
periodically,定期地  
conduit through,穿过  
propagated from,从传播  
tolerate machine failures,容忍机器故障  
inaccessible,无法进入  
eventually,最终  
checkpoints,检查点  
Semantics,语义学  
deterministic,确定性的  
sequential,顺序的  
rely on ,依靠  
atomic commits,原子提交  
underlying,潜在的  
vast majority,绝大多数  
relatively,相对地  
Locality,地区性  
particular,特定  
weaker semantics,较弱的语义  
conserve,养护  
replica,复制品  
significant fraction,很大一部分  
Task Granularity,任务粒度  
phase,阶段  
  
## Notes  
Q: git-diff to ignore ^M  
A: git config --global core.autocrlf true  
  
Q: Add User to WSL Linux Distro in Windows 10  
A: https://winaero.com/add-user-wsl-linux-distro-windows-10/  




