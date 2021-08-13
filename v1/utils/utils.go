package utils

import (
	"fmt"
	"context"
	redis "github.com/0187773933/RedisManagerUtils/manager"
)

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB int `json:"db"`
	Password string `json:"password"`
	Prefix string `json:"prefix"`
}

func GetRedisConnection( config *RedisConfig ) ( redis_client redis.Manager ) {
	redis_client = redis.Manager{}
	redis_client.Connect(
		fmt.Sprintf( "%s:%s" , config.Host , config.Port ) ,
		config.DB ,
		config.Password ,
	)
	return
}

func RedisKeyExists( redis *redis.Manager , redis_key string ) ( result bool ) {
	result = false
	var ctx = context.Background()
	exists_result , exists_error := redis.Redis.Exists( ctx , redis_key ).Result()
	if exists_error != nil { fmt.Println( exists_error ); }
	if exists_result == 1 { result = true }
	return
}