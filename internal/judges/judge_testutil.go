package judges

import "testing"

func RunTestContestAndTaskId(t *testing.T, urls []string, judge Judge, xContestId string, xTaskId string) {
	for _, url := range urls {
		t.Run(url, func(t *testing.T) {
			contestId, taskId, err := judge.ContestAndTaskId(url)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if contestId != xContestId {
				t.Fatalf("contestId is not correct: expected %s, got %s\n", xContestId, contestId)
			}
			if taskId != xTaskId {
				t.Fatalf("taskId is not correct: expected %s, got %s\n", xTaskId, taskId)
			}
		})
	}
}
