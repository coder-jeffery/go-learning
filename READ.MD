golang web框架分两类:
    基于标准库net/http 比如:gin-gonic/gin, labstack/echo, astaxie/beego
    基于valyala/fasthttp  偏性能: kataras/iris , gofiber/fiber

代表性框架:
    gin[http] 相当于falsk
    iris[宇宙最快]
    beego [http] 相当于django  url: https://github.com/beego
        快速安装:
            go get -u github.com/beego/beego/v2
            go get -u github.com/beego/bee/v2
        创建项目
            bee api beego_api_demo
        创建web工程
            bee new projectname
help: https://go.dev/doc/go-get-install-deprecation
beego 最新下载方式:  go install github.com/beego/bee/v2@latest

配置代理：
export GO111MODULE=on
export GOPROXY="https://goproxy.cn,direct"

下载：
git clone https://github.com/beego/bee.git
git clone https://github.com/beego/beego.git

https://github.com/beego/beego
本地部署beego：https://blog.csdn.net/handsomewzj/article/details/144573586

step1.
go get beego_learning
step2.
生成swaggerui
bee run -gendoc=true -downdoc=true

Go:
    服务器编程 处理日志或者数据打包、虚拟机处理、文件系统。
    分布式系统：数据库代理
    网络编程 web应用、API应用、下载应用
    内存数据库
    云平台
    区块链 分布式数据存储 点对点传输 共识机制 加密算法
Go应用场景：
    消息推送 监控 容器化

Go入门基础：

Go:
1.Go语言的花括号不能单独放在一行  否则编译错误
2.Go语言不需要加分号
3.变量声明放在函数外面  变量赋值放在函数内部  变量未使用需要删除
4.Go format  / Tab键  --->  go fmt xx.go

Go:
25个关键字
37个保留字
变量：Variable 存储数据
静态编程语言

变量类型：
整型 浮点型 布尔型 字符串型
变量声明：
var 变量名 变量类型


golang func main入口函数
    变量和常量
    数据类型 (int/string/bool/float64) 复合类型(array/slice/map)
    流程控制、函数、包管理(Go 1.11+ 支持 Go Modules（go mod），替代传统的 GOPATH，是项目管理的核心)、
Go 特色类型：
    slice切片 动态数组
    map键值对
    struct自定义结构体
    错误处理
Go并发编程
    异步执行函数 goroutine
    channel通信
    同步工具
    接口： go鸭子类型接口
go标准库 std ：IO/网络/加密/并发/时间处理/并发工具/
Go项目驱动：
    命令行参数解析/文件存储/简易http请求/
进阶：
    简易聊天室 goroutine/channel 多人聊天室 / TCP/UDP网络编程
    并发管理
    Gin框架/Bingee框架/ORM框架/GORM框架
    



        
            
    

    
