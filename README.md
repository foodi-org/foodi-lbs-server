# foodi-lbs-server

location based services of foodi


### 项目组件修改

* 配置文件改为使用`nacos`。`/etc`目录下配置只保留`nacos`连接配置，请按照环境区分。

  > `/config/congif.go`文件中定义`nacos`配置结构体和解析方法. 在`/svc/servicecontext.go`中加载
  >
* 数据库改为使用`gorm`。在`/svc/servicecontext.go`中加载。

### 注意

* 更改包名: 请使用有效的github包名。统一在`foodi-org`组织下。
* 基于`proto`文件生成代码时请使用服务分组模式。

  ```shell
  // 服务分组带上参数 -m
  goctl rpc protoc foodiLBS.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
  ```
