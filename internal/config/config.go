package config

import (
	"github.com/foodi-org/foodi-lbs-server/internal/pkg/pkgconsul"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var (
	servConf = Config{}
)

type (
	// Config 项目配置
	Config struct {
		zrpc.RpcServerConf
		Mysql struct {
			DataSource string
		}
		Consul consul.Conf
	}

	SrvYaml struct {
		UnmarshalKey string     `yaml:"UnmarshalKey"`
		ConsulYaml   ConsulYaml `yaml:"Consul"`
	}

	ConsulYaml struct {
		Key   string   `yaml:"Key"`
		Meta  Meta     `yaml:"Meta"`
		Tag   []string `yaml:"Tag"`
		Host  string   `yaml:"Host"`
		Token string   `yaml:"Token"`
		TTL   int      `yaml:"TTL"`
	}

	Meta struct {
		Protocol string `yaml:"Protocol"`
	}

	// LBSConf nacos 配置映射结构体
	LBSConf struct {
		// ServiceName 服务名
		ServiceName string `json:"serviceName"`

		// ListenOn 声明监听端口 0.0.0.0:xxx
		ListenOn string `json:"listenOn"`

		Mysql struct {
			Datasource string `json:"datasource"`
		} `json:"mysql"`

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
		srvYaml SrvYaml
		lbsConf LBSConf
	)

	// 解析yaml配置文件
	file, err := os.ReadFile(filepath.Join(path, "etc", filename))
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, &srvYaml); err != nil {
		return err
	}

	servConf.Consul = consul.Conf{
		Host:  srvYaml.ConsulYaml.Host,
		Key:   srvYaml.ConsulYaml.Key,
		Token: srvYaml.ConsulYaml.Token,
		Tag:   srvYaml.ConsulYaml.Tag,
		TTL:   srvYaml.ConsulYaml.TTL,
	}

	if err = servConf.Consul.Validate(); err != nil {
		return err
	}

	consulClient := pkgconsul.NewConsulDO()
	consulClient.SetConfig(servConf.Consul)

	// consul connect
	if err = consulClient.Connect(); err != nil {
		panic(err)
	}

	// load service config from consul with key
	if err = consulClient.LoadJsonConfig(srvYaml.UnmarshalKey, &lbsConf); err != nil {
		panic(err)
	}

	// 设置 zrpc service config
	servConf.Name = lbsConf.ServiceName
	servConf.ListenOn = lbsConf.ListenOn
	servConf.Mysql.DataSource = lbsConf.Mysql.Datasource

	servConf.Redis = redis.RedisKeyConf{
		RedisConf: redis.RedisConf{
			Host: lbsConf.Redis.Host,
			Type: lbsConf.Redis.Type,
			Pass: lbsConf.Redis.Password,
			Tls:  lbsConf.Redis.TLS,
		},
		Key: lbsConf.ServiceName,
	}

	// 日志配置
	servConf.Log = logx.LogConf{
		ServiceName: lbsConf.ServiceName,
		Mode:        "file",
		//TimeFormat:  "",
		Path:       path,
		Level:      "info",
		KeepDays:   10,
		MaxBackups: 10,
		Rotation:   "daily",
	}

	return nil
}
