package judges

import "testing"

func TestAtCoderContestAndTaskId(t *testing.T) {
	urls := []string{
		`https://atcoder.jp/contests/abc246/tasks/abc246_a`,
		`https://www.atcoder.jp/contests/abc246/tasks/abc246_a`,
	}
	xContestId, xTaskId := "abc246", "a"
	RunTestContestAndTaskId(t, urls, AtCoder{}, xContestId, xTaskId)
}
