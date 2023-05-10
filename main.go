package main

import (
	dao "ToDoList/dao"
	"ToDoList/moudle"
	"ToDoList/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	//启动数据库
	if err := dao.InitMySQL(); err != nil {
		log.Fatal(err)
	}
	defer dao.Close()
	//连接数据库与结构体模型
	dao.DB.AutoMigrate(&moudle.Todo{})
	r := router.InitRouter()
	//运行
	log.Fatal(r.Run(":8080"))
}
