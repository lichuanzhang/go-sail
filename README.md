<div align="center">
    <h1><img src="static/sailboat-solid-colorful.svg" alt="sailboat-solid" title="sailboat-solid" width="300" /></h1>
</div> 

[![Go](https://github.com/keepchen/go-sail/actions/workflows/go.yml/badge.svg)](https://github.com/keepchen/go-sail/actions/workflows/go.yml)  [![CodeQL](https://github.com/keepchen/go-sail/actions/workflows/codeql.yml/badge.svg)](https://github.com/keepchen/go-sail/actions/workflows/codeql.yml)  [![Go Report Card](https://goreportcard.com/badge/github.com/keepchen/go-sail/v3)](https://goreportcard.com/report/github.com/keepchen/go-sail/v3)  

简体中文 | [English](./README_EN.md)

## go-sail是什么？  

**go-sail**是一个轻量的渐进式web框架，使用Go语言实现。它并**不是重复造轮子的产物**，而是站在巨人的肩膀上，整合现有的优秀组件，旨在帮助使用者以最简单的方式构建稳定可靠的服务。  
正如它的名字一般，你可以把它视作自己在golang生态的一个开始。go-sail将助力你从轻出发，扬帆起航。  

## 如何使用  
> 推荐go version >= 1.19  

> go get -u github.com/keepchen/go-sail/v3

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/keepchen/go-sail/v3/sail"
    "github.com/keepchen/go-sail/v3/sail/config"
)

var (
    conf = &config.Config{
        LoggerConf: logger.Conf{
            Filename: "logs/running.log",
        },
        HttpServer: config.HttpServerConf{
            Debug: true,
            Addr:  ":8000",
        },
    }
    registerRoutes = func(ginEngine *gin.Engine) {
        ginEngine.GET("/hello", func(c *gin.Context){
            c.String(http.StatusOK, "%s", "hello, world!")
        })
    }
)

sail.WakeupHttp("go-sail", conf).
    Hook(registerRoutes, nil, nil).
    Launch()
```  
当你看到终端如下图所示内容就表示服务启动成功了：  

<img src="static/launch.png" alt="launch.png" title="launch.png" width="600" />  

## 文档  
[文档传送门](https://go-sail.keepchen.com)

## 特性  
- 获取组件  
> go-sail启动时，会根据配置文件启动相应的应用组件，可使用`sail`关键字统一获取  
```go
import (
    "github.com/keepchen/go-sail/v3/sail"
)

//获取日志组件
sail.GetLogger()

//获取数据库连接（读、写实例）
sail.GetDB()

//获取redis连接(单例模式)
sail.GetRedis()

//获取redis连接(cluster模式)
sail.GetRedisCluster()

//获取nats连接
sail.GetNats()

//获取kafka完整连接实例
sail.GetKafkaInstance()

//获取etcd连接实例
sail.GetEtcdInstance()
```  
更多组件持续开发中，也欢迎大家提PR👏🏻👏🏻

- 返回响应  
```go
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/keepchen/go-sail/v3/constants"
    "github.com/keepchen/go-sail/v3/sail"
)

//handler
func SayHello(c *gin.Context) {
    sail.Response(c).Builder(constants.ErrNone, nil, "OK").Send()
}
```  

- 返回响应实体  
```go
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/keepchen/go-sail/v3/constants"
    "github.com/keepchen/go-sail/v3/http/pojo/dto"
    "github.com/keepchen/go-sail/v3/sail"
)

type UserInfo struct {
    dto.Base
    Data struct {
        Nickname string `json:"nickname" validate:"required" format:"string"` //nickname
        Age int `json:"int" validate:"required" format:"number"` //age
    } `json:"data" validate:"required"` //body data
}

//implement dto.IResponse interface
func (v UserInfo) GetData() interface{} {
    return v.Data
}

//handler
func GetUserInfo(c *gin.Context) {
    var resp UserInfo
    resp.Data.Nickname = "go-sail"
    resp.Data.Age = 18
	
    sail.Response(c).Builder(constants.ErrNone, resp).Send()
}
```

#### 其他插件  
[README.md](plugins/README.md)  

## 大感谢  
感谢在体验、使用过程中提出宝贵建议和意见以及提供过其他各种帮助的各位小伙伴！  
- 配置模块化优化 [@fujilin](https://github.com/fujilin)  
- 响应器语法糖增强优化 [@lichuanzhang](https://github.com/lichuanzhang)  
- Logo美化 [@ShuaiRen34](https://twitter.com/ShuaiRen34)  

## 其他  
- 欢迎大家提PR: [pull request](https://github.com/keepchen/go-sail/compare)  
- 欢迎大家提出自己的想法: [issue](https://github.com/keepchen/go-sail/issues/new/choose)  
- 感谢你的star如果你喜欢这个项目的话 :)  

## 使用案例  
<img src="static/usecases/pikaster-metaland.png" alt="Pikaster" width="300" />
<img src="static/usecases/wingoal-metaland.png" alt="WinGoal" width="300" />
<img src="static/usecases/miniprogram-hpp.png" alt="生活好评助手-小程序" width="180" />

