package config

import (
	"encoding/json"
	pkgnacos "github.com/foodi-org/foodi-pkg/nacos"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var (
	servConf Config
)

type (
	// Config 项目配置
	Config struct {
		zrpc.RpcServerConf
		Mysql struct {
			DataSource string
		}
	}

	// yamlConf yaml配置文件映射结构体
	yamlConf struct {
		Username  string   `yaml:"Username"`
		Password  string   `yaml:"Password"`
		Address   []string `yaml:"Address"`
		Namespace string   `yaml:"Namespace"`
		DataID    string   `yaml:"DataID"`
		Group     string   `yaml:"Group"`
	}

	// greetConf nacos 配置映射结构体
	greetConf struct {
		// ServiceName 服务名
		ServiceName string `json:"serviceName"`

		// ListenOn 声明监听端口 0.0.0.0:xxx
		ListenOn string `json:"listenOn"`

		Mysql struct {
			Datasource string `json:"datasource"`
		} `json:"mysql"`

		Etcd struct {
			Host []string `json:"host"` // etcd 集群
			Key  string   `json:"key"`  // 该服务注册到etcd的key
		} `json:"etcd"`

		Redis struct {

			// redis 服务地址 ip:port, 如果是 redis cluster 则为 ip1:port1,ip2:port2,ip3:port3...(暂不支持redis sentinel)
			Host string `json:"host"`

			// node:单节点 redis;cluster:redis 集群
			Type string `json:"type"`

			Password string `json:"password"`

			// 是否开启tls
			TLS bool `json:"tls"`
		} `json:"redis"`
	}
)

func ServConf() *Config {
	return &servConf
}

// InitServConf
//
//	@dDescription: 初始化 service 配置
//	@param path: 项目路径
//	@param filename: 使用的配置文件名称
//	@return error
func InitServConf(path string, filename string) error {
	var (
		conf  yamlConf
		gConf greetConf
		data  string
	)

	// 解析yaml配置文件
	if file, err := os.ReadFile(filepath.Join(path, "etc", filename)); err != nil {
		return err
	} else {
		if err = yaml.Unmarshal(file, &conf); err != nil {
			return err
		}

		// 加载nacos并解析配置
		n := pkgnacos.NewNacosClient(conf.Namespace, conf.Username, conf.Password, conf.Address)
		n.SetCacheDir(path)
		n.SetLogDir(path)

		if err = n.CreateConfigClient(); err != nil {
			return err
		}
		if data, err = n.GetConfig(conf.DataID, conf.Group); err != nil {
			return err
		}
		if err = json.Unmarshal([]byte(data), &gConf); err != nil {
			return err
		}

		// 设置 zrpc service config
		servConf.Name = gConf.ServiceName
		servConf.ListenOn = gConf.ListenOn
		servConf.Mysql.DataSource = gConf.Mysql.Datasource

		servConf.Etcd = discov.EtcdConf{
			Hosts: gConf.Etcd.Host,
			Key:   gConf.Etcd.Key,
		}

		servConf.Redis = redis.RedisKeyConf{
			RedisConf: redis.RedisConf{
				Host: gConf.Redis.Host,
				Type: gConf.Redis.Type,
				Pass: gConf.Redis.Password,
				Tls:  gConf.Redis.TLS,
			},
			Key: gConf.ServiceName,
		}

		// 日志配置
		servConf.Log = logx.LogConf{
			ServiceName: gConf.ServiceName,
			Mode:        "file",
			//TimeFormat:  "",
			//Path:        "",
			Level:      "info",
			KeepDays:   10,
			MaxBackups: 10,
			Rotation:   "daily",
		}
	}

	return nil
}
