package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

// DB 定义全局数据库变量
var (
	DB *gorm.DB
)

const (
	user     = "root"
	password = "149661377.."
	addr     = "127.0.0.1"
	port     = "3306"
	database = "ToDoList"
)

// InitMySQL :初始化数据库
func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, port, database)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close() //程序退出关闭数据库连接
}
