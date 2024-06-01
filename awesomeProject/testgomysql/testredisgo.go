package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var ctx = context.Background()

func main() {
	// 连接到本地Redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 没有密码则留空
		DB:       0,                // 默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully!")
	}

	// 设置Gin路由
	r := gin.Default()

	r.POST("/users/signup", func(c *gin.Context) {
		// 你的处理逻辑
		c.JSON(http.StatusOK, gin.H{"message": "signup successful"})
	})

	// 启动Gin服务器
	r.Run()
}
