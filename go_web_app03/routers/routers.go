package routers

import (
	"github.com/gin-gonic/gin"
	"go_web_app03/controller"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	inthegroup := r.Group("/v1")
	inthegroup.GET("/todo", controller.GetTodoList)
	inthegroup.POST("/todo", controller.CreateATodo)
	inthegroup.PUT("/todo/:id", controller.UpdateATodo)
	inthegroup.DELETE("/todo/:id", controller.DeleteATodo)
	return r
}