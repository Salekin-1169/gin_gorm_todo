package main

import (
	"fmt"
	"gin-web/ctrl"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")

	{
		v1.POST("/", ctrl.CreateTodo)
		v1.GET("/", ctrl.FetchAllTodo)
		v1.GET("/:id", ctrl.FetchSingleTodo)
		v1.PUT("/:id", ctrl.UpdateTodo)
		v1.DELETE("/:id", ctrl.DeleteTodo)
	}

	if err := router.Run(":9090"); err != nil {
		fmt.Println("error:", err.Error())
	}
}
