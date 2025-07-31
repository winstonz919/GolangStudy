/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 11:59:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:24:13
 * @FilePath: /33-go-zero/user-api/internal/logic/account/loginlogic.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package account

import (
	"context"
	"fmt"
	"time"

	"user-api/internal/biz"
	"user-api/internal/model"
	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 校验用户名密码
	userModel := model.NewUserModel(l.svcCtx.Conn)
	user, err := userModel.FindByUsernameAndPassword(l.ctx, req.Username, req.Password)
	if err != nil {
		return nil, biz.DBError
	}
	if user == nil {
		return nil, biz.UserNotFound
	}
	// 生成token
	secret := l.svcCtx.Config.JwtAuth.Secret
	expire := l.svcCtx.Config.JwtAuth.Expire
	token, err := biz.GetJwtToken(secret, time.Now().Unix(), int64(expire), user.Id)
	if err != nil {
		return nil, biz.TokenError
	}
	// 写入缓存
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, "token:"+token, fmt.Sprintf("%d", user.Id), expire)
	if err != nil {
		return nil, biz.CacheError
	}
	return &types.LoginResp{Token: token}, nil
}
