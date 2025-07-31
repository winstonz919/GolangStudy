/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 19:44:12
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:53:18
 * @FilePath: /33-go-zero/grpc-server/internal/logic/pinglogic.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package logic

import (
	"context"

	"grpc-server/greet"
	"grpc-server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	// todo: add your logic here and delete this line

	return &greet.Response{Pong: "pong"}, nil
}
