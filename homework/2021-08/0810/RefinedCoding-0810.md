# Beyond REST - 1·Rapid Development with GraphQL Microservices
- https://netflixtechblog.com/beyond-rest-1b76f7c20ef6

## 架构
- Docker 容器里面运行，便于发布和获取
- 一个简单的 NodeJS 网页服务器和 Graphile 库
- TODO：无需编写更多的REST接口，而是在GraphQL里嵌入数据库查询
- 除此之外，还有奈飞内部的组件，包括安全，日志，指标和监控
- 提供更好的 REST 或者 REST ++ 服务

## 特点
- 用数据库视图替代 API 接口层
- 允许使用现有的 GraphQL 模式来修改数据库
- 利用 PostgreSQL 复杂类型，操作 PostgreSQL 聚合功能
- 
