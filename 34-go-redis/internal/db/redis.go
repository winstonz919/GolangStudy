/*
 * @Author: WinstonRD
 * @Date: 2025-07-18 16:01:15
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-18 16:11:20
 * @FilePath: /34-go-redis/internal/db/redis.go
 * @Description: redis
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package db

import (
	"34-redis-go/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

func NewCache(ctx context.Context) (client *redis.Client) {
	conf := ctx.Value(config.CtxKey("conf")).(*config.Config)
	opt := redis.Options{
		Addr:     conf.RedisConf.Addr,
		Password: conf.RedisConf.Passd,
		DB:       conf.RedisConf.Db,
	}
	client = redis.NewClient(&opt)
	if client == nil {
		panic("redis 连接异常")
	}
	return
}
