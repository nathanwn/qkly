package judges

import "testing"

func TestHackerRankContestAndTaskId(t *testing.T) {
	url := `https://hackerrank.com/contests/uqcs-codejam-2021/challenges/lucky-code`

	xContestId, xTaskId := "uqcs-codejam-2021", "lucky-code"

	hackerrank := HackerRank{}
	contestId, taskId := hackerrank.ContestAndTaskId(url)
	if contestId != xContestId {
		t.Fatalf(`contestId is not correct: expected "%s", got "%s"`, xContestId, contestId)
	}
	if taskId != xTaskId {
		t.Fatalf(`taskId is not correct: expected "%s", got "%s"`, xTaskId, taskId)
	}
}
