/*
 * @Author: WinstonRD
 * @Date: 2025-07-15 11:59:05
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-15 14:12:22
 * @FilePath: /33-go-zero/user-api/internal/handler/account/registerhandler.go
 * @Description:
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package account

import (
	"net/http"

	"user-api/internal/biz"
	"user-api/internal/logic/account"
	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			httpx.OkJsonCtx(r.Context(), w, biz.Success(resp))
		}
	}
}
