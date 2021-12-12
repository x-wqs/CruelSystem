## 图数据模型

如果数据大多是一对多关系或没有关系，文档模型最合适。

关系模型能够处理简单的多对多关系。随着数据之间关联越来越复杂，将数据建模转换为图模型会更加自然。例如：社交网络、Web图、公路网。算法：导航搜索道路网中最短路径。PageRank计算Web图上网页流行度，从而确定搜索排名。

## 属性图

每个顶点包括：唯一标识符、出边的集合、入编的集合、属性的集合。

每个边包括：唯一标识符、边开始的顶点、边结束的顶点、描述两个顶点间关系类型的标签、属性的集合。

图存储看成两个关系表组成，一个用于顶点，一个用于边。

## Cypher查询语言

用于属性图的声明式查询语言。最早为Neo4j图形数据库创建。

## SQL中的图查询

join操作数量并不是预先确定。

## 三元存储与SPARQL

三元存储等同于属性图模型。三部分：主体、谓语、课题。主体相当于顶点，课题是以下两种之一：原始数据类型的值如字符串或数字，图中另一个顶点。

语义网：资源描述框架RDF让不同网站以一致格式发布数据。来自不同网站的数据合并成数据网络。一种互联网级别包含所有数据的数据库。