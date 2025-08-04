package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	//Cookie 是浏览器和服务器之间自动传递的小型数据，常用于会话和身份管理。
	//在 Gin 项目中，Cookie 用于保存和加密 session 数据，实现用户登录等功能。
	//只需理解：Cookie 让你的服务能“记住”用户是谁，实现登录、个性化等功能。
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, "success")
		})
		// 用户服务
		v1.POST("user/register", handlers.UserRegister)
		v1.POST("user/login", handlers.UserLogin)

		// 需要登陆保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", handlers.GetTaskList)
			authed.POST("task", handlers.CreatTask)
			authed.GET("task/:id", handlers.GetTaskDetail) // task_id
			authed.PUT("task/:id", handlers.UpdateTask)    // task_id
			authed.DELETE("task/:id", handlers.DeleteTask) // task_id
		}
	}
	return ginRouter
}
