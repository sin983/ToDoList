package controller

import (
	dao "ToDoList/dao"
	"ToDoList/moudle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddToDo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 会发送请求到这里
	//1.从请求中把数据拿出
	var todo moudle.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	//2.存入数据库
	if err := dao.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetAllToDos(c *gin.Context) {
	//查询todo表里的所有数据
	var todoList []moudle.Todo
	if err := dao.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetToDo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	var todo moudle.Todo
	if err := dao.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateToDo(c *gin.Context) {
	//从请求中把id取出
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error:": "无效id"})
		return
	}
	var todo moudle.Todo
	//根据id查询数据库中是否有此数据
	if err := dao.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	//将更改后的数据传给todo{}
	c.ShouldBind(&todo)
	//将todo{}中的数据更新到数据库
	if err := dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteToDo(c *gin.Context) {
	//从请求中把id取出
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error:": "无效id"})
		return
	}
	//根据id字段删除数据
	if err := dao.DB.Delete(moudle.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
