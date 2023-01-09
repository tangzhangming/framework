## 初始化

redis.Start()

传入 github.com/go-redis/redis 连接 redis.SetConnection("default", client)

可以传入多个连接 redis.SetConnection("cache", client)
redis.SetConnection("redis2", client)

## 调用

```
rds1 := redis.Conn("default")
rds2 := redis.Conn("cache")

if rds1 == nil || rds2 == nil {
    ...错误处理
}

/**
 * 默认使用 context.Background() 作为上下文 
 * 在gin等框架中 你可以把 *gin.Context 传入作为上下文
 */
rds1.WithContext(ctx)
rds2.WithContext(ctx)

/**
 * 普通调用
 * conn.Client 可以获取到 github.com/go-redis/redis 包
 */
rds1.Get("key")
rds2.Get("cache_kay")
rds2.Client.HSet(...) 


/**
 * 快速调用
 * 默认使用名称为 default 的连接
 */  
redis.Get("key11") //等同于 redis.Conn("default").Get("key11")
redis.Set(...)
redis.Expire(...)
...

```

