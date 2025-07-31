/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 20:13:36
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 20:26:23
 * @FilePath: /33-go-zero/grpc-client/main.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"context"
	"grpc-server/greet"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := greet.NewGreetClient(conn.Conn())
	resp, err := client.Ping(context.Background(), &greet.Request{Ping: "ping"})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp)
}
