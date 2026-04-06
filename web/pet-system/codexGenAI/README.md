# Coffee Inventory System

`codexGenAI` 目录下包含一套完整的咖啡库存系统：

- `frontend`：Next.js + React 前端管理界面
- `backend`：.NET 8 Minimal API + EF Core
- `docker-compose.yml`：本地启动 MySQL 8 的可选方式

## 功能

- 咖啡库存列表展示
- 新增、编辑、删除库存记录
- 低库存预警
- 库存总量、库存货值、产地覆盖统计
- 初始化示例数据

## 后端启动

1. 准备 MySQL，确保数据库连接与 `backend/appsettings.json` 一致：

```json
"Server=localhost;Port=3306;Database=coffee_db;User=root;Password=password;"
```

2. 启动数据库后执行：

```bash
cd backend
dotnet restore
dotnet run
```

默认 API 地址为 `http://localhost:5216`。

## 前端启动

1. 进入前端目录并安装依赖：

```bash
cd frontend
npm install
```

2. 复制环境变量并按需修改：

```bash
cp .env.example .env.local
```

3. 启动开发服务器：

```bash
npm run dev
```

默认前端地址为 `http://localhost:3000`。

## 可选：使用 Docker 启动 MySQL

```bash
cd codexGenAI
docker compose up -d
```

数据库启动后，后端首次运行会自动建表并写入种子数据。
