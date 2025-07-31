/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 19:44:12
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:48:27
 * @FilePath: /33-go-zero/grpc-server/internal/svc/servicecontext.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package svc

import "grpc-server/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
