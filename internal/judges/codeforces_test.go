package judges

import "testing"

func TestCodeForcesContestAndTaskIdWithProblemsetUrl(t *testing.T) {
	url := `https://codeforces.com/problemset/problem/348/D`
	xContestId, xTaskId := "348", "d"

	codeforces := CodeForces{}
	contestId, taskId, err := codeforces.ContestAndTaskId(url)
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

func TestCodeForcesContestAndTaskIdWithContestUrl(t *testing.T) {
	url := `https://codeforces.com/contest/1657/problem/A`
	xContestId, xTaskId := "1657", "a"

	codeforces := CodeForces{}
	contestId, taskId, err := codeforces.ContestAndTaskId(url)
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
