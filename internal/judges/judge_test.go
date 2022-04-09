package judges

import "testing"

func TestJudgeId(t *testing.T) {
	t.Run("CodeForces", func(t *testing.T) {
		judge := CodeForces{}
		xId := "codeforces"
		if judge.Id() != xId {
			t.Fatalf("Wrong judge name for Codeforces: expected \"%s\", found %s\n", xId, judge.Id())
		}
	})

	t.Run("AtCoder", func(t *testing.T) {
		judge := AtCoder{}
		xId := "atcoder"
		if judge.Id() != xId {
			t.Fatalf("Wrong judge name for AtCoder: expected \"%s\", found %s\n", xId, judge.Id())
		}
	})

	t.Run("HackerRank", func(t *testing.T) {
		judge := HackerRank{}
		xId := "hackerrank"
		if judge.Id() != xId {
			t.Fatalf("Wrong judge name for HackerRank: expected \"%s\", found %s\n", xId, judge.Id())
		}
	})
}
