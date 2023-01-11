package cache

var cache_store_list map[string]cache_store

/**
 * 获得一个缓存库实例
 * @param name string
 * @return 储存驱动 cache_store
 */
func Store(name string) cache_store {
	store, ok := cache_store_list[name]

	if ok {
		return store
	}

	return nil
}

/**
 * 设置一个缓存库
 */
func SetStore(name string, driver string, client interface{}) {

	if cache_store_list == nil {
		cache_store_list = make(map[string]cache_store)
	}

	var store cache_store

	if driver == "redis" {
		store = redis_cache_instance(client)

	} else if driver == "memcached" {
		// store = memcached_cache_instance(client)

	}

	cache_store_list[name] = store
}
