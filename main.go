package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由器
	r := gin.Default()

	// 设置一个简单的 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 设置一个 POST 路由来接收 JSON 数据
	r.POST("/user", func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": json.Name,
			"age":  json.Age,
		})
	})

	// 启动 HTTP 服务，监听默认端口8080
	r.Run("localhost:8880")
}
