package judges

import (
	"regexp"
	"strings"
)

type SPOJ struct{}

func (judge SPOJ) Id() string {
	return "spoj"
}

func (judge SPOJ) ContestAndTaskId(url string) (string, string, error) {
	if matched, _ := regexp.MatchString("https://(www.)?spoj.com/*", url); matched {
		// example url: "https://www.spoj.com/problems/LCS0/"
		urlSplits := strings.Split(url, "/")
		contestId := ""
		taskId := urlSplits[4]
		return contestId, taskId, nil
	} else {
		return "", "", ParseContestAndTaskIdError{url}
	}
}
