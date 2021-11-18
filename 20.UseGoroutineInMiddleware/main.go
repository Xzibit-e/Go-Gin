package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// 用time.Sleep()模拟一个长任务
			time.Sleep(5 * time.Second)

			// 请注意使用的是复制的上下文"cCp",这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用time.Sleep()模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// 监听并在127.0.0.1:8080上启动服务
	r.Run(":8080")
}