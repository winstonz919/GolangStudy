/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 12:28:34
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 12:29:06
 * @FilePath: /33-go-zero/user-api/internal/db/mysql.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package db

import (
	"context"
	"time"
	"user-api/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)
	db, err := mysql.RawDB()
	if err != nil {
		panic(err)
	}
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mysqlConfig.ConnectTimeout))
	defer cancel()
	err = db.PingContext(cxt)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return mysql
}
