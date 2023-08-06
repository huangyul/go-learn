package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 静态路由
	r.GET("/hello", func(ctx *gin.Context) {
		// 获取查询参的方法
		ctx.String(http.StatusOK, "hello, gin")
	})

	// 参数路由
	r.GET("/users/:name", func(ctx *gin.Context) {
		// 获取参数的方法
		p := ctx.Param("name")
		ctx.String(http.StatusOK, fmt.Sprintf("参数路由：%s", p))
	})

	// 匹配路由
	r.GET("/order/*.html", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "匹配路由")
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}

}
