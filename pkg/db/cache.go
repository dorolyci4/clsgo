package db

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

/*
By default, the *gcache.Cache object of the ORM provides a single-process memory Cache,
which is highly efficient but can only be used within a single process.
If the service is deployed with multiple nodes, the cache between multiple nodes
may cause data inconsistency. Therefore, in most scenarios, we use Redis server
to cache database query data. The gcache.Cache object uses the adaptor design pattern to
easily switch from a single-process memory Cache to a distributed Redis Cache.
*/

func init() {

}

func GlobalCacheSetRedis() {
	redisCache := gcache.NewAdapterRedis(g.Redis())
	g.DB().GetCache().SetAdapter(redisCache)
}
