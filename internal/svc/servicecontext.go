package svc

import (
	"fmt"
	"github.com/foodi-org/foodi-lbs-server/internal/config"
	foodipkg "github.com/foodi-org/foodi-pkg/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

var svc ServiceContext

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

// NewServiceContext
//
//	@Description: 加载 svc
//	@param c
//	@param dir 项目路径
//	@param file 配置文件名称
//	@return *ServiceContext
func NewServiceContext(c *config.Config, dir string, file string) error {

	// 加载配置
	if err := config.InitServConf(dir, file); err != nil {
		return err
	}

	// mysql client
	fmt.Println("mysql: ", c.Mysql.DataSource)
	foodipkg.InitConn(c.Mysql.DataSource)

	// redis client
	red := redis.MustNewRedis(redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		Pass:        c.Redis.Pass,
		Tls:         c.Redis.Tls,
		PingTimeout: 3 * time.Second,
	})
	svc = ServiceContext{
		Config: *c,
		Redis:  red,
	}
	return nil
}

func RedisCli() *redis.Redis {
	return svc.Redis
}

func Svc() *ServiceContext {
	return &svc
}
