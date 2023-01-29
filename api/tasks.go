package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context){
	var createTask service.CreateTaskService
	claim , _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask);err== nil{
		res := createTask.Create(claim.Id)
		c.JSON(200,res)
	}else{
		logging.Error(err)
		c.JSON(400,err)
	}
}

func ShowTask(c *gin.Context) {
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(200, res)
}

func ListTask(c *gin.Context) {
	listTask := service.ListTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(chaim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateTask(c *gin.Context){
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask);err== nil{
		res := updateTask.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		logging.Error(err)
		c.JSON(400,err)
	}
}

func SearchTask(c *gin.Context){
	searchTask := service.SearchTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteTask(c *gin.Context){
	deleteTask := service.DeleteTaskService{}
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}