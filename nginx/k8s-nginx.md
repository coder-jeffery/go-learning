K8S:
    用户 → DNS → CDN → LB（云厂商SLB/ELB）→ NGINX Ingress → 微服务（Service）
    生产级 Ingress 网关配置:
        在 K8s 里，Nginx 网关 = Ingress Controller 用 Ingress 资源配置路由，自动管理 Nginx
        