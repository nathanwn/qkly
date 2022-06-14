package judges

import "testing"

func TestJudgeId(t *testing.T) {
	xJudgeNames := make(map[Judge]string)

	xJudgeNames[AtCoder{}] = "atcoder"
	xJudgeNames[CodeForces{}] = "codeforces"
	xJudgeNames[CSES{}] = "cses"
	xJudgeNames[HackerRank{}] = "hackerrank"
	xJudgeNames[SPOJ{}] = "spoj"

	for judge, xId := range xJudgeNames {
		t.Run(xId, func(t *testing.T) {
			if judge.Id() != xId {
				t.Fatalf("Wrong judge name: expected \"%s\", found %s\n", xId, judge.Id())
			}
		})
	}
}

func TestGetJudge(t *testing.T) {
	xJudges := make(map[string]string)

	xJudges["https://atcoder.jp/contests/abc246/tasks/abc246_a"] = "atcoder"
	xJudges["https://www.atcoder.jp/contests/abc246/tasks/abc246_a"] = "atcoder"
	xJudges["https://codeforces.com/contest/547/problem/C"] = "codeforces"
	xJudges["https://www.codeforces.com/contest/547/problem/C"] = "codeforces"
	xJudges["https://codeforces.com/problemset/problem/1697/F"] = "codeforces"
	xJudges["https://www.codeforces.com/problemset/problem/1697/F"] = "codeforces"
	xJudges["https://cses.fi/problemset/task/1680/"] = "cses"
	xJudges["https://www.cses.fi/problemset/task/1680/"] = "cses"
	xJudges["https://spoj.com/problems/SCUBADIV/"] = "spoj"
	xJudges["https://www.spoj.com/problems/SCUBADIV/"] = "spoj"

	for url, xJudgeId := range xJudges {
		t.Run(url, func(t *testing.T) {
			judge, err := GetJudge(url)
			if err != nil {
				t.Fatalf("Unsupported judge for url\"%s\"\n", url)
			}
			if judge.Id() != xJudgeId {
				t.Fatalf("Wrong judge name: expected \"%s\", found %s\n", xJudgeId, judge.Id())
			}
		})
	}
}
