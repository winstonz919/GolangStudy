/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 13:53:16
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 13:53:21
 * @FilePath: /33-go-zero/user-api/internal/biz/result.go
 * @Description: 返回结果
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package biz

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{
		Code: Ok,
		Msg:  "success",
		Data: data,
	}
}

func Fail(err *Error) *Result {
	return &Result{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
