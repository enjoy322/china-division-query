package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

func RedisInit(conf Redis) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.Db,
		PoolSize:     120,
		MinIdleConns: 20,
	})
	err := conn.Ping(context.Background()).Err()
	if err != nil {
		return nil, errors.New("Redis启动失败，" + err.Error())
	}
	//使用0号数据库
	conn.Do(context.Background(), "select", conf.Db)
	return conn, nil
}

type Redis struct {
	Addr     string `json:"Addr"`
	Password string `json:"Password"`
	Db       int    `json:"Db"`
}
