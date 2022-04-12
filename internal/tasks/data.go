package tasks

import (
	"strings"

	"github.com/nathan-wien/qkly/internal/judges"
)

type InputType string

const (
	InputKindStdin = "stdin"
	InputKindFile  = "file"
	InputKindRegex = "regex"
)

type InputConfiguration struct {
	Type     InputType `json:"type"`
	FileName string    `json:"fileName"`
	Pattern  string    `json:"pattern"`
}

type OutputType string

const (
	OutputKindStdout OutputType = "stdout"
	OutputKindFile              = "file"
)

type OutputConfiguration struct {
	Type     OutputType `json:"type"`
	FileName string     `json:"fileName"`
}

type TestFileType string

const (
	TestFileTypeInput  TestFileType = "input"
	TestFileTypeOutput TestFileType = "output"
)

type TestData struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type TaskData struct {
	Name         string              `json:"name"`
	Group        string              `json:"group"`
	Url          string              `json:"url"`
	Interactive  bool                `json:"interactive"`
	MemoryLimit  int                 `json:"memoryLimit"`
	TimeLimit    int                 `json:"timeLimit"`
	Tests        []TestData          `json:"tests"`
	TestType     string              `json:"testType"`
	InputConfig  InputConfiguration  `json:"input"`
	OutputConfig OutputConfiguration `json:"output"`
}

type UnsupportedJudge struct{}

func (err UnsupportedJudge) Error() string {
	return "Judge"
}

func GetJudge(taskData *TaskData) (judges.Judge, error) {
	if strings.HasPrefix(taskData.Url, "https://codeforces.com/") {
		return judges.CodeForces{}, nil
	} else if strings.HasPrefix(taskData.Url, "https://atcoder.jp/") {
		return judges.AtCoder{}, nil
	} else if strings.HasPrefix(taskData.Url, "https://hackerrank.com/") {
		return judges.HackerRank{}, nil
	} else if strings.HasPrefix(taskData.Url, "https://www.cses.fi/") {
		return judges.CSES{}, nil
	} else {
		return nil, UnsupportedJudge{}
	}
}
