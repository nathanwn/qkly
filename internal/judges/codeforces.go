package judges

import (
	"regexp"
	"strings"
)

type CodeForces struct{}

func (judge CodeForces) Id() string {
	return "codeforces"
}

func (judge CodeForces) ContestAndTaskId(url string) (string, string, error) {
	if matched, _ := regexp.MatchString("https://(www.)?codeforces.com/problemset/*", url); matched {
		// example url: "https://codeforces.com/problemset/problem/348/D"
		urlSplits := strings.Split(url, "/")
		contestId := urlSplits[5]
		taskId := strings.ToLower(urlSplits[6])
		return contestId, taskId, nil
	} else if matched, _ := regexp.MatchString("https://(www.)?codeforces.com/contest/*", url); matched {
		// example url: "https://codeforces.com/contest/1657/problem/A"
		urlSplits := strings.Split(url, "/")
		contestId := urlSplits[4]
		taskId := strings.ToLower(urlSplits[6])
		return contestId, taskId, nil
	} else {
		return "", "", ParseContestAndTaskIdError{url}
	}
}
