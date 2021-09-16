#### cache

好处：提高系统效率和速度，支持多语言

作用：如果数据在缓存里，则可以直接给用户。如果数据不在缓存里，则从数据库拿，然后存在缓存里。

论文：6 数据库缓存能提高多少的研究todo

用的时候的一些考虑：

1 读多些少的时候

2 使用过期政策

3 数据一致性，如果对缓存和数据库的修改不是在同一个事务里的话，有可能造成数据不一致（论文：scaling memcache at facebook）

4 single point of failure 有可能发生在cache系统里，建议使用多个（那会不会更多不一致性？）

5 eviction 政策：LRU和LFU（啥时候用LRU啥时候用LFU？to do）



#### CDN

存储一些图像， 视频，CSS，Javascript

当用户需要的时候再发送给他

Dynamic content caching



