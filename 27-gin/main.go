/*
 * @Author: WinstonRD
 * @Date: 2025-06-28 22:19:48
 * @LastEditors: WinstonRD && winstonz919@sina.com
 * @LastEditTime: 2025-07-25 10:21:26
 * @FilePath: /27-gin/main.go
 * @Description: 用gin实现restful风格的api
 *
 * Copyright (c) 2025 by WinstonRD, All Rights Reserved.
 */
package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var count int
var m sync.Mutex

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

type UserLevel int

type User struct {
	Name   string
	Age    int
	Level  UserLevel
	Errors []error
}

func UserListLogic(pageIndex, pageSize int) []User {
	ul := make([]User, 0)
	ul = append(ul, User{Name: "recker", Age: 1, Level: 10}, User{Name: "winston", Age: 20, Level: 12})
	return ul
}

func UserList(c *gin.Context) {
	users := UserListLogic(1, 10)
	c.JSON(http.StatusOK, gin.H{"users": users}) // 返回json格式数据
}

func Add(c *gin.Context) {
	count++
	time.Sleep(1 * time.Second)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1") // 包含接口的版本号
	// v1.Use(middles.CheckTokenMiddleWare()) // v1路由下的接口调用中间件
	{
		calc := v1.Group("/calc")
		{
			calc.GET("/add", Add)
		}
		user := v1.Group("/users") // uri只包含跟资源相关的名词
		{
			user.GET("", UserList) // 用get方法获取数据
		}
	}
	router.Run("127.0.0.1:8888")
}
