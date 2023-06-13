package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "pong"})
	//})
	//r.Run()

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // MaxHeaderBytes控制服务器在解析请求头的键和值（包括请求行）时读取的最大字节数。它不限制请求正文的大小。
	}

	s.ListenAndServe()
}
