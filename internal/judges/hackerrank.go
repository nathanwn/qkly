package judges

import (
	"regexp"
	"strings"
)

type HackerRank struct{}

func (judge HackerRank) Id() string {
	return "hackerrank"
}

func (judge HackerRank) ContestAndTaskId(url string) (string, string, error) {
	var contestId, taskId string
	// example: https://www.hackerrank.com/contests/uqcs-codejam-2021/challenges/lucky-code
	if matched, _ := regexp.MatchString("https://(www.)?hackerrank.com/*", url); matched {
		urlSplits := strings.Split(url, "/")
		contestId = urlSplits[4]
		taskId = urlSplits[6]
		return contestId, taskId, nil
	} else {
		return "", "", nil
	}
}
