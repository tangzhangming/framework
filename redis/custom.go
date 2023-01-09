package redis

/**
 * 下列方法为定制封装
 * @author tangzhangming@live.com
 */

func GetString(key string) string {
	return defaultConn().GetString(key)
}
func (rds RedisClient) GetString(key string) string {
	value, err := rds.Client.Get(rds.Context, key).Result()

	if err != nil {
		return ""
	}

	return value
}
