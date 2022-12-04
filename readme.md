# web项目模板

main.go为项目入口文件，其中init已经将配置，数据库等链接建立，main()主要是启动对应的服务

而服务都在applications目录下，一个目录是一个服务，web和job分别写了一个示例



library目录存储一些公共的资源
- business为项目用到的常量
- e为错误，里面定义了一个错误结构体以及各种错误码，在自定义的功能中返回*e.Error作为error，而不是使用系统自带的error
- redis_key统一管理项目中使用的各种redisKey

