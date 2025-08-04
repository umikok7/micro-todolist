package service

import (
	"encoding/json"
	"log"
	"mq-server/model"
)

// CreateTask 从RabbitMQ中接收消息，写入数据库中
func CreateTask() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, err := ch.QueueDeclare("task_name", true, false, false, false, nil)
	err = ch.Qos(1, 0, false) // 控制速度
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 处于持续监听状态，一直监听生产者的生产，所以我们要阻塞主进程
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()
}
