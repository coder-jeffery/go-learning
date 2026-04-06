
my-project/
├── go.mod
├── go.sum          # 依赖校验文件，自动生成，提交到 Git
├── main.go         # 入口文件（可选，推荐放在 cmd 下）
├── cmd/            # 存放主要应用程序的 main 函数
│   └── my-project/
│       └── main.go
├── internal/       # 私有代码，不允许外部模块导入
│   ├── handler/    # HTTP 处理器
│   ├── service/    # 业务逻辑
│   └── model/      # 数据模型
├── pkg/            # 可被外部引用的公共库（如果需要）
├── config/         # 配置文件
├── scripts/        # 脚本
└── README.md


