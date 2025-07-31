/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 19:08:54
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:09:02
 * @FilePath: /33-go-zero/user-api/internal/db/cache.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package db

import "github.com/zeromicro/go-zero/core/stores/redis"

func NewRedis(conf redis.RedisConf) *redis.Redis {
	return redis.MustNewRedis(conf)
}
