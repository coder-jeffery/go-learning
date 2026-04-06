嵌套创建文件：
结构体：type xxxx struct
func (xx T) method()

组件类型	推荐工具	类比 Java 生态
服务注册发现	ETCD/Consul/Nacos	Eureka/Nacos
配置中心	ETCD/Nacos/Viper	Nacos Config/Apollo
熔断限流	Sentinel-Go/Hystrix-Go/Go-Zero 内置	Sentinel/Hystrix
链路追踪	Jaeger/Zipkin	SkyWalking/Zipkin
日志	Zap/Logrus	Logback/Log4j2
ORM	GORM/XORM	MyBatis/Hibernate
网关	Hertz/Go-Zero Gateway/APISIX	Spring Cloud Gateway


特性	Java	Go
类型系统	强类型，泛型（1.5+）	强类型，极简泛型（1.18+），无装箱拆箱
并发模型	线程 + 锁，CompletableFuture	Goroutine（轻量级线程）+ Channel（通道）
内存管理	JVM GC，OOM 风险高	更轻量的 GC，goroutine 栈动态扩容
语法	冗长（try-catch、分号、访问修饰符）	极简（无分号、无 class、无继承）
工程结构	Maven/Gradle，包名 = 目录结构	Go Module，GOPATH/GOMOD 模式
错误处理	异常（Checked/Unchecked）	显式返回 error（无 try-catch）
面向对象	类、继承、接口（需显式实现）	结构体 + 方法，组合代替继承，接口隐式实现

核心语法点：
    结构体（struct）代替 Java 的 class，方法绑定（func (t T) method()）
        常用数据接受方式：
            1 值接收方式
            2 指针接收方式
            3 访问控制：
                大写：public 外包可见
                小写：private 内包可见
                外包和内包区别：
    接口（interface）：隐式实现（无需implements），Go 的 “鸭子类型”
        结构体 + 绑定方法
    错误处理：error接口 + 多返回值（func f() (int, error)）
        
    协程（goroutine）：go func() 一行启动，理解其轻量级（百万级并发）
        
    通道（channel）：chan用于 goroutine 通信，替代 Java 的锁 / 队列

    内置函数：make（创建切片 / 通道 / 映射）、new（分配内存）、append等

    切片（slice）：动态数组，区别于 Java 的 ArrayList

    映射（map）：哈希表，无需像 Java 一样指定泛型