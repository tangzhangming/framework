package redis

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"strconv"
	"sync"
)

type RedisClient struct {
	name    string //连接名称
	Client  *redis.Client
	Context context.Context
}

/**
 * 连接储存器 The Redis connections
 */
var connections map[string]*RedisClient

/**
 * 默认连接的名称 default connection name
 */
var dcname string

/**
 * once 确保全局的 Redis 对象只实例一次
 */
var once sync.Once

func Start() {
	//初始化连接储存器
	connections = make(map[string]*RedisClient)

	//设置默认连接
	dcname = "default"
}

/**
 * 根据 name 获取一个 RedisClient
 */
func Conn(name string) *RedisClient {
	connection, ok := connections[name]

	if ok {
		return connection
	}

	return nil
}

/**
 * 设置一个redis连接(RedisClient) 并且取名
 */
func SetConnection(name string, client *redis.Client) {
	rdb := &RedisClient{
		name:    name,
		Context: context.Background(),
		Client:  client,
	}

	//Ping
	_, err := rdb.Client.Ping(rdb.Context).Result()
	if err != nil {
		return
	}

	//定时Ping 保持连接
	//t := time.NewTicker(time.Second * 10)
	//for range t.C {
	//	rdb.Client.Ping(rdb.Context).Result()
	//	fmt.Printf("%s is ping \n", name)
	//}

	connections[name] = rdb
}

func SetDefaultConnectionName(name string) {
	dcname = name
}

func SetConnectionMap(name string, options map[string]string) {
	var err error

	database, err := strconv.Atoi(options["database"])
	if err != nil {
		fmt.Printf("redis连接 %s 的配置项 database=%s 有误 \n", name, options["database"])
	}

	client := redis.NewClient(&redis.Options{
		Addr:     options["host"] + ":" + options["port"],
		Password: options["password"],
		DB:       database,
	})

	//PING
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("redis连接 %s PING失败 \n", name)
	}

	SetConnection(name, client)
}

func defaultConn() *RedisClient {
	return Conn(dcname)
}

func (rds RedisClient) WithContext(ctx context.Context) {
	rds.Context = ctx
}

func (rds RedisClient) GetName() string {
	return rds.name
}
