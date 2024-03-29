package tasks

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nathan-wien/qkly/internal/filesystem"
	"github.com/nathan-wien/qkly/internal/judges"
	"github.com/nathan-wien/qkly/internal/schema"
)

type Judge judges.Judge

type Task struct {
	Judge       Judge
	Data        *schema.TaskData
	ContestId   string
	TaskId      string
	FsMgr       *filesystem.FileSystemManager
	ProjectDir  string
	TemplateDir string
}

var DefaultTemplateDir = filepath.Join("templates", "default")

func NewTask(taskData *schema.TaskData, fsMgr *filesystem.FileSystemManager, projectDir string, templateDir string) (*Task, error) {
	judge, err := judges.GetJudge(taskData.Url)
	if err != nil {
		return nil, err
	}
	contestId, taskId, err := judge.ContestAndTaskId(taskData.Url)
	if err != nil {
		return nil, err
	}
	if templateDir == "" {
		templateDir = DefaultTemplateDir
	}
	return &Task{judge, taskData, contestId, taskId, fsMgr, projectDir, templateDir}, nil
}

func (task Task) CreateTask() error {
	if err := task.writeTestFiles(); err != nil {
		return err
	}
	if err := task.copyTemplateFiles(); err != nil {
		return err
	}
	return nil
}

func (task Task) Dir() string {
	return filepath.Join(task.ProjectDir, "solutions", task.Judge.Id(), task.ContestId, task.TaskId)
}

func (task Task) CreateAndWriteFile(filepath string, content string) error {
	f, err := task.FsMgr.Fs.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}

	return nil
}

func (task Task) writeTestFiles() error {
	task.FsMgr.Fs.MkdirAll(task.Dir(), os.ModePerm)

	for id, testData := range task.Data.Tests {
		testId := id + 1
		inFilePath := filepath.Join(task.Dir(), fmt.Sprintf("%v.in.txt", testId))
		outFilePath := filepath.Join(task.Dir(), fmt.Sprintf("%v.out.txt", testId))

		if err := task.CreateAndWriteFile(inFilePath, testData.Input); err != nil {
			return err
		}
		if err := task.CreateAndWriteFile(outFilePath, testData.Output); err != nil {
			return err
		}
	}
	return nil
}

func (task Task) copyTemplateFiles() error {
	// Example: https://github.com/moby/moby/blob/63a9e3cc93efabb3759412a11a41a8c07a6e95e5/daemon/graphdriver/copy/copy.go#L117
	var walkFunc filepath.WalkFunc = func(path string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		copySubPath := path[strings.Index(path, task.TemplateDir)+len(task.TemplateDir):]
		copyPath := filepath.Join(task.Dir(), copySubPath)
		switch mode := fileInfo.Mode(); {
		case mode.IsDir():
			if err := task.FsMgr.Util.Mkdir(copyPath, fileInfo.Mode()); err != nil && !os.IsExist(err) {
				return err
			}
		default:
			content, err := task.FsMgr.Util.ReadFile(path)
			if err != nil {
				return err
			}
			if err := task.CreateAndWriteFile(copyPath, string(content)); err != nil {
				return err
			}
		}
		return nil
	}
	if err := task.FsMgr.Util.Walk(task.TemplateDir, walkFunc); err != nil {
		return err
	}
	return nil
}
