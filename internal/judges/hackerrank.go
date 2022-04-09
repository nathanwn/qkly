package judges

import (
	"strings"
)

type HackerRank struct{}

func (judge HackerRank) Id() string {
	return "hackerrank"
}

func (judge HackerRank) ContestAndTaskId(url string) (string, string) {
	var contestId, taskId string
	// example: https://hackerrank.com/contests/uqcs-codejam-2021/challenges/lucky-code
	if strings.HasPrefix(url, "https://hackerrank.com/contests/") {
		urlSplits := strings.Split(url, "/")
		contestId = urlSplits[4]
		taskId = urlSplits[6]
	}
	return contestId, taskId
}
