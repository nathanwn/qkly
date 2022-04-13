package judges

import "testing"

func TestAtCoderContestAndTaskId(t *testing.T) {
	url := `https://atcoder.jp/contests/abc246/tasks/abc246_a`
	xContestId, xTaskId := "abc246", "a"

	atcoder := AtCoder{}
	contestId, taskId, err := atcoder.ContestAndTaskId(url)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if contestId != xContestId {
		t.Fatalf("contestId is not correct: expected %s, got %s\n", xContestId, contestId)
	}
	if taskId != xTaskId {
		t.Fatalf("taskId is not correct: expected %s, got %s\n", xTaskId, taskId)
	}
}
