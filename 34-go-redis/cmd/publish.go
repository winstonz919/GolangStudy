/*
 * @Author: WinstonRD
 * @Date: 2025-07-19 23:55:35
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-20 00:03:46
 * @FilePath: /34-go-redis/cmd/publish.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"34-redis-go/internal/config"
	"34-redis-go/internal/db"
	"context"
)

func main() {
	conf := config.NewConfig("../internal/etc/config.yaml")
	ctx := context.WithValue(context.Background(), config.CtxKey("conf"), conf)
	r := db.NewCache(ctx)
	r.Publish(ctx, "channel_1", "message_1")
}
