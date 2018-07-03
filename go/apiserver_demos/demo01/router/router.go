package router

import (
    `learnlanguage/go/apiserver_demos/demo01/handler/sd`
    `learnlanguage/go/apiserver_demos/demo01/handler/user`
    "learnlanguage/go/apiserver_demos/demo01/router/middleware"

    "github.com/gin-gonic/gin"

    "net/http"
)

// 加载路由
// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
    // Middlewares.
    g.Use(gin.Recovery())
    g.Use(middleware.NoCache)
    g.Use(middleware.Options)
    g.Use(middleware.Secure)
    g.Use(mw...)

    // 404 Handler.
    g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect API route.")
    })


    // 用户路由设置
    u := g.Group("/v1/user")
    {
        u.POST("", user.Create)         // 创建用户
        u.DELETE("/:id", user.Delete)   // 删除用户
        u.PUT("/:id", user.Update)      // 更新用户
        u.GET("", user.List)            // 用户列表
        u.GET("/:username", user.Get)   // 获取指定用户的详细信息
    }


    // The health check handlers
    svcd := g.Group("/sd")
    {
        svcd.GET("/health", sd.HealthCheck)
        svcd.GET("/disk", sd.DiskCheck)
        svcd.GET("/cpu", sd.CPUCheck)
        svcd.GET("/ram", sd.RAMCheck)
    }

    return g

}

var s = `
该代码块定义了一个叫 "sd" 的分组，在该分组下注册了 /health、/disk、/cpu、/ram HTTP 路径，

分别路由到 sd.HealthCheck、sd.DiskCheck、sd.CPUCheck、sd.RAMCheck 函数。

sd 分组主要用来检查 API Server 的状态：健康状况、服务器硬盘、CPU 和内存使用量。

`
var h = `
router.Load 函数通过 g.Use() 来为每一个请求设置 Header，在 router/router.go 文件中设置 Header：

gin.Recovery()：在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器

middleware.NoCache：强制浏览器不使用缓存

middleware.Options：浏览器跨域 OPTIONS 请求设置 

middleware.Secure：一些安全设置

`
