# RefinedCoding-D09-018    
- GFS P11 - P15    
- https://pdos.csail.mit.edu/6.824/papers/gfs.pdf    
    
## 6 测试      
    
      
### 6.1 微基准      
    
#### 6.1.1 读测试    
    
#### 6.1.2 写测试    
    
#### 6.1.3 记录追加    
      
### 6.2 现实的集群    
    
#### 6.2.1 存储    
    
#### 6.2.2 元数据    
    
#### 6.2.3 读写速率    
    
#### 6.2.4 集群主的负载    
    
#### 6.2.5 恢复时间    
    
    
### 6.3 负载详细分析    
    
#### 6.3.1 方法和警告    
    
#### 6.3.2 块服务器的负载    
    
#### 6.3.3 追加还是写入    
    
#### 6.3.4 集群主的工作量    
    
## 7 Experience    
    
## 8 相关工作    
    
## 9 结论    
    
    
      
## Terminology      
    
## Experience    
- GFS有哪些应用场景？作者就稍微提了一下    
    
## Terminology    
micro-benchmarks	,	微观基准    
inherent	,固有的；内在的；与生俱来的，遗传的    
cold cache result	,	冷缓存结果    
aggregate read rate	,	总读取率    
theoretical limit	,	理论极限    
saturated	,	饱和的    
observe	,	观察    
breakdown	,	分解    
caveats	,	注意事项    
efficiency	,	效率    
probability	,	可能性    
simultaneously	,	同时地    
aggregate write	,	汇总写入    
theoretical	,	理论的    
plateaus	,	高原    
culprit	,	罪魁祸首    
pipelining schema	,	流水线模式    
theoretical limit	,	理论极限    
concurrently	,	同时    
collision	,	碰撞    
involve	,	涉及    
in practise	,	在实践中    
occasional human intervention	,	偶尔的人为干预    
simultaneously	,	同时地    
virtually	,	几乎    
proportion	,	部分    
reclaimed	,	回收    
in aggregate	,	合计    
in addition	,	此外    
copy-on-write	,	写时复制    
somewhat	,	有些    
hobbled for a period	,	蹒跚了一段时间    
various	,	各种各样的    
burst of	,	一阵    
propagated	,	传播    
network topology	,	网络拓扑结构    
curve	,	曲线    
measured throughput	,	实测吞吐量    
error bar	,	错误栏    
confidence intervals	,	置信区间    
be illegible in 	,	难以辨认    
low variance	,	低方差    
much higher than	,	远高于    
in particular	,	特别是    
sustain	,	支持    
keep up with	,	跟上    
bottlebeck	,	瓶颈    
occasionally	,	偶尔    
sequentially	,	依序    
scan throughput, 扫描吞吐量 
binary search	,	   二进制搜索
name lookup cache	,	名称查询缓存    
be cloned to	,	克隆到    
provide leeway	,	提供余地    
reflect	,	反映    
carry out	,	执行    
heuristically	,	启发式地    
parallelism	,	并行性    
infer	,	推断    
differ	,	不同    
access pattern	,	访问模式    
stylize	,	程式化    
in the noise	,	在噪音中    
accurate data	,	准确的数据    
logistically impossible	,	逻辑上不可能    
cumbersome	,	笨重的；累赘的；难处理的    
overly generalize	,	过度概括    
tend to be tuned for something	,	倾向于为某事而调整    
conversely	,	反过来    
mutual influence	,	相互影响    
rather than	,	而不是    
more pronounced in	,	更明显的    
exhibit a bimodal distribution	,	呈现双峰分布    
seek intensive	,	寻求密集    
producer consumer queue	,	生产者消费者队列    
occasionally	,	偶尔    
outpaces	,	超越    
short-lived	,	短暂的    
significant portion	,	重要部分    
moreover	,	而且    
fairly	,	相当    
skewed	,	歪斜    
dominated	,	占主导地位    
breakdown, 击穿
differ	,	差异    
atempts	,	尝试    
proximates	,	近邻    
deliberately	,	故意地    
account for	,	占    
turn out	,	结果发现    
chunk server	,	块服务器    
significantly	,	显着地    
regularly	,	经常    
regenerated	,	再生的    
implicitly	,	隐含地    
from scratch	,	从头开始    
Unix open terminology	,	Unix开放术语    
automated	,	自动化的    
in contrast	,	相比之下    
concveived	,	对流    
over time	,	随着时间的推移    
evolved	,	进化的    
quota	,	配额    
rudimentary form	,	基本形式    
disciplined	,	纪律严明    
interfere with one another	,	互相干扰    
claim	,	宣称    
in fact	,	实际上    
reliably	,	可靠地    
occasionally	,	偶尔    
disagree about	,	不同意    
motivated	,	有动力    
corruption	,	腐败    
rather than	,	而不是    
modified portion,	修改部分    
checkpointing,   
synchronous write	,	同步写    
eventually migrated	,	最终迁移    
transient	,	短暂的    
sporadic hardware failure	,	偶发的硬件故障    
at the cost of	,	以…为代价    
transparently	,	透明地    
sophisticated	,	复杂的    
redundancy	,	冗余    
in contrast	,	相比之下    
opt forin paticular	,	选择特定    
sophisticated	,	复杂的    
write ahead log	,	提前写日志    
adapt a primary copy schema	,	调整主副本架构    
strong consistency guarantee	,	强大的一致性保证    
in terms of 	,	按照    
aggregate performance	,	综合表现    
POSIX compliant	,	符合POSIX    
unreliable components	,	不可靠的组件    
resemble architecture	,	类似于建筑    
network attached disk drive	,	网络连接的磁盘驱动器    
variable length object	,	可变长度对象    
persistent file	,	永久文件    
demonstrate	,	证明    
similar magntitude	,	类似的    
cost consciousness	,	成本意识    
reexamine	,	复审    
anticipate	,	预料    
radically	,	反复地    
conrurrently	,	相应地    
crucial	,	至关重要的    
motivate a novel	,	激发小说    
compensate	,	补偿    
high aggregate throughput	,	高总吞吐量    
a variety of	,	各种各样    
pass throughput	,	通过吞吐量    
delegate authority	,	委托权限    
dedicated	,	专用    
delegate	,	代表    
innovate	,	创新    
attack problem	,	攻击问题    
    
    
    
    
    
    
    
    
    
    
    