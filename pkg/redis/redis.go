// Package redis use gofame-gredis implement common-usage redis functions.
package redis

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	// "github.com/lovelacelee/clsgo/pkg/log"
)

type Client struct {
	Instance *gredis.Redis
}

func init() {
	gredis.SetConfig(&gredis.Config{
		Address: "192.168.3.207:6379, 192.168.3.207:6379",
		Db:      1,
	}, "default")
}

func New() *Client {
	client := Client{
		Instance: gredis.Instance("default"),
	}
	return &client
}

func (client *Client) Close() {
	ctx := context.Background()
	client.Instance.Close(ctx)
}

func (client *Client) SET(args ...any) (*g.Var, error) {
	ctx := context.Background()
	return client.Instance.Do(ctx, "SET", args...)
}
