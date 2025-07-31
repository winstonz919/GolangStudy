/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 11:59:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 14:07:56
 * @FilePath: /33-go-zero/user-api/internal/logic/account/registerlogic.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package account

import (
	"context"
	"time"

	"user-api/internal/biz"
	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUserModel(l.svcCtx.Conn)
	user, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("find user err:", err)
		return nil, biz.DBError
	}

	if user != nil {
		return nil, biz.AlreadyRegister
	}

	_, err = userModel.Insert(l.ctx, &model.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return
}
