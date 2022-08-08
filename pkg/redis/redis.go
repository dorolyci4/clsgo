// Package redis use gofame-gredis implement common-usage redis functions.
package redis

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/log"
	"sync"
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

// Name valid in [default/cache], initialized with config.yaml
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
	ctx := context.Background()
	return client.conn.Do(ctx, command, args...)
}

func (client *Client) Receive() (*g.Var, error) {
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
