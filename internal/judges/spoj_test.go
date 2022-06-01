package judges

import "testing"

func TestSPOJContestAndTaskId(t *testing.T) {
	urls := []string{
		`https://www.spoj.com/problems/LCS0`,
		`https://www.spoj.com/problems/LCS0/`,
		`https://spoj.com/problems/LCS0`,
		`https://spoj.com/problems/LCS0/`,
	}
	xContestId, xTaskId := "", "LCS0"
	RunTestContestAndTaskId(t, urls, SPOJ{}, xContestId, xTaskId)
}
