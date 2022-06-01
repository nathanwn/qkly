package judges

import (
	"fmt"
	"regexp"
)

type Judge interface {
	Id() string
	ContestAndTaskId(url string) (string, string, error)
}

func GetJudge(url string) (Judge, error) {
	if matched, _ := regexp.MatchString("https://(www.)?codeforces.com/*", url); matched {
		return CodeForces{}, nil
	} else if matched, _ := regexp.MatchString("https://(www.)?atcoder.jp/*", url); matched {
		return AtCoder{}, nil
	} else if matched, _ := regexp.MatchString("https://(www.)?hackerrank.com/*", url); matched {
		return HackerRank{}, nil
	} else if matched, _ := regexp.MatchString("https://(www.)?cses.fi/*", url); matched {
		return CSES{}, nil
	} else if matched, _ := regexp.MatchString("https://(www.)?spoj.com/*", url); matched {
		return SPOJ{}, nil
	} else {
		return nil, UnsupportedJudge{url}
	}
}

type ParseContestAndTaskIdError struct {
	url string
}

func (err ParseContestAndTaskIdError) Error() string {
	return fmt.Sprintf("Could not parse contest id and task id from url \"%s\"", err.url)
}
