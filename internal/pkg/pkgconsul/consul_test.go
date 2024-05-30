package pkgconsul

import (
	"github.com/foodi-org/foodi-lbs-server/internal/config"
	"github.com/hashicorp/consul/api"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"testing"
)

func TestConsulConf_LoadJsonConfig(t *testing.T) {
	type fields struct {
		Consul consul.Conf
	}
	type args struct {
		client *api.Client
		key    string
		v      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test get key of consul",
			fields: fields{Consul: consul.Conf{
				Host:  "8.134.206.241:8500",
				Key:   "",
				Token: "",
				Tag:   nil,
				Meta:  nil,
				TTL:   0,
			}},
			args: args{
				client: nil,
				key:    "foodi-dev/foodi-lbs-service",
				v:      config.LBSConf{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ConsulConf{
				Consul: tt.fields.Consul,
			}
			cli, err := c.NewClient()
			if err != nil {
				t.Error(err)
			}
			if err = c.LoadJsonConfig(cli, tt.args.key, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("LoadJsonConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
