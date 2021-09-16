Http请求可以打在任意的web service上，然后webservice可以共享database

然后，我们可以把缓存分离出来，缓存中存储session

NoSQL可以方便地AutoScale，也就是根据流量来调整里面的service的多少

我们也可以用两个Datacenter，然后DNS根据ip地址路由请求。

> 里面主要的问题就是数据同步：netflex有一个多datacenter的论文

Message Queue

Message queue一般接在Web servers后面


