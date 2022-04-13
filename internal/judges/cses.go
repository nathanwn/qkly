package judges

import (
	"strings"
)

type CSES struct{}

func (judge CSES) Id() string {
	return "cses"
}

func (judge CSES) ContestAndTaskId(url string) (string, string, error) {
	if strings.HasPrefix(url, "https://www.cses.fi/problemset/") {
		// example url: "https://www.cses.fi/problemset/task/1068"
		urlSplits := strings.Split(url, "/")
		contestId := ""
		taskId := urlSplits[5]
		return contestId, taskId, nil
	} else {
		return "", "", ParseContestAndTaskIdError{url}
	}
}
