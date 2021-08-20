# Design Tiktok
- https://www.youtube.com/watch?v=Z-0g_aJL5Fw

# 假设
- 每天用户上传2个video，每个video根据h.264编码，5m
- SLA 99.999%，这个部分要求非常高了，全年停机时间少于6分钟
- 1M 用户

# 设计
## 高可用
## 低延迟
## 易扩展

# 接口
## 上传接口
- 使用s3存储video数据，使用RDB存储meta，这点存疑，应该使用kv存储meta会好一点，每个video meta里面包含video_link, video_id, create_id, create_time等等，这些信息不需要join其他信息，也不需要到RDB中取，kv接口会更快，适合缓存
## 信息流
- 使用uuid作为key缓存视频
- 后台加载热点视频
## 用户关系
- 喜欢和关注，使用外键关联
## CDN加速
## 多节点部署，防止单点故障
