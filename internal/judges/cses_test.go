package judges

import "testing"

func TestCSESContestAndTaskId(t *testing.T) {
	urls := []string{
		`https://cses.fi/problemset/task/1068`,
		`https://www.cses.fi/problemset/task/1068`,
	}
	xContestId, xTaskId := "", "1068"
	RunTestContestAndTaskId(t, urls, CSES{}, xContestId, xTaskId)
}
