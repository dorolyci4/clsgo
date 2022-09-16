// Package redis use gofame-gredis implement common-usage redis functions.
package redis

import (
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/log"
)

// Redis Client wrappers
type Client struct {
	instance *gredis.Redis
	conn     *gredis.RedisConn
	// Is Subscriber loop running or not
	subscriberAlive bool
	// Subscriber receive g.Var and output
	subscriberNotify chan *g.Var
	wg               sync.WaitGroup
}

func init() {
	// Note that g.Redis() return nil in test mode, before SetConfig called.

	gredis.SetConfig(loadRedisConfig("redis.default"), "default")
	gredis.SetConfig(loadRedisConfig("redis.cache"), "cache")

}

func loadRedisConfig(configpath string) *gredis.Config {
	c := gredis.Config{
		Address:         config.Cfg.GetString(configpath + ".address"),
		Db:              config.Cfg.GetInt(configpath + ".db"),
		Pass:            config.Cfg.GetString(configpath + ".pass"),
		MinIdle:         config.Cfg.GetInt(configpath + ".minIdle"),
		MaxIdle:         config.Cfg.GetInt(configpath + ".maxIdle"),
		MaxActive:       config.Cfg.GetInt(configpath + ".maxActive"),
		IdleTimeout:     config.Cfg.GetDuration(configpath + ".idleTimeout"),
		MaxConnLifetime: config.Cfg.GetDuration(configpath + ".maxConnLifetime"),
		WaitTimeout:     config.Cfg.GetDuration(configpath + ".waitTimeout"),
		DialTimeout:     config.Cfg.GetDuration(configpath + ".dialTimeout"),
		ReadTimeout:     config.Cfg.GetDuration(configpath + ".readTimeout"),
		WriteTimeout:    config.Cfg.GetDuration(configpath + ".writeTimeout"),
		MasterName:      config.Cfg.GetString(configpath + ".masterName"), //Used in Sentinel mode
		TLS:             config.Cfg.GetBool(configpath + ".tls"),
		TLSSkipVerify:   config.Cfg.GetBool(configpath + ".tlsSkipVerify"),
	}
	return &c
}

// Name valid in [default/cache], initialized with config.yaml,
// See order for more https://www.redis.net.cn/order/
func New(name string) *Client {
	var (
		ctx = context.Background()
	)
	client := Client{
		instance:         gredis.Instance(name),
		subscriberAlive:  false,
		subscriberNotify: make(chan *g.Var),
		wg:               sync.WaitGroup{},
	}
	conn, err := client.instance.Conn(ctx)
	if err != nil {
		log.Errori(err)
	}
	client.conn = conn
	go subscriberRoutine(&client)
	return &client
}

func subscriberRoutine(client *Client) {
	client.subscriberAlive = true
	client.wg.Add(1)
	for client.subscriberAlive {
		reply, err := client.Receive()
		if err != nil {
			break
		}
		if reply != nil {
			client.subscriberNotify <- reply
		}
	}
	client.wg.Done()
}

func (client *Client) Close() {
	client.subscriberAlive = false
	client.wg.Wait()
	close(client.subscriberNotify)
	ctx := context.Background()
	client.conn.Close(ctx)
	// When instance closed, subscriber do not work any more
	// client.instance.Close(ctx)
}

func (client *Client) Do(command string, args ...any) (*g.Var, error) {
	if client.conn == nil {
		return nil, errors.New("redis not connected")
	}
	ctx := context.Background()
	return client.conn.Do(ctx, command, args...)
}

func (client *Client) Receive() (*g.Var, error) {
	if client.conn == nil {
		return nil, errors.New("redis not connected")
	}
	ctx := context.Background()
	return client.conn.Receive(ctx)
}

// Redis subscribe, return a go chan for receive message notification
func (client *Client) Subscribe(channel string) chan *g.Var {
	_, err := client.Do("SUBSCRIBE", channel)
	if err != nil {
		log.Errori(err)
	}
	return client.subscriberNotify
}
