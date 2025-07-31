/*
 * @Author: WinstonRD
 * @Date: 2025-06-30 20:13:47
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-06-30 20:24:08
 * @FilePath: /27-gin/middles/middles.go
 * @Description: 中间件
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package middles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// token验证
func CheckTokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token != "admin123456" {
			c.JSON(http.StatusForbidden, gin.H{"msg": "请先登录"})
			c.Abort()
		}
		c.Next()
	}
}
