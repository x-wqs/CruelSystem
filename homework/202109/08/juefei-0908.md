### 数据类型

int家族：int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr

### 切片
没啥好讲的，和python差不多

### 类型转化

用括号括起来，比如int(xxx)

### 循环
只有for

### 方法
方法取代了类，简单说一个接口体里可以定义很多方法，比如:
```go

func(p person) setName(name string) {
 p.name = name
}
```
就是给person结构体定义了方法，还有就是在go中，首字母大写代表public，小写代表private

### 接口
```
type animal interface {
    description() string
}
```

只要实现了这个方法的结构体都算实现了这个接口，都可以被传到这个接口对象上。

### 通道
通道可以在两个Go routine之间传输信息，比如在lab中，就是通过通道来进行状态转移等等，通道在传输过程中是阻塞的，且是双向阻塞，也就是说如果发送了没接收，那发送端阻塞，接收的时候还没有发送，那接受端阻塞。通道和锁一起使用的时候要特别注意，很容易死锁，比如某个线程拿了锁，然后发送了channel，结果那边又在 等这个锁，就死了，做lab的时候这个地方吃了大亏。
