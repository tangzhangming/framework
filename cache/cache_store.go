package cache

import (
	"context"
	"time"
)

type cache_store interface {

	//设置缓存驱动使用的context
	WithContext(ctx context.Context)

	//根据 key 获取一条缓存数据
	Get(key string) string

	//获取缓存，如果缓存不存在 将闭包的运行结果设为缓存并且返回
	GetOrSet(key string, make func() string) string

	//设置缓存
	Set(key string, val string, seconds time.Duration) bool

	//如果不存在则设置缓存
	Add(key string, val string, seconds uint) bool

	//删除缓存
	Delete(key string) bool

	//DefaultGet(key string, defaultVal string) string
	//
	//Set(key string, val string, seconds int) bool
	//
	//Exists(key string) bool
	//
	//Delete(key string) bool

}
