package redis

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis/v8"
	"log"
)

var CTX = context.Background()


var Squirrel *redis.Client

// Init 初始化/*
func Init() {
	Squirrel = createRedisClient()
}

/*
 * 创建redis连接
 */
func createRedisClient() *redis.Client  {
	var host = beego.AppConfig.String("redis-host")
	var password = beego.AppConfig.String("redis-pwd")
	var db, _ = beego.AppConfig.Int("redis-db")

	log.Println("redis 连接参数：", host, password, db)
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	// 通过 client.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping(CTX).Result()
	log.Println(pong, err)

	return client
}