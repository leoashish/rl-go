package agent_test

import (
	"testing"

	"rl/internal/agent"
	"rl/internal/env"
)

// Step 02 -- the random policy and episode rollout. Run: go test ./internal/agent -run Step02 -v

func TestStep02_RandomPolicy_ActReturnsAValidAction(t *testing.T) {
	p := agent.NewRandomPolicy(1)
	for i := 0; i < 100; i++ {
		a := p.Act(env.State(0))
		if a < 0 || a > 3 {
			t.Fatalf("Act returned %d, want 0..3", a)
		}
	}
}

func TestStep02_Rollout_TerminatesWithinMaxSteps(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	p := agent.NewRandomPolicy(2)
	reward, steps, _ := agent.Rollout(e, p, 100)
	if steps > 100 {
		t.Fatalf("Rollout took %d steps, want <= 100", steps)
	}
	_ = reward
}

func TestStep02_Rollout_SucceedsWhenItReachesTheGoal(t *testing.T) {
	e := env.New(2, 1, env.State(0), env.State(1), nil)
	p := agent.NewRandomPolicy(3)
	succeeded := false
	for i := 0; i < 50; i++ {
		_, _, ok := agent.Rollout(e, p, 20)
		if ok {
			succeeded = true
			break
		}
	}
	if !succeeded {
		t.Fatal("no rollout succeeded in 50 tries on a 2-cell grid")
	}
}
