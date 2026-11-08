package qlearn_test

import (
	"testing"

	"rl/internal/env"
	"rl/internal/qlearn"
)

// Step 03 -- tabular Q-learning: table + Bellman update. Run: go test ./internal/qlearn -run Step03 -v

func TestStep03_NewTable_IsZeroInitialized(t *testing.T) {
	tbl := qlearn.NewTable(4, 2)
	if tbl.Get(env.State(0), env.Action(0)) != 0 {
		t.Fatalf("Q[0][0] = %v, want 0", tbl.Get(env.State(0), env.Action(0)))
	}
	if tbl.Get(env.State(3), env.Action(1)) != 0 {
		t.Fatalf("Q[3][1] = %v, want 0", tbl.Get(env.State(3), env.Action(1)))
	}
}

func TestStep03_BestAction_ReturnsTheArgmax(t *testing.T) {
	tbl := qlearn.NewTable(2, 4)
	tbl.Set(env.State(0), env.Action(2), 5.0)
	tbl.Set(env.State(0), env.Action(1), 1.0)
	if got := tbl.BestAction(env.State(0)); got != env.Action(2) {
		t.Fatalf("BestAction = %d, want 2", got)
	}
}

func TestStep03_Update_DoneTerminal_UsesOnlyTheReward(t *testing.T) {
	tbl := qlearn.NewTable(2, 2)
	qlearn.Update(tbl, env.State(0), env.Action(0), env.State(1), 10.0, 0.5, 0.9, true)
	got := tbl.Get(env.State(0), env.Action(0))
	if got != 5.0 {
		t.Fatalf("Q[0][0] = %v, want 5.0 (0 + 0.5*(10-0))", got)
	}
}

func TestStep03_Update_NotDone_UsesRewardPlusDiscountedMaxQ(t *testing.T) {
	tbl := qlearn.NewTable(2, 2)
	tbl.Set(env.State(1), env.Action(0), 8.0) // maxQ(s') = 8
	qlearn.Update(tbl, env.State(0), env.Action(0), env.State(1), 1.0, 0.5, 0.9, false)
	got := tbl.Get(env.State(0), env.Action(0))
	// 0 + 0.5*(1 + 0.9*8 - 0) = 0.5*(1+7.2) = 4.1
	want := 4.1
	if got != want {
		t.Fatalf("Q[0][0] = %v, want %v", got, want)
	}
}
