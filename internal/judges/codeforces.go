package judges

import (
	"strings"
)

type CodeForces struct{}

func (judge CodeForces) Id() string {
	return "codeforces"
}

func (judge CodeForces) ContestAndTaskId(url string) (string, string, error) {
	if strings.HasPrefix(url, "https://codeforces.com/problemset/") {
		// example url: "https://codeforces.com/problemset/problem/348/D"
		urlSplits := strings.Split(url, "/")
		contestId := urlSplits[5]
		taskId := strings.ToLower(urlSplits[6])
		return contestId, taskId, nil
	} else if strings.HasPrefix(url, "https://codeforces.com/contest/") {
		// example url: "https://codeforces.com/contest/1657/problem/A"
		urlSplits := strings.Split(url, "/")
		contestId := urlSplits[4]
		taskId := strings.ToLower(urlSplits[6])
		return contestId, taskId, nil
	} else {
		return "", "", ParseContestAndTaskIdError{url}
	}
}
