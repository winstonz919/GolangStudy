/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 15:21:11
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 15:47:30
 * @FilePath: /33-go-zero/user-api/internal/logic/account/userinfologic.go
 * @Description: 用户信息逻辑层
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package account

import (
	"context"
	"encoding/json"

	"user-api/internal/biz"
	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	// 获取userId
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, biz.InvalidToken
	}
	userModel := model.NewUserModel(l.svcCtx.Conn)
	user, err := userModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, biz.DBError
	}
	if user == nil {
		return nil, biz.UserNotFound
	}

	return &types.UserInfoResp{UserId: int(user.Id), Username: user.Username}, nil
}
