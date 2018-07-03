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
  


# demo04


`curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user`

``` 
 curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin"}'
{"code":10001,"message":"password is empty"} 


curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"password":"admin"}'
{"code":20102,"message":"The user was not found. This is add message."}

 curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"admin","password":"admin"}'
{"code":0,"message":"OK"}                              
```

# demo06 读取request 返回response

``` 
curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/admin2?desc=test -d'{"username":"admin","password":"admin"}'

```


# demo07创建用户逻辑

``` 
1.从 HTTP 消息体获取参数（用户名和密码）
2.参数校验
3.加密密码
4.在数据库中添加数据记录
5.返回结果（这里是用户名）

```

```json
create user

curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}'



get users

curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'
{"code":0,"message":"OK","data":{"totalCount":2,"userList":[{"id":2,"username":"kong","sayHello":"Hello M3d04aDiR","password":"$2a$10$nC7PnMjU2XsDHcbbnlodEujBBuKbrNG4vrvcNuzRCE.gMM3r7rBKC","createdAt":"2018-07-04 00:04:00","updatedAt":"2018-07-04 00:04:00"},{"id":1,"username":"admin","sayHello":"Hello M3dAVavmgz","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-27 16:25:33","updatedAt":"2018-05-27 16:25:33"}]}}%



get user details
curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong
{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$nC7PnMjU2XsDHcbbnlodEujBBuKbrNG4vrvcNuzRCE.gMM3r7rBKC"}}

update user 
 curl -XPUT -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"kong","password":"kongmodify'
{"code":0,"message":"OK","data":null}

curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong
{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$MWxOY4U02Y2ixn7WLr.XS.6E40BpqtjYWhJL0gqSCv.EMZCXZPLq."}}

    

delete user 
 curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2
{"code":0,"message":"OK","data":null}            


curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'
{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":1,"username":"admin","sayHello":"Hello JXpE4-vig","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-27 16:25:33","updatedAt":"2018-05-27 16:25:33"}]}}%

```


