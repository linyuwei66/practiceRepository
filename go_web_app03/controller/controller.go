package controller

import (
	"go_web_app03/models"

	"github.com/gin-gonic/gin"

)

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetTodoList(c *gin.Context) {
	//final version
	// var todos []Todo
	// err = DB.Find(&todos).Error
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"error": err.Error(),
	// 	})
	// } else {
	// 	c.JSON(200, todos)
	// }

	//my version
	todos,err := models.GetTodoListInModel()
	if err!=nil{
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(200, todos)
	}

	//teacher's version
	// var todoList []Todo
	// 	if err = DB.Find(&todoList).Error; err!= nil {
	// 		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 	}else {
	// 		c.JSON(http.StatusOK, todoList)
	// 	}
}

func CreateATodo(c *gin.Context) {
	//final version
	var todo models.Todo
	c.BindJSON(&todo)
	err:=models.CreateATodoInModel(&todo)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, todo)
	}

	//my version
	// var todo Todo
	// c.BindJSON(&todo)
	// DB.Create(&todo)
	// c.JSON(200, todo)

	//teacher's version
	// 	var todo Todo
	// 	c.BindJSON(&todo)
	// 	//err = DB.Create(&todo).Error
	// 	//if err!= nil {
	// 	//}
	// 	if err = DB.Create(&todo).Error;err != nil {
	// 		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 	}else{
	// 		c.JSON(http.StatusOK, todo)
	// 	}
}

func UpdateATodo(c *gin.Context) {
	//final version
	// anum, ok := c.Params.Get("id")
	// if !ok {
	// 	c.JSON(200, gin.H{
	// 		"error": "invalid id",
	// 	})
	// 	return
	// }
	// var todo Todo
	// err := DB.First(&todo, anum).Error
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	// c.BindJSON(&todo)
	// err = DB.Save(&todo).Error
	// if err != nil {
	// 	c.JSON(200, gin.H{
	// 		"error": err.Error(),
	// 	})
	// } else {
	// 	c.JSON(200, todo)
	// }

	//my version
	anum,ok := c.Params.Get("id")
	if !ok{
		c.JSON(200,gin.H{
			"error":"无效的id",
		})
		return
	}
	todo,err:=models.GetATodoInModel(anum)
	if err!=nil{
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	}

	c.BindJSON(&todo)
	err=models.UpdateATodoInModel(todo)
	if err!=nil{
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	}
	//c.JSON(200, todo)

	//teacher's version
	// id, ok := c.Params.Get("id")
	// 	if !ok {
	// 		c.JSON(200, gin.H{"error": "无效的id"})
	// 		return
	// 	}
	// 	var todo Todo
	// 	if err = db.Where("id=?", id).First(&todo).Error; err!=nil{
	// 		c.JSON(200, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.BindJSON(&todo)
	// 	if err = db.Save(&todo).Error; err!= nil{
	// 		c.JSON(200, gin.H{"error": err.Error()})
	// 	}else{
	// 		c.JSON(200, todo)
	// 	}
}

func DeleteATodo(c *gin.Context) {
	//final version
	anum, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{
			"error": "invalid id",
		})
	}
	err := models.DeleteATodoInModel(anum)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"id": "deleted",
		})
	}

	//my version
	// anum := c.Param("id")
	// DB.Where("id=?", anum).Delete(Todo{})

	//teacher's version
	// id, ok := c.Params.Get("id")
	// if !ok {
	// 	c.JSON(200, gin.H{"error": "无效的id"})
	// 	return
	// }
	// if err = db.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
	// 	c.JSON(200, gin.H{"error": err.Error()})
	// }else{
	// 	c.JSON(200, gin.H{id:"deleted"})
	// }
}