# 依赖

- gin - web服务器
- **validator - 用于验证传参**
- viper - 处理配置文件
- zap - 日志记录

# 目录结构

```bash
.
├── README.md
├── api
├── cmd
│   └── template-api
│       └── main.go
├── configs
│   └── conf.yaml
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── common
│   │   └── logger
│   │       └── logger.go
│   ├── models
│   │   ├── common.go
│   │   ├── conf.go
│   │   └── sample.go
│   ├── proxy
│   └── template-api
│       ├── conf
│       │   └── conf.go
│       ├── constants.go
│       ├── dao
│       │   ├── dao.go
│       │   └── sample.go
│       ├── http
│       │   ├── controller
│       │   │   ├── controller.go
│       │   │   └── sample.go
│       │   ├── http.go
│       │   ├── middleware
│       │   │   └── logHandler.go
│       │   ├── router.go
│       │   └── validate.go
│       └── logic
│           └── sample.go
├── script
├── sql
│   └── sample.sql
├── test
└── tools
    └── tools.go
```

cmd — 目录保存主函数入口.

internal — 主要逻辑入口

internal/common — 存储公共函数, 比如log设置.

internal/models — 存储主要数据结构

internal/proxy — 存放外部服务接口

internal/script — 存放项目脚本

internal/sql — 存放sql 文件

internal/test — 存放测试文件

internal/tools — 存放公共方法

internal/template-api — 程序逻辑, 可自定义 目录/包 名字.

internal/template-api/dao — 数据库交互逻辑

internal/template-api/http — httpServer 相关配置

internal/template-api/logic — 存放主逻辑函数

## 模板代码逻辑

> 一个mvc 架构
>
- main.go 负责注册structure. 并启动web server
- 以HttpServer 当做入口. 进入主逻辑

```go
type HttpServer struct {
   c  *conf.AppConfig
   v  *Validator
   cl *controller.Controller
   s  *gin.Engine
}
```

- Controller 结构定包含的controller

```
type Controller struct {
   SampleController *SampleController
}
```

- SampleController 中包含用到的的逻辑层

```
type SampleController struct {
   logic *logic.SampleLogic
}
```

- 逻辑层中包含了 用到的DAO层

```
type SampleLogic struct {
   dao *dao.Dao
}
```

- Dao 中 存放了配置和 db相关信息

```
type Dao struct {
   conf *conf.AppConfig
   db *gorm.DB
}
```

# 如何使用

> 模板中已经定义好log, conf 等配置, web server 可以直接写逻辑.
>

### 从一个简单的CRUD开始

1. 在dao中增加和数据库交互的的函数

```go
func (d *Dao)GetSampleList() (samples []models.Sample) {
	d.db.Find(&samples)
	return
}
```

2. logic 中增加相关查询逻辑.

```go
func (s *SampleLogic) GetSampeList() []models.Sample {
   return s.dao.GetSampleList()
}
```

3. controller 中增加控制层逻辑

```go
func (s SampleController) GetAllSampleList(c *gin.Context) {
   result := s.logic.GetSampeList()
   c.JSON(http.StatusOK, gin.H{"err":"", "result": result})
}
```

4. router 中增加相关控制逻辑

```go
noAuth.GET("/sample", h.cl.SampleController.GetAllSampleList)
```

### 添加的结构

所有使用到的结构都在main函数中New 出来. 因此. 新增任何一个结构都需要在main函数中创建.

1. 新建一个结构体
2. 创建构造函数.
3. main 函数中创建

以一个proxy 的创建为例

1. proxy中增加 nameProxy

```go
type NameProxy struct {
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
	resty *resty.Client
}
```

2. 生成构造函数

```go
func NewNameProxy(url string, token string) *NameProxy {
	return &NameProxy{Url: url, Token: token, resty: resty.New()}
}
```

3. 编写相关逻辑

```go
balabalbala
```

4. logic结构中增加相关的proxy结构

```go
type SampleLogic struct {
   dao       *dao.Dao
	 nameProxy *proxy.NameProxy
}
```

5. 更新logic的构造函数
6. main函数中定义相关变量. 传入需要的构造函数