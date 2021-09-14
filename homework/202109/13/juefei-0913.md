### 缓存
缓存可以暂时存储request和response

### 缓存的层次

一般在和database上面，先去查缓存，缓存没有再去database查

注意事项：
> 什么时候使用缓存
> 过期政策
> 一致性：和数据库的一致性
> 单点错误处理：用缓存的话，缓存就是一个单点错误，需要处理。
> Eviction策略：比如FIFO，LFU，LRU

### CDN
可以暂时把内容进行存储，然后有近的CDN可以用就直接在里面那从CDN取内容。

Consideration:
> cost
> setting an appropriate cache expriy
> CDN fallback
> Invalidating files

存储的内容：静态文件：JS CSS images

### 无状态服务层
服务一定要是无状态的，方便扩容

