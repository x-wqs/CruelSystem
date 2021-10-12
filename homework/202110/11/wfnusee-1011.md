# Rate Limiter
限流器：限制访问流量大小。
- 避免DoS攻击
- 减少成本
- 防止服务器过载

实现在 server-side 或者 gateway 都是可以的

## 常用算法
令牌桶
``` go
package ratelimit

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	rate   int64
	max    int64
	last   int64
	amount int64
	lock   sync.Mutex
}

// (now - last) * rate
func cur() int64 {
	return time.Now().Unix()
}

func New(rate int64, max int64) *RateLimiter {
	// TODO: 检查一下rate和max是否合法
	return &RateLimiter{
		rate:   rate,
		max:    max,
		last:   cur(),
		amount: max,
	}
}

func (rl *RateLimiter) Pass() bool {
	rl.lock.Lock()
	defer rl.lock.Unlock()

	// 当前桶中 是否还有令牌
	passed := cur() - rl.last
	fmt.Println("passed is: ", passed)

	amount := rl.amount + passed*rl.rate

	if amount > rl.max {
		amount = rl.max
	}

	if amount <= 0 {
		return false
	}

	amount--
	rl.amount = amount
	rl.last = cur()
	return true
}
```
并不需要维护真的令牌数据结构，用时间来计算即可。

漏桶
维护一个队列，队列满则丢弃。
一个问题在于，如果过去有一些没有完成的任务，最近的任务会全部被丢弃。而不是像令牌桶一样分散到不同的时间里。

滑动窗口相关算法：
1. 固定滑窗：无法处理峰值。
2. 滑窗日志：需要耗费大量的内存。
3. 滑窗计数：假设请求分布均匀。
