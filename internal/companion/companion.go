package companion

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nathan-wien/qkly/internal/filesystem"
	"github.com/nathan-wien/qkly/internal/tasks"
)

func Handle(context *gin.Context) {
	var taskData tasks.TaskData
	if err := context.BindJSON(&taskData); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	task := tasks.NewTask(&taskData, filesystem.Mgr(), ".", "")
	if err := task.CreateTask(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Parsed task to %s\n", task.Dir())
	}
}
