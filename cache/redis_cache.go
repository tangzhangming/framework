package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type redis_cache struct {
	context context.Context
	redis   *redis.Client
}

func redis_cache_instance(client interface{}) *redis_cache {
	rclient, ok := client.(*redis.Client)

	if !ok {
		return nil
	}

	return &redis_cache{
		context: context.Background(),
		redis:   rclient,
	}
}

func (r redis_cache) WithContext(ctx context.Context) {
	r.context = ctx
}

func (r redis_cache) Get(key string) string {
	value, err := r.redis.Get(r.context, key).Result()

	if err != nil {
		return ""
	} else {
		return value
	}
}
func (r redis_cache) GetOrSet(key string, make func() string) string {
	var val string

	val = r.Get(key)
	if val == "" {
		val = make()
		r.Set(key, val, time.Hour)
	}

	return val
}

func (r redis_cache) Set(key string, val string, seconds time.Duration) bool {
	if r.redis.Set(r.context, key, val, seconds).Err() != nil {
		return false
	} else {
		return true
	}
}

func (r redis_cache) Add(key string, val string, seconds uint) bool {
	var incrBy = redis.NewScript("return redis.call('exists',KEYS[1])<1 and redis.call('setex',KEYS[1],ARGV[2],ARGV[1])")

	keys := []string{key}
	values := []interface{}{val, seconds}

	err := incrBy.Run(r.context, r.redis, keys, values).Err()
	if err != nil {
		return false
	} else {
		return true
	}
}

func (r redis_cache) Delete(key string) bool {
	if r.redis.Del(r.context, key).Err() != nil {
		return false
	} else {
		return true
	}
}
