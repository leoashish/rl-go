package train_test

import (
	"math"
	"testing"

	"rl/internal/env"
	"rl/internal/qlearn"
	"rl/internal/train"
)

// Step 04 -- the Q-learning training loop. Run: go test ./internal/train -run Step04 -v

func TestStep04_Train_RunsTheRequestedNumberOfEpisodes(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	tbl := qlearn.NewTable(e.NumStates(), e.NumActions())
	res := train.Train(e, tbl, 50, 100, 0.1, 0.9, 0.3, 1)
	if len(res.EpisodeRewards) != 50 {
		t.Fatalf("len(EpisodeRewards) = %d, want 50", len(res.EpisodeRewards))
	}
	for i, r := range res.EpisodeRewards {
		if math.IsNaN(r) || math.IsInf(r, 0) {
			t.Fatalf("EpisodeRewards[%d] = %v, want finite", i, r)
		}
	}
}

func TestStep04_Train_LearnsToReachTheGoal(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	tbl := qlearn.NewTable(e.NumStates(), e.NumActions())
	res := train.Train(e, tbl, 2000, 100, 0.1, 0.9, 0.3, 42)
	if res.SuccessCount < 5 {
		t.Fatalf("SuccessCount = %d after 2000 episodes, want >= 5 (the agent should learn to reach the goal)", res.SuccessCount)
	}
}
