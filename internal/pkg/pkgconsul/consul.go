package pkgconsul

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/consul/api"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var cli = &ConsulCli{}

type (
	ConsulCli struct {
		Consul consul.Conf
		client *api.Client
		err    error
	}

	ConsulDO interface {
		SetConfig(conf consul.Conf)

		Err() error

		/*Validate

		@Description: 验证 consul.Conf 配置必要参数
		@return error
		*/
		Validate() error

		/*Connect

		@Description: 建立consul连接
		@return error
		*/
		Connect() error

		/*LoadJsonConfig load and unmarshal config from consul by json key

		@Description: load and unmarshal config from consul by json key
		@param key: 读取key
		@param v: 反序列化结构体, 接收指针形式
		@return error
		*/
		LoadJsonConfig(key string, v interface{}) error
	}
)

func NewConsulDO() ConsulDO {
	return cli
}

func (c *ConsulCli) SetConfig(conf consul.Conf) {
	c.Consul = conf
}

func (c *ConsulCli) Err() error {
	return c.err
}

// Validate
//
//	@Description: 验证 consul.Conf 配置必要参数
//	@return error
func (c *ConsulCli) Validate() error {
	if len(c.Consul.Host) == 0 {
		return errors.New("consul host is required")
	}
	if len(c.Consul.Key) == 0 {
		return errors.New("consul key is required")
	}
	if c.Consul.TTL == 0 {
		c.Consul.TTL = 20
	}
	return nil
}

// Connect
//
//	@Description: 建立consul连接
//	@return error
func (c *ConsulCli) Connect() error {
	if c.client, c.err = api.NewClient(&api.Config{
		Address: c.Consul.Host,
		Scheme:  "http",
		Token:   c.Consul.Token,
	}); c.err != nil {
		logx.Must(c.err)
	}
	return c.err
}

// LoadJsonConfig load and unmarshal config from consul by json key
//
//	@Description: load and unmarshal config from consul by json key
//	@param key: 读取key
//	@param v: 反序列化结构体, 接收指针形式
//	@return error
func (c *ConsulCli) LoadJsonConfig(key string, v interface{}) error {
	kv := c.client.KV()
	data, _, err := kv.Get(key, nil)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("consul unmarshalKey key not found")
	}
	return json.Unmarshal(data.Value, v)
}
