Nginx: Nginx 是基础服务器，只认配置，不认代码
Lua 是一门脚本语言，用来写逻辑。
OpenResty 把 Nginx 和 Lua 绑在一起，让你能用 Lua 扩展 Nginx
OpenResty 是基于 Nginx 核心深度定制的 Web 平台，集成了 LuaJIT 和 ngx_lua 模块，使开发者可以使用 Lua 脚本语言编写 Nginx 逻辑，实现动态路由、限流、鉴权、缓存、灰度发布等高级网关功能



                      功能	  原生 Nginx	 OpenResty(Nginx+Lua)
                      静态网站	✅	        ✅
                      反向代理	✅      	✅
                      负载均衡	✅      	✅
                      写代码逻辑	❌      	✅
                      鉴权、限流	弱      	强

