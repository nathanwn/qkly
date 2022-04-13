package judges

import "testing"

func TestJudgeId(t *testing.T) {
	xJudgeNames := make(map[Judge]string, 4)

	xJudgeNames[AtCoder{}] = "atcoder"
	xJudgeNames[CodeForces{}] = "codeforces"
	xJudgeNames[CSES{}] = "cses"
	xJudgeNames[HackerRank{}] = "hackerrank"

	for judge, xId := range xJudgeNames {
		t.Run(xId, func(t *testing.T) {
			if judge.Id() != xId {
				t.Fatalf("Wrong judge name: expected \"%s\", found %s\n", xId, judge.Id())
			}
		})
	}
}
