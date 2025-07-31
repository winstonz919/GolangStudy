/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 11:59:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:16:33
 * @FilePath: /33-go-zero/user-api/internal/config/config.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package config

import (
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	RedisConfig redis.RedisConf
	JwtAuth     JwtAuth
}

type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int
}

type JwtAuth struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}

type RedisConfig struct {
	Host     string
	Type     string `json:",default=node,options=node|cluster"`
	Pass     string `json:",optional"`
	Tls      bool   `json:",optional"`
	NonBlock bool   `json:",default=true"`
	// PingTimeout is the timeout for ping redis.
	PingTimeout time.Duration `json:",default=1s"`
}
