package router

import (
	"ToDoList/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	//启动路由
	r := gin.Default()
	//告诉gin框架模版文件引用的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	ListGroup := r.Group("/v1")
	{
		//添加
		ListGroup.POST("/todo", controller.AddToDo)
		//查看所有待办事项
		ListGroup.GET("/todo", controller.GetAllToDos)
		//查看某一个待办事项
		ListGroup.GET("/todo/:id", controller.GetToDo)
		//修改某个待办事项
		ListGroup.PUT("/todo/:id", controller.UpdateToDo)
		//删除某个待办事项
		ListGroup.DELETE("/todo/:id", controller.DeleteToDo)
	}
	return r
}
