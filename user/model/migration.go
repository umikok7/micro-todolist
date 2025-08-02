package model

/*
功能： 将模型映射到数据库上
*/

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&User{})
}
