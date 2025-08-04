package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq)) // 解析前端的请求，映射绑定到userReq结构体当中
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq) // 调用对应rpc服务
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": userResp,
	})
}

func UserLogin(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq)) // 解析前端的请求，映射绑定到userReq结构体当中
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
