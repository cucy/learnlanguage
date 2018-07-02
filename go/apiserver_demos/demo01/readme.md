实战：启动一个最简单的RESTful API服务器

# main.go

main.go 中的 main() 函数是 Go 程序的入口函数，在 main() 函数中主要做一些配置文件解析、程序初始化和路由加载之类的事情，最终调用 http.ListenAndServe() 在指定端口启动一个 HTTP 服务器。本小节是一个简单的 HTTP 服务器，仅初始化一个 Gin 实例，加载路由并启动 HTTP 服务器。

# **加载路由**
  
  `main()` 函数通过调用 `router.Load` 函数来加载路由（函数路径为 router/router.go，具体函数实现参照 [demo01/router/router.go](https://link.juejin.im/?target=https%3A%2F%2Fgithub.com%2Flexkong%2Fapiserver_demos%2Fblob%2Fmaster%2Fdemo01%2Frouter%2Frouter.go)）：
  
  ```go
  "apiserver/handler/sd"
      
      ....
      
      // The health check handlers
      svcd := g.Group("/sd")
      {   
          svcd.GET("/health", sd.HealthCheck)
          svcd.GET("/disk", sd.DiskCheck)
          svcd.GET("/cpu", sd.CPUCheck)
          svcd.GET("/ram", sd.RAMCheck)
      }
      ...
  ```
  
