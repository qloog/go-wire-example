package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app: %v", err))
	}

	// 初始化 Gin 引擎
	router := gin.Default()
	// 注册 UserHandler 的路由
	v1 := router.Group("/v1")
	{
		v1.POST("/login", app.UserHandler.LoginHandler)
		v1.GET("/products", app.ProductHandler.GetProducts)
		// 添加更多路由
	}

	// 启动服务
	if err := router.Run(":8080"); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
