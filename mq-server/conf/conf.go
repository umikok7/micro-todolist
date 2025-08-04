package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mq-server/model"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	LoadMysqlData(file)
	pathMySQL := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=" + Charset + "&parseTime=true"}, "")
	model.Database(pathMySQL) // 连接数据库

	LoadRabbitMQ(file)
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	model.RabbitMQ(pathRabbitMQ) // 连接MQ
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}
