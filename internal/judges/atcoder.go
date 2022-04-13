package judges

import (
	"strings"
)

type AtCoder struct{}

func (judge AtCoder) Id() string {
	return "atcoder"
}

func (judge AtCoder) ContestAndTaskId(url string) (string, string, error) {
	if strings.HasPrefix(url, "https://atcoder.jp/contests/") {
		// example url: "https://atcoder.jp/contests/abc246/tasks/abc246_a"
		urlSplits := strings.Split(url, "/")
		contestId := urlSplits[4]
		taskId := urlSplits[6][len(urlSplits[6])-1:]
		return contestId, taskId, nil
	} else {
		return "", "", ParseContestAndTaskIdError{url}
	}
}
