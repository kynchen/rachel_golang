package redis

import (
	"log"
	"time"
)

var redisClient = &Squirrel
var redisContext = CTX

// ------------------------------------- string OPERATE ----------------------------------------------/*

func Set(key string, value interface{}, expire int) bool {
	err := (*redisClient).Set(redisContext, key, value, time.Second*time.Duration(expire)).Err()
	if err != nil {
		log.Println("redis缓存异常", err)

	}
	return true
}

func Get(key string) string {
	var redisResult string
	stringCmd := (*redisClient).Get(redisContext, key)
	if stringCmd != nil {
		result, _ := stringCmd.Result()
		redisResult = result
	}
	return redisResult
}

// ------------------------------------- HASH OPERATE ----------------------------------------------/*

func HSet(hKey string, value interface{}) bool {
	err := (*redisClient).HSet(redisContext, hKey, value).Err()
	if err != nil {
		log.Println("redis缓存异常", err)

	}
	return true
}

func HMSet(hKey string, value interface{}) bool {
	err := (*redisClient).HMSet(redisContext, hKey, value).Err()
	if err != nil {
		log.Println("redis缓存异常", err)
	}
	return true
}

func HSetNx(hKey string, key string, value interface{}) bool {
	err := (*redisClient).HSetNX(redisContext, hKey, key, value).Err()
	if err != nil {
		log.Println("redis缓存异常", err)

	}
	return true
}

func HGet(hKey string, key string) string {
	var redisResult string
	result := (*redisClient).HGet(redisContext, hKey, key)
	if result != nil {
		result, _ := result.Result()
		redisResult = result
	}
	return redisResult
}

func HMGet(hKey string,key string) []interface{} {
	var redisResult []interface{}
	result := (*redisClient).HMGet(redisContext, hKey,key)
	if result != nil {
		result, _ := result.Result()
		redisResult = result
	}
	return redisResult
}

func HGetAll(hKey string) map[string]string {
	var redisResult map[string]string
	result := (*redisClient).HGetAll(redisContext, hKey)
	if result != nil {
		result, _ := result.Result()
		redisResult = result
	}
	return redisResult
}

func HDelete(hKey string, key string) bool {
	(*redisClient).HDel(redisContext, hKey, key)
	return true
}
