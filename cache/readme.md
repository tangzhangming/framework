## 使用
```

//Store方法访问各种缓存存储
cc := cache.Store("redis")


val := cc.Get("key111")

cc.Set("key333", "val content", time.Hour)

val := cc.GetOrSet("key222", func() string {
    return "...build value"
})

```