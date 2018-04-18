```
.
├─01_string_to_html字符串转html
├─02_parseExecute解析执行
│  ├─01_stdout读取模板到stdout
│  ├─02_tofile_解析执行模板到文件
│  ├─03_parse_files解析文件_ParseFiles_Execute结合
│  ├─04_parse_Glob_一次加载整个目录模板_分开执行各个模板
│  │  └─templates
│  └─05_performantParsing_高性能版本一次加载所有模板
│      └─templates
├─03_data
├─04_variable变量
├─05_dataStructures
│  ├─01_slice
│  │  ├─01
│  │  └─02_variable
│  ├─02_map
│  │  ├─01
│  │  └─02_variable
│  ├─03_struct
│  │  ├─01
│  │  └─02_variable
│  ├─04_sliceOfStruct
│  └─05_structSliceStruc结构体_结构体放在切片里
│      ├─01
│      └─02_anonStructType
├─06_func
│  ├─01
│  ├─02_dateFormatting_日期格式化
│  └─03_pipeline
├─07_predefinedGlobalFunctions_内置_全局函数
│  ├─01_index
│  │  ├─01
│  │  └─02
│  ├─02_and
│  └─03_comparison_比较大小
├─08_nestedTemplates模板嵌套
│  ├─01_nestedTemplates
│  │  └─templates
│  └─02_dataInTemplates_子模板传入数据
│      └─templates
├─09_composition_组合
│  ├─01_struct类型加数据类型判断
│  ├─02_method_方法在模板中使用
│  └─03_handsOnExercises_结构体嵌套结构体
├─10_characterEscaping_字符转义
│  ├─01_textTemplateNoEscaping_文字模板不转义
│  └─02_htmlTemplateEscaping_字符串进行转义
├─11_TCPServers
│  ├─03_readWrite
│  ├─04_readWriteSetDeadline_链接的生命周期为5秒
│  ├─05_dialRead_tcp连接
│  ├─06_dialWrite
│  ├─07_tcpApps
│  │  ├─01_rot13_字符串偏移13个
│  │  └─02_memoryDatabase
│  ├─08_tcpServerForHttp_简单的http服务器
│  ├─09_tcpMultiplexer
│  ├─1_write
│  └─2_read
├─12_netHttp
│  ├─01_listenAndServe
│  └─02_request
│      ├─01_parseForm
│      ├─02_requestStruct_请求首部结构
│      └─03_responseWriter_设置响应头
├─13_netHttpServeMux
│  ├─01_routing
│  ├─02_NewServeMux_多个请求路由
│  ├─03_defaultServeMux
│  ├─04_handleFunc_处理函数
│  └─05_handlerFunc
├─14_handsOnExercises_练习
│  ├─01
│  ├─02_HandleFunc_向模板传入数据
│  ├─03_Handle
│  ├─04_简单的tcpserver
│  ├─05_稍微复杂的tcpserver程序
│  ├─06_tcpserver设置断开时信息
│  ├─07_tcpserver_实现http服务
│  └─08
├─15_servingFiles
│  ├─01_notServing
│  ├─02_ioCopy
│  ├─03_serveContent
│  ├─04_serveFile
│  ├─05_fileServer
│  │  ├─01_列出当前的文件列表
│  │  ├─02_stripPrefix_前缀
│  │  │  └─assets
│  │  └─03_static
│  │      └─assets
│  │          ├─css
│  │          └─img
│  │              └─svg
│  ├─06_handsOn
│  │  ├─01
│  │  ├─02
│  │  │  ├─css
│  │  │  └─pic
│  │  ├─03_通用模板templates
│  │  │  ├─public
│  │  │  │  └─pics
│  │  │  └─templates
│  │  ├─04
│  │  │  ├─public
│  │  │  │  └─pics
│  │  │  └─templates
│  │  ├─05_静态文件
│  │  │  ├─public
│  │  │  │  ├─css
│  │  │  │  └─pic
│  │  │  └─templates
│  │  └─06
│  │      └─templates
│  └─07_notFoundHandler
├─16_passingData_表单上传相关
│  ├─01_url
│  ├─02_formPost
│  ├─03_formGet
│  ├─04_form
│  │  └─templates
│  ├─05_formFile
│  │  ├─01_read_读取文件
│  │  └─02_store_保存上传的表单文件
│  │      ├─sample-file
│  │      ├─templates
│  │      └─user
│  └─06_enctype
│      ├─01_default
│      │  └─templates
│      ├─02_multipart
│      │  └─templates
│      └─03_textPlain
│          └─templates
├─17_redirects
│  ├─01_303_seeOther
│  │  └─templates
│  ├─02_307_temporaryRedirect
│  │  └─templates
│  ├─03_301_movedPermanently_永久重定向
│  └─04_writeHeader_头部设置location重定向
│      └─templates
├─18_cookies
│  ├─01_setAndGet
│  ├─02_multipleCookies
│  ├─03_handsOn_手动设置cookie_浏览总量
│  └─04_maxAge
├─19_sessions
│  ├─01_uuid
│  ├─02_createSession
│  │  └─templates
│  ├─03_signup_注册
│  │  └─templates
│  ├─04_bcrypt
│  │  └─templates
│  ├─05_login
│  │  └─templates
│  ├─06_logout
│  │  └─templates
│  └─07_permissions
│      └─templates
├─20_rdbms_关系型数据库
│  ├─01_connect_连接
│  └─02_SQL
├─21_photoBlog
│  ├─public
│  │  ├─pics
│  │  └─picsToUpload
│  └─templates
├─22_HMAC_认证算法
│  ├─01
│  └─02
├─23_base64
├─24_context
│  ├─01
│  ├─02
│  ├─03
│  └─04
├─25_json
│  ├─01
│  ├─02_unmarshal
│  ├─03_unmarshal
│  ├─04_unmarshalTags
│  └─05_handsOn
└─26_ajax
    ├─01
    │  └─templates
    └─02
        └─templates
```