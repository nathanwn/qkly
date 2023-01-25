package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nathan-wien/qkly/internal/filesystem"
	"github.com/nathan-wien/qkly/internal/schema"
	"github.com/nathan-wien/qkly/internal/tasks"
)

func HandleCompetitiveCompanion(context *gin.Context) {
	var taskData schema.TaskData
	if err := context.BindJSON(&taskData); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	task, err := tasks.NewTask(&taskData, filesystem.Mgr(), ".", "")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	if err := task.CreateTask(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("Parsed task to %s\n", task.Dir())
}
