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

Go数据类型：
    



        
            
    

    
