package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}