## KubeVela 客户端连接池实现
该项目使用KubeVela官方接口文档生成客户端API文件，在此基础上实现了简单通用的Client Pool。

### 使用方式

```go
package service

import (
	"context"
	"errors"
	vela "github.com/taweizhong/velaclient"
	"log"
	"time"
)

func NewConn() (interface{}, error) {
	velaUX := config.Get().VelaUX
	cfg := vela.NewConfiguration(vela.WithBasePath(velaUX.BaseUrl))
	client := vela.NewAPIClient(cfg)
	loginRes, _, err := client.AuthenticationApi.Login(context.Background(), vela.V1LoginRequest{
		Code:     "",
		Username: velaUX.UserName,
		Password: velaUX.PassWord,
	})
	if err != nil {
		log.Print(err)
	}
	token := loginRes.AccessToken
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	return client, nil
}

func CloseConn(client interface{}) error {
	if c, ok := client.(*vela.APIClient); ok {
		c.CloseConn()
	} else {
		return errors.New("client is not a Vela Client")
	}
	return nil
}

func PingConn(client interface{}) error {
	if c, ok := client.(*vela.APIClient); ok {
		token, _, err := c.AuthenticationApi.RefreshToken(context.Background())
		if err != nil {
			return err
		}
		cfg := c.Config()
		cfg.AddDefaultHeader("Authorization", "Bearer "+token.AccessToken)
	} else {
		return errors.New("client is not a Vela Client")
	}
	return nil
}

func NewVelaClintPool() vela.Pool {
	conf := &vela.Config{
		InitialCap:  10,
		MaxCap:      1000,
		MaxIdle:     20,
		Factory:     NewConn,
		Close:       CloseConn,
		Ping:        PingConn,
		IdleTimeout: 600 * time.Second,
	}
	pool, err := vela.NewChannelPool(conf)
	if err != nil {
		return nil
	}
	return pool
}
```



> [!WARNING]
>
> 注意事项：部分配置的底层实现使用通用接口，在实际使用时按需修改。

