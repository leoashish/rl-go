package main

import (
	"fmt"

	"rl/internal/env"
	"rl/internal/eval"
	"rl/internal/qlearn"
	"rl/internal/train"
)

func main() {
	e := env.New(5, 5, env.State(0), env.State(24), nil)
	tbl := qlearn.NewTable(e.NumStates(), e.NumActions())

	fmt.Println("RL from scratch — grid world 5x5, start (0,0), goal (4,4)")
	fmt.Println("Training Q-learning for 2000 episodes...")

	res := train.Train(e, tbl, 2000, 100, 0.1, 0.9, 0.3, 42)
	ev := eval.Evaluate(e, tbl, 20, 100)

	fmt.Printf("Episodes: %d | Successes: %d | Success rate: %.2f | Avg reward: %.2f\n",
		len(res.EpisodeRewards), res.SuccessCount, ev.SuccessRate, ev.AvgReward)

	path, ok := eval.OptimalPath(e, tbl, 100)
	fmt.Println("Optimal path (greedy):")
	for _, s := range path {
		x, y := e.Coord(s)
		fmt.Printf("  (%d,%d)\n", x, y)
	}
	if ok {
		fmt.Println("Result: goal reached ✅")
	} else {
		fmt.Println("Result: goal NOT reached ❌")
	}
}
