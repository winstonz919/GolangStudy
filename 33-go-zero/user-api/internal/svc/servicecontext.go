/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 11:59:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:17:02
 * @FilePath: /33-go-zero/user-api/internal/svc/servicecontext.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package svc

import (
	"user-api/internal/config"
	"user-api/internal/db"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	Conn        sqlx.SqlConn
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := db.NewMysql(c.MysqlConfig)
	redisClient := db.NewRedis(c.RedisConfig)
	return &ServiceContext{
		Config:      c,
		Conn:        sqlConn,
		RedisClient: redisClient,
	}
}
