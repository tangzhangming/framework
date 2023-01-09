package redis

import (
	redis "github.com/go-redis/redis/v8"
	"time"
)

/**
 * 下列方法为 redis.Client 中的方法二次封装
 * 1、Context统一传入 调用时少写参数
 * 2、支持快速操作 不需要调用Connection方法即可快速使用默认连接
 */

// value, err := Get().Result()
func Get(key string) *redis.StringCmd {
	return defaultConn().Get(key)
}
func (rds RedisClient) Get(key string) *redis.StringCmd {
	return rds.Client.Get(rds.Context, key)
}

// err := Set().Err()
func Set(key string, value string, expiration time.Duration) *redis.StatusCmd {
	return defaultConn().Set(key, value, expiration)
}
func (rds RedisClient) Set(key string, value string, expiration time.Duration) *redis.StatusCmd {
	return rds.Client.Set(rds.Context, key, value, expiration)
}

/**
 * 设置一个key的值，并返回这个key的旧值
 */
func GetSet(key string, value interface{}) *redis.StringCmd {
	return defaultConn().GetSet(key, value)
}
func (rds RedisClient) GetSet(key string, value interface{}) *redis.StringCmd {
	return rds.Client.GetSet(rds.Context, key, value)
}

/**
 * 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
 */
func SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return defaultConn().SetNX(key, value, expiration)
}
func (rds RedisClient) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	return rds.Client.SetNX(rds.Context, key, value, expiration)
}

/**
 * 递增操作
 */
func Incr(key string) *redis.IntCmd {
	return defaultConn().Incr(key)
}
func (rds RedisClient) Incr(key string) *redis.IntCmd {
	return rds.Client.Incr(rds.Context, key)
}
func IncrBy(key string, value int64) *redis.IntCmd {
	return defaultConn().IncrBy(key, value)
}
func (rds RedisClient) IncrBy(key string, value int64) *redis.IntCmd {
	return rds.Client.IncrBy(rds.Context, key, value)
}

/**
 * 递减操作
 */
func Decr(key string) *redis.IntCmd {
	return defaultConn().Incr(key)
}
func (rds RedisClient) Decr(key string) *redis.IntCmd {
	return rds.Client.Incr(rds.Context, key)
}
func DecrBy(key string, decrement int64) *redis.IntCmd {
	return defaultConn().IncrBy(key, decrement)
}
func (rds RedisClient) DecrBy(key string, decrement int64) *redis.IntCmd {
	return rds.Client.IncrBy(rds.Context, key, decrement)
}

/**
 * 删除key操作,支持批量删除
 * err := redis.Del("a", "b", "c").Err()
 */
func Del(keys ...string) *redis.IntCmd {
	return defaultConn().Del(keys...)
}
func (rds RedisClient) Del(keys ...string) *redis.IntCmd {
	return rds.Client.Del(rds.Context, keys...)
}

/**
 * 设置key的过期时间,单位秒
 */
func Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return defaultConn().Expire(key, expiration)
}
func (rds RedisClient) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return rds.Client.Expire(rds.Context, key, expiration)
}

/**
 * 给数据库中名称为key的string值追加value
 */
func Append(key string, value string) *redis.IntCmd {
	return defaultConn().Append(key, value)
}
func (rds RedisClient) Append(key string, value string) *redis.IntCmd {
	return rds.Client.Append(rds.Context, key, value)
}

/**
 * 事务
 *
 * 调用方法:
 * rdb := redis.Conn('default')
 * pipe := rdb.TxPipeline() //开始一个事务
 *     rdb.Incr()
 *     pipe.Expire() //可以用pipe操作redis
 * pipe.Exec() //提交事务
 */
func (rds RedisClient) TxPipeline() redis.Pipeliner {
	return rds.Client.TxPipeline()
}

/**
 * 使用Lua脚本操作Redis
 * script = redis.NewScript(scriptString)
 * script.Run(ctx, rdb, keys, values)
 */
func NewScript(src string) *redis.Script {
	return redis.NewScript(src)
}
