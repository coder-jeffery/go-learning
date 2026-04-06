网关：
    基于：OpenResty（Nginx+LuaJIT）
        Nginx服务器：
            前端服务器 （Web 服务器）
                Nginx 
                Apache Http Server
                Lighttpd
                Caddy
                Tomacat / Jetty / Undertow / Express / Koa / net /http/  kestrel / OpenResty
                Windows系统：IIS
            反向代理 （浏览器 ↔ Nginx ↔ 后端服务）
                优点：
                    1.隐藏用户IP，安全
                    2.统一入口，负载均衡
                    3.做缓存，限流，权限，SSL
                    4.前后端分离部署
                    CDN vs DNS vs Nginx： 浏览器 → DNS 解析 → 得到 IP → 访问 CDN/Nginx → 微服务
                    CDN：Content Delivery Network 内容分发网络
                    DNS：Domain Name System 域名系统
            负载均衡
                硬件 LB：F5、A10（企业级，性能强）
                软件 LB：Nginx、HAProxy、LVS（云原生常用）
                云厂商 LB：阿里云 SLB、AWS ELB、腾讯云 CLB
                Nginx负载均衡:
                    轮询（默认）：依次分发
                    weight 加权：性能不同的机器
                    ip_hash：会话保持
                    least_conn：压力均衡
            网关入口
                统一入口、路由转发、负载均衡、安全防护、流量控制
K8S：
    Nginx 网关 = Ingress Controller

Nginx在K8S作用：
    1.NGINX Ingress Controller（最主流）
    2.独立部署 Nginx 做业务网关
    3.作为 sidecar 代理单个服务

K8s + Nginx Ingress + 微服务  
有 K8s 了，Nginx 少干了什么 (Nginx 专心做网关、代理、流量管控) | (K8s = 容器编排平台（管应用生命周期）) | (Nginx = 流量入口网关（管外网访问、路由、安全、加速）)
    未出现K8S：1.手动配置 upstream 列表 2.手动维护后端 IP 3.手动扩缩容后改配置
    出现K8S后：1.服务发现 2.负载均衡 3.自动扩缩容 4.自愈
    K8s + Nginx Ingress + 微服务
    