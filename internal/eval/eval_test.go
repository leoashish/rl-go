package eval_test

import (
	"testing"

	"rl/internal/env"
	"rl/internal/eval"
	"rl/internal/qlearn"
	"rl/internal/train"
)

// Step 05 -- evaluation and the optimal path. Run: go test ./internal/eval -run Step05 -v

func TestStep05_Evaluate_ReturnsRatesInRange(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	tbl := qlearn.NewTable(e.NumStates(), e.NumActions())
	train.Train(e, tbl, 500, 100, 0.1, 0.9, 0.3, 7)
	r := eval.Evaluate(e, tbl, 20, 100)
	if r.SuccessRate < 0 || r.SuccessRate > 1 {
		t.Fatalf("SuccessRate = %v, want in [0,1]", r.SuccessRate)
	}
}

func TestStep05_OptimalPath_ReachesTheGoalAfterTraining(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	tbl := qlearn.NewTable(e.NumStates(), e.NumActions())
	train.Train(e, tbl, 2000, 100, 0.1, 0.9, 0.3, 7)
	path, ok := eval.OptimalPath(e, tbl, 100)
	if !ok {
		t.Fatal("OptimalPath did not reach the goal after training")
	}
	if len(path) == 0 || path[0] != e.Start() {
		t.Fatalf("OptimalPath = %v, want a non-empty path starting at %d", path, e.Start())
	}
	if path[len(path)-1] != e.Goal() {
		t.Fatalf("OptimalPath last = %d, want goal %d", path[len(path)-1], e.Goal())
	}
}
