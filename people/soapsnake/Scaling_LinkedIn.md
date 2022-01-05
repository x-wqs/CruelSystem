A Brief History of Scaling LinkedIn 读书笔记
指标:3亿5千万用户(到2015年3月)
系统最开始是一个叫Leo的大application,随着用户增长,大系统开始拆分,拆分的子系统利用RPC进行调用.
接着是DB的读写分离,利用databus,随着用户量进一步的增长,系统逐渐演变到SOA架构
系统引入缓存进行加速,并且开发了kafka作为集成总线,后面又开发了rest.li并且构建了全球化的数据中心.

