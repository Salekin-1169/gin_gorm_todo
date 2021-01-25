package ctrl

import (
	"gin-web/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {

	completed, err := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{Title: c.PostForm("title")}
	if err == nil {
		todo.Completed = completed
	}

	model.DBEngin.Create(&todo)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo item created successfully",
		"data": gin.H{
			"todoId": todo.ID,
		},
	})
}

func FetchAllTodo(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TodoModelGetData

	model.DBEngin.Find(&todos)

	if len(todos) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "nothing todo",
		})
	}

	for _, todo := range todos {
		var completed bool
		if todo.Completed == 0 {
			completed = false
		} else {
			completed = true
		}
		_todos = append(_todos, model.TodoModelGetData{ID: todo.ID, Title: todo.Title, Completed: completed})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

func FetchSingleTodo(c *gin.Context) {
	var todo model.TodoModel

	todoId := c.Param("id")

	model.DBEngin.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "not found",
		})
		return
	}

	var _todo = &model.TodoModelGetData{ID: todo.ID, Title: todo.Title}
	if todo.Completed == 0 {
		_todo.Completed = false
	} else {
		_todo.Completed = true
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todo,
	})
}

func UpdateTodo(c *gin.Context) {
	var todo model.TodoModel

	todoId := c.Param("id")

	model.DBEngin.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Change failed, Invalid todo",
		})
		return
	}

	_title, _completed := c.PostForm("title"), c.PostForm("completed")

	if len(_title) == 0 && len(_completed) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please change some value",
		})
	}

	if len(_title) > 0 {
		todo.Title = _title
	}

	if len(_completed) > 0 {
		_completed, err := strconv.Atoi(_completed)
		if err == nil {
			todo.Completed = _completed
		}
	}

	model.DBEngin.Save(&todo)

	var _todo = &model.TodoModelGetData{ID: todo.ID, Title: todo.Title}
	if todo.Completed == 0 {
		_todo.Completed = false
	} else {
		_todo.Completed = true
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todo,
	})
}

func DeleteTodo(c *gin.Context) {
	var todo model.TodoModel

	todoId := c.Param("id")

	model.DBEngin.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Change failed, Invalid todo",
		})
		return
	}

	model.DBEngin.Unscoped().Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "todo delete successfully",
	})
}
