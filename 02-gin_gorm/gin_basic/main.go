package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello, gin")
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
