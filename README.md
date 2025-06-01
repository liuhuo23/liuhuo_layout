# liuhuo_layout 项目结构

```
.
├── .gitignore                # Git忽略规则文件
├── Makefile                  # 项目构建文件
├── README.md                 # 项目说明文档
├── api/                      # API协议定义目录
├── bin/                      # 可执行文件输出目录
├── cmd/                      # 命令行程序入口
│   └── liuos/                # 主程序包
│       ├── main.go           # 程序入口文件
│       ├── wire.go           # 依赖注入定义
│       └── wire_gen.go       # 自动生成的依赖注入代码
├── config/                   # 配置文件目录
│   └── config.yaml          # 应用配置文件
├── go.mod                    # Go模块定义
├── go.sum                    # Go模块校验文件
├── internal/                 # 内部实现代码
│   ├── app/                  # 应用层代码
│   │   └── app.go            # 应用初始化逻辑
│   ├── conf/                 # 配置相关代码
│   │   ├── config.pb.go      # Protobuf生成的配置代码
│   │   └── config.proto     # Protobuf配置定义
│   ├── data/                 # 数据访问层
│   │   └── data.go           # 数据访问实现
│   ├── server/               # 服务层代码
│   │   └── http.go           # HTTP服务实现
│   └── services/             # 业务服务目录
│       └── .gitkeep         
└── third_party/              # 第三方依赖
    └── .gitkeep            
```

# 安装不要文件
```shell
make init
```

# 编译
```shell
make build
```
# 运行
```shell
make run
```
# 按照 proto 生成客户端文件
规则：
`./api/[应用 名称]/v1/[服务名称].proto`
```shell
make client name=[应用名称]
```