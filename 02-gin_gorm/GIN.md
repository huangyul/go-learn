# gin入门

基本要点：
- 定义路由：包括路由参数、通配符路由
- 处理输入输出
- 使用middleware解决AOP问题

### 最简单的实例

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 相当于server
	r := gin.Default()
	// 注册路由
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello, gin")
	})
	// 监听端口
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

```

## 路由

- 静态路由：就是完全匹配的
- 参数路由：路由上带有参数
- 通配符路由


