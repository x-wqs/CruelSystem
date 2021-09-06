# Design Youtube
## 格式选择
- webp 小文件压缩率高
- mp4 h.264，稍微大点的文件
## 存储选择
- s3
- fastdfs 分布式小文件存储系统
- Tectonic facebook
## 视频上传
- youtube不同于tiktok有很多类似相同的文件，实现文件秒传，可以计算md5或者使用merkle tree hash
- 为了保证高可用可以使用分布式存储系统
## 视频观看
- CDN加速
- 冷热缓存，热点视频通过后台计算推到CDN节点
## 用户关系
- 关系型数据库存储
- 大V用户使用cache加速获取
- 推送关注用户更新，采用pull的形式+后台计算push结合
## 分享视频
- 使用短链接方式分享，base32计算短链接
## like/dislike
- 使用MQ发送用户操作，后台计算，推荐视频
- like个数在缓存中计数
## comments
- 通常存储在关系型数据库中，并缓存最热评论到video meta中
