阅读区+理解区
back to the envelope estimation
要记住一些基本的速度是多少
L1 cache reference 0.5ns 
mutex lock 100ns 
压缩算法 压缩1k bytes with Zippy需要10us 应该是很快的了
read 1MB sequentially from memory 250us 应该是指连续性的，比如图像有2m，我读取第一部分的1m
disk seek 10ms 好慢啊，比L1 cache和L2 cache慢很多倍，要尽量避免
数据包从加州发到荷兰一个来回要150ms，所以发送数据之前尽量压缩（而且压缩算法很快，1kb压缩只需要2us）
高可用的测量：一般用百分比
可用性的范围：大部分服务都是99-100之间
service level agreement (SLA)  一般都是跟接受服务方协商的up time，比如接收方需要我们的服务可用性在99.9%以上。
up time的测量：越多9越好，99.99999 > 99.9
up time测量大概：99%可用=每天可以用14.4 分钟



心得区
这一章光看可能没什么用，但可以当作一个reference来用，之后在看本书的其他章节中，如果遇到了类似的estimation的数字，可以翻回到这一张。
