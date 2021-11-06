# What's NewsFeed
类似于FaceBook的主页 或者微博的主页.
可以展示出用户的好友动态, 以及ta关注的groups的动态.
动态按照时间顺序, 展示最新的前XXX条. 动态包含文本,视频,图片,音乐等等
用户可以发布News, 以及获取自己的Feeds Timeline

# Requirements
## Functional
1. 支持用户发布动态
2. 支持用户获取自己的Feed主页,包含所关注friends/groups的最细动态
3. 动态包含文字,图片,视频等多媒体内容
4. 对于Active User(正在使用webpage/app浏览), 支持将最新的post发送到他们的newsFeed上(**这里不太理解**)


## Non-Functional
1. 最大生成newsFeed延迟 <> 2s
2. 新post变为可见的延迟 < 5s

# Capacity & Constraints
## User
每个用户300个friends, 关注200个pages
## Traffic
300M DAU & 5 Fetches/user --> 1500M fetch reqs/day --> 17000 fetch req/s (0.1M or 0.09M second/day)
## Storage
这里主要考虑In-memory的post, 也可以关注一下persistent storage的量级

1 KB/post& 500 posts in mem --> 500KB in mem per user

500Kb * 300M --> 150KB * 1e9 = 150 TB mem usage

100 GB/server -> 1.5 * 1e3 = 1500 servers

# APIs
完成了前面粗略的系统规模的估计,这里再看具体如何设计API

这里只考虑获取feed的API, **发布post的API怎么设计比较好呢?**

> 输入: getUserFeed(api_dev_key, user_id, since_id, count, max_id, exclude_replies)
- api_dev_key, 如果有API developer请求接口的话, 需要使用这个来进行标识, 从而进行管控和治理
- user_id
- from_id, count, max_id 这里用时间戳会不会更好一点
- exclude_replies **这里感觉没必要加这个? 比如用户屏蔽的内容,可以在application server侧去掉**

> 输出: json, 包含feed的items 

**怎么设计items比较好?**
- itemID
- timestamp
- source 表示是由谁发布的post
- itemType 表示item的类型, 如text, image, video
- resourceLinks 如果包含image/video之类的内容, 记录其cdn的链接
  

# Database Design
数据库的设计上围绕系统中有哪些对象. 比如我们支持user/pages/group, 可以把pages/groups抽象出来作为Entity, user单独作为User, 以及发布的post作为PostItem
## User
## Entity
## PostItem
对于Item, 记录text or media, 一个Feed可能包含多个media? 多张图片
## FeedItem
- item id
- user id
- visibility
- content
- date
- num_likes

对于 num_likes 感觉可以单独存储?

## Media
记录每一个media的具体内容, 包括 
- media id 
- URL
- date
- type
- description


## FeedMedia 
记录Feed中包含了哪些Media(ID), 有可能包含多个? 然后进行join操作
- item id
- media id

此外,对于用户间的关注关系也需要有数据结构进行记录
## UserFollow


To Be Cont.
