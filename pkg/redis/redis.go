// Package redis use gofame-gredis implement common-usage redis functions.
package redis

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/log"
)

type Client struct {
	Instance   *gredis.Redis
	Connection *gredis.RedisConn
}

func init() {
	// Note that g.Redis() return nil, before SetConfig called.

	gredis.SetConfig(loadRedisConfig("redis.default"), "default")
	gredis.SetConfig(loadRedisConfig("redis.cache"), "cache")
	gredis.SetConfig(loadRedisConfig("redis.group"), "group")

	// log.Info(gredis.GetConfig("default"))
	// log.Info(gredis.GetConfig("cache"))
	// log.Info(gredis.GetConfig("group"))
}

func loadRedisConfig(configpath string) *gredis.Config {
	c := gredis.Config{
		Address:         clsgo.Cfg.GetString(configpath + ".address"),
		Db:              clsgo.Cfg.GetInt(configpath + ".db"),
		Pass:            clsgo.Cfg.GetString(configpath + ".pass"),
		MinIdle:         clsgo.Cfg.GetInt(configpath + ".minIdle"),
		MaxIdle:         clsgo.Cfg.GetInt(configpath + ".maxIdle"),
		MaxActive:       clsgo.Cfg.GetInt(configpath + ".maxActive"),
		IdleTimeout:     clsgo.Cfg.GetDuration(configpath + ".idleTimeout"),
		MaxConnLifetime: clsgo.Cfg.GetDuration(configpath + ".maxConnLifetime"),
		WaitTimeout:     clsgo.Cfg.GetDuration(configpath + ".waitTimeout"),
		DialTimeout:     clsgo.Cfg.GetDuration(configpath + ".dialTimeout"),
		ReadTimeout:     clsgo.Cfg.GetDuration(configpath + ".readTimeout"),
		WriteTimeout:    clsgo.Cfg.GetDuration(configpath + ".writeTimeout"),
		MasterName:      clsgo.Cfg.GetString(configpath + ".masterName"), //Used in Sentinel mode
		TLS:             clsgo.Cfg.GetBool(configpath + ".tls"),
		TLSSkipVerify:   clsgo.Cfg.GetBool(configpath + ".tlsSkipVerify"),
	}
	return &c
}

// name valid in [default/cache/group]
func New(name string) *Client {
	var (
		ctx = context.Background()
	)
	client := Client{
		Instance: gredis.Instance(name),
	}

	// r, err := gredis.New()
	// log.Info(r, err)
	conn, err := client.Instance.Conn(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}
	client.Connection = conn
	log.Info(client)
	return &client
}

func (client *Client) Close() {
	ctx := context.Background()
	client.Connection.Close(ctx)
}

func (client *Client) SET(args ...any) (*g.Var, error) {
	ctx := context.Background()
	return client.Connection.Do(ctx, "SET", args...)
}
