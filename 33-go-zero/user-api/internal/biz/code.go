/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 13:49:41
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 19:24:03
 * @FilePath: /33-go-zero/user-api/internal/biz/code.go
 * @Description: 业务错误码
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package biz

const Ok = 200

var (
	DBError         = NewError(10000, "数据库错误")
	CacheError      = NewError(10002, "缓存错误")
	AlreadyRegister = NewError(10100, "用户已注册")
	UserNotFound    = NewError(10101, "用户名或密码错误")
	TokenError      = NewError(10102, "token获取失败")
	InvalidToken    = NewError(10103, "非法token")
)
