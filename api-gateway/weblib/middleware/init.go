package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitMiddleware 接受服务实例，并存在gin.key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存入gin.keys中
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Keys["taskService"] = service[1]
		context.Next()
	}
}

// ErrorMiddleware 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			// 通过 recover() 捕获异常，如果有异常发生，则返回一个 JSON 响应，包含错误信息
			if r := recover(); r != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": http.StatusNotFound,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
