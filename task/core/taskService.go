package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"task/model"
	"task/services"
)

// CreateTask 创建备忘录，将备忘录信息生产， 放在RabbitMQ消息队列中
func (*TaskService) CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("RabbitMQ channel err:" + err.Error())
	}
	q, err := ch.QueueDeclare("task_name", true, false, false, false, nil)
	body, _ := json.Marshal(req) // 序列化：例如 title content
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json", // 传输的类型
		Body:         body,               // 发送的消息
	})
	if err != nil {
		err = errors.New("RabbitMQ Publish err:" + err.Error())
	}
	return nil
}

// GetTasksList 实现备忘录服务接口 获取备忘录列表
func (*TaskService) GetTasksList(ctx context.Context, req *services.TaskRequest, resp *services.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count uint32
	// 查找备忘录
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find:" + err.Error())
	}
	// 统计数量
	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	var taskRes []*services.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = count

	return nil
}

// GetTask 获取备忘录的详细信息
func (*TaskService) GetTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

// UpdateTask 修改备忘录
func (*TaskService) UpdateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).First(&taskData)
	// 执行对应的修改
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content

	model.DB.Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

// DeleteTask 删除备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Delete(&model.Task{})
	return nil
}
