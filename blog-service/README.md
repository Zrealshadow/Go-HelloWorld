# Go Http Server Example : Blog Service

>这是一个关于Go语言http后端的学习，下面记录一些可复用的代码和模式
>
> 该源码选择了部分功能进行了实现



- [x] 进行项目设计
- [x] 编写公共组件
- [x] 生成接口文档
- [x] 为接口文档做参数校验
- [x] 标签管理
- [x] 上传图片和文件服务
- [x] 对接口进行访问控制
- [x] 应用中间件
- [ ] 链路追踪
- [ ] 应用配置
- [ ] 应用编译
- [ ] 优雅重启和停止



### **项目目录结构**

- configs：配置文件。
- docs：文档集合。
- global：全局变量。
- internal：内部模块。
  - dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。
  - middleware：HTTP 中间件。
  - model：模型层，用于存放 model 对象。
  - routers：路由相关逻辑处理。
  - service：项目核心业务逻辑。
- pkg：项目相关的模块包。
- storage：项目生成的临时文件。
- scripts：各类构建，安装，分析等操作的脚本。
- third_party：第三方的资源工具，例如 Swagger UI。

其中要注意的是internal中的部分代码模式可以借鉴，pkg中的代码块可以直接复用











### Gin Web Framework

该项目的主体是依靠Gin网络框架

下面是一个简单的http Service， 该代码并不能运行，主要用来表示Gin的基本模式。

```go

func main() {
	router := NewRouter()
  // assign Handler and Addr and listen
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
	}
	s.ListenAndServe()
}

func NewRouter() *gin.Engine {
  // step one : create a new gin Engine
	r := gin.New()
  // step two : register middleware, there gin.Logger and gin.Recovery is two default middleware
  // some middleware only designed for groups instead of all router
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
  /*
  register middleware
  */
  
  // step three : declare some data model
	upload := NewUpload() // declare model 
  
  // step four : bind model process function to router
	r.POST("/upload/file", upload.UploadFile) 
  
  /*
  process Function is gin.HandlerFunc
  => func (ctx *gin.Context)
  
  Actually, we can directly bind function to router, no need to create a model
  */
	return r
}
```





### Model 创建Demo

使用gorm进行ORM操作，下面是gorm操作实例(以tag为例)，更多细节可以查看`internal/model/model.go`和`internal/model/tag.go`代码

```go
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
// public model contains public columns

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "BlogTag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 注意其中 Model() 方法 指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的 TableName 方法用于指定其对应返回的表名。

```



### 配置管理

使用了Viper对yaml数据进行了读取，下面是基本读取方法

```go
type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
// 这里注意AddConfigPath可以添加多个，在多个路径寻找配置文件

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
// 这里是将整个的Setting分配到各个global.Setting 结构上

// 例如声明
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
	DefaultContextTimeout time.Duration
}
```

对应的Config.yaml 文件

```yaml
Server:
  RunMode: debug
  HttpPort: 8000
  ReadTimeout: 60
  WriteTimeout: 60
App:
  DefaultPageSize: 10
  MaxPageSize: 100
  LogSavePath: storage/logs
  LogFileName: app
  LogFileExt: .log
  UploadSavePath: storage/uploads
  UploadServerUrl: http://127.0.0.1:8000/static
  UploadImageMaxSize: 5 #MB
  UploadImageAllowExts:
    - .jpg
    - .jpeg
    - .png
  DefaultContextTimeout: 60
```

通常在main函数开头读取并设置配置文件: main.go

```go
func init(){
  err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err : %v", err)
	}
}
//init 文件在 main函数开始之前执行

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
  // 通过ReadSection 将二级配置分配到全局变量上
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
```



### Swagger Example

注解编写demo

```go
// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

// 注意: 这些Param输入参数都放在了 Struct Tag t 的属性里面了

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
```

main函数注解

```go
// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	...
}
```



注解完成之后需要生成,回到根目录执行一下命令，在docs下生成了docs.go swagger.json swagger.yaml三个文件

```shell
$ swag init
```



在该服务器下还得编写路由

```go
import (
	...
	_ "github.com/go-programming-tour-book/blog-service/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
  
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  
	/*
	...
	*/
	return r
}
```





### 参数校验

对于每一个Router，最好我们在service层声明每一个request所携带的参数的结构体，并在对应结构体字段的标签上定义校验规则。

应用到了两个 tag 标签，分别是 form 和 binding

常见标签如下

| 标签     | 含义                      |
| -------- | ------------------------- |
| required | 必填                      |
| gt       | 大于                      |
| gte      | 大于等于                  |
| lt       | 小于                      |
| lte      | 小于等于                  |
| min      | 最小值                    |
| max      | 最大值                    |
| oneof    | 参数集内的其中之一        |
| len      | 长度要求与 len 给定的一致 |

```go
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

```

其中form指的是request body中该值的Key，通过`gin.Context.SHouldBind() 方法将body中的Param绑定到struct中的属性上`







### 其他公共组件复用

- 数据库连接
- 错误代码标准化
- 日志标准化
- 响应处理 （这个提供了几个非常好用的wrapper，对于Response来说）
- 上传文件服务可以直接进行复用



### 应用中间件复用

不要忘记在Router中进行中间件的注册

其中接口限流控制，可以通过限制 获取权限接口的流量来控制整个App的访问流量

- 访问日志
- 异常捕获处理
- 服务信息储存
- 接口限流控制
- 统一超时控制

