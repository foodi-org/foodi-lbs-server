# foodi-lbs-server

location based services of foodi


### 项目组件修改

在原`go-zero`基础上做了相应调整.

* 源生加载配置方式修改，配置文件修改为使用`consul`,在`NewServiceContext`时解析配置文件并建立`consul`连接.
* 服务发现修改
  * 服务注册放弃使用`etcd`,改为使用`consul`。服务注册需要显示声明
    ```go
       // 注册中心使用 consul
       if err := zrpcconsul.RegisterService(c.ListenOn, c.Consul); err != nil {
           panic(err)
       }
  
  * 服务发现相应的使用`consul`

* 静态配置文件`/etc`修改，目录下配置只保留`consul`连接配置，请按照环境区分。不能再使用`go-zero`源生方法加载配置。在`NewServiceContext`中 init 配置。
  ```go
  # consul://[user:passwd]@host/service?param=value'
  # format like this
  Add:
  Target: consul://127.0.0.1:8500/add.rpc?wait=14s
  Check:
  Target: consul://127.0.0.1:8500/check.rpc?wait=14s

  # ACL token support
  Add:
  Target: consul://127.0.0.1:8500/add.rpc?wait=14s&token=f0512db6-76d6-f25e-f344-a98cc3484d42
  Check:
  Target: consul://127.0.0.1:8500/check.rpc?wait=14s&token=f0512db6-76d6-f25e-f344-a98cc3484d42

* 数据库改为使用`gorm`。在`/svc/servicecontext.go`中加载。

### 注意

* 更改包名: 请使用有效的github包名。统一在`foodi-org`组织下。
* 基于`proto`文件生成代码时请使用服务分组模式。

  ```shell
  // 服务分组带上参数 -m
  goctl rpc protoc foodiLBS.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m
  ```
