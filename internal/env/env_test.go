package env_test

import (
	"testing"

	"rl/internal/env"
)

// Step 01 -- the grid-world environment. Run: go test ./internal/env -run Step01 -v

func TestStep01_New_ReturnsTheConfiguredGridWorld(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	if e.NumStates() != 25 {
		t.Fatalf("NumStates = %d, want 25", e.NumStates())
	}
	if e.NumActions() != 4 {
		t.Fatalf("NumActions = %d, want 4", e.NumActions())
	}
	if e.Start() != 0 {
		t.Fatalf("Start = %d, want 0", e.Start())
	}
	if e.Goal() != 24 {
		t.Fatalf("Goal = %d, want 24", e.Goal())
	}
}

func TestStep01_Reset_ReturnsTheStartState(t *testing.T) {
	e := env.New(5, 5, env.State(3), env.State(24), nil)
	if got := e.Reset(); got != 3 {
		t.Fatalf("Reset = %d, want 3", got)
	}
}

func TestStep01_Step_IntoTheGoal_ReturnsDoneWithPositiveReward(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	// from (3,4) step down? Actually goal is (4,4); from (4,3) step down reaches goal
	s := e.Index(4, 3)
	res, err := e.Step(s, env.ActionDown)
	if err != nil {
		t.Fatalf("Step: unexpected error: %v", err)
	}
	if !res.Done {
		t.Fatalf("Step into goal: Done = false, want true")
	}
	if res.Reward <= 0 {
		t.Fatalf("Step into goal: Reward = %v, want > 0", res.Reward)
	}
	if res.Next != 24 {
		t.Fatalf("Step into goal: Next = %d, want 24", res.Next)
	}
}

func TestStep01_Step_OutOfBounds_StaysAndPenalises(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	s := env.State(0) // (0,0)
	res, err := e.Step(s, env.ActionLeft)
	if err != nil {
		t.Fatalf("Step: unexpected error: %v", err)
	}
	if res.Done {
		t.Fatalf("Step out of bounds: Done = true, want false")
	}
	if res.Next != s {
		t.Fatalf("Step out of bounds: Next = %d, want %d (stay)", res.Next, s)
	}
	if res.Reward >= 0 {
		t.Fatalf("Step out of bounds: Reward = %v, want < 0", res.Reward)
	}
}

func TestStep01_Step_IntoAnObstacle_StaysAndPenalises(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), []env.State{env.State(1)})
	s := env.State(0)
	res, err := e.Step(s, env.ActionRight) // (1,0) is an obstacle
	if err != nil {
		t.Fatalf("Step: unexpected error: %v", err)
	}
	if res.Done {
		t.Fatalf("Step into obstacle: Done = true, want false")
	}
	if res.Next != s {
		t.Fatalf("Step into obstacle: Next = %d, want %d (stay)", res.Next, s)
	}
	if res.Reward >= 0 {
		t.Fatalf("Step into obstacle: Reward = %v, want < 0", res.Reward)
	}
}

func TestStep01_Actions_ReturnsAllFourMoves(t *testing.T) {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	acts := e.Actions(env.State(0))
	if len(acts) != 4 {
		t.Fatalf("Actions len = %d, want 4", len(acts))
	}
}
