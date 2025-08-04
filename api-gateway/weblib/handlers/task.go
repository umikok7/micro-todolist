package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTaskList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到的是当前访问的用户的id，拿到用户自己的备忘录信息
	taskReq.Uid = uint64(claim.Id)
	// 调用服务端函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data":  taskResp.TaskList,
		"count": taskResp.Count,
	})
}

func CreatTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到的是当前访问的用户的id，拿到用户自己的备忘录信息
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	// 调用服务端函数
	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	//if err != nil {
	//	PanicIfTaskError(err)
	//}
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到的是当前访问的用户的id，拿到用户自己的备忘录信息
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id")) //获取task_id，前面传进来的那个
	taskReq.Id = uint64(id)
	productService := ginCtx.Keys["taskService"].(services.TaskService)
	// 调用服务端函数
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": productRes.TaskDetail,
	})
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到的是当前访问的用户的id，拿到用户自己的备忘录信息
	id, _ := strconv.Atoi(ginCtx.Param("id"))                       //获取task_id，前面传进来的那个
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	// 调用服务端函数
	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}

func DeleteTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization")) // 拿到的是当前访问的用户的id，拿到用户自己的备忘录信息
	id, _ := strconv.Atoi(ginCtx.Param("id"))                       //获取task_id，前面传进来的那个
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	// 调用服务端函数
	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": taskResp.TaskDetail,
	})
}
