package judges

import "testing"

func TestCodeForcesContestAndTaskIdWithProblemsetUrl(t *testing.T) {
	urls := []string{
		`https://codeforces.com/problemset/problem/348/D`,
		`https://www.codeforces.com/problemset/problem/348/D`,
	}
	xContestId, xTaskId := "348", "d"
	RunTestContestAndTaskId(t, urls, CodeForces{}, xContestId, xTaskId)
}

func TestCodeForcesContestAndTaskIdWithContestUrl(t *testing.T) {
	urls := []string{
		`https://codeforces.com/contest/1657/problem/A`,
		`https://www.codeforces.com/contest/1657/problem/A`,
	}
	xContestId, xTaskId := "1657", "a"
	RunTestContestAndTaskId(t, urls, CodeForces{}, xContestId, xTaskId)
}
