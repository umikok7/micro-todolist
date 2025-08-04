package model

/*
功能： 将模型映射到数据库上
*/

func migration() {
	// 通过 gorm 的 AutoMigrate 方法，将 Go 语言中的结构体（如 User）映射为数据库表。
	// 使用 Set 方法设置表的选项，例如字符集为 utf8mb4，以支持存储多语言字符
	DB.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&Task{})
}
