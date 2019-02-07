package dao

import (
	redis "gopkg.in/redis.v5"
)

// GetDao method
func GetDao() (DataDao, error) {
	clientStore := initRedis()
	return NewDataRedis(clientStore), nil
}

func initRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		panic(err)
	}
	return redisCli
}
