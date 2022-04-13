package judges

import "testing"

func TestHackerRankContestAndTaskId(t *testing.T) {
	urls := []string{
		`https://hackerrank.com/contests/uqcs-codejam-2021/challenges/lucky-code`,
		`https://www.hackerrank.com/contests/uqcs-codejam-2021/challenges/lucky-code`,
	}
	xContestId, xTaskId := "uqcs-codejam-2021", "lucky-code"
	RunTestContestAndTaskId(t, urls, HackerRank{}, xContestId, xTaskId)
}
