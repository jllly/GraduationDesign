package Middleware

import (
	"Homework_system/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

//是否登录的验证中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authkey, err := c.Cookie("auth")
		//，没cookie会报错
		if err != nil {
			fmt.Printf("%v", err)
			c.Header("Pragma", "no-cache")
			c.Header("Cache-Control", "no-cache")
			c.Redirect(307, "/") //没有cookie重定向到登录页
			return
		}
		//在有cookie的情况下，查看是否有session
		_, ok := global.Session[authkey]
		if !ok {
			c.Header("Pragma", "no-cache")
			c.Header("Cache-Control", "no-cache")
			c.Redirect(307, "/")
			return
		}
		// 调用业务代码
		c.Next()
	}
}
