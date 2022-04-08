package judges

import (
	"strings"
)

func GetJudge(url string) Judge {
	if strings.HasPrefix(url, "https://codeforces.com/") {
		return CodeForces{}
	} else if strings.HasPrefix(url, "https://atcoder.jp/") {
		return AtCoder{}
	} else if strings.HasPrefix(url, "https://hackerrank.com/") {
		return HackerRank{}
	} else {
		return nil
	}
}
