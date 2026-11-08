// REFERENCE SOLUTION for Step 05. Try it yourself first!

// Package eval evaluates a trained Q-table: success rate, average reward, and the greedy optimal path.
package eval

import (
	"rl/internal/env"
	"rl/internal/qlearn"
)

// Result is the outcome of evaluating a policy. GIVEN.
type Result struct {
	SuccessRate float64
	AvgReward   float64
}

// Evaluate runs `episodes` greedy episodes and reports success rate and average reward. GIVEN.
func Evaluate(e *env.Env, table *qlearn.Table, episodes, maxSteps int) Result {
	successes := 0
	totalReward := 0.0
	for ep := 0; ep < episodes; ep++ {
		s := e.Reset()
		for step := 0; step < maxSteps; step++ {
			a := table.BestAction(s)
			r, _ := e.Step(s, a)
			totalReward += r.Reward
			s = r.Next
			if r.Done {
				successes++
				break
			}
		}
	}
	rate := 0.0
	if episodes > 0 {
		rate = float64(successes) / float64(episodes)
	}
	avg := 0.0
	if episodes > 0 {
		avg = totalReward / float64(episodes)
	}
	return Result{SuccessRate: rate, AvgReward: avg}
}

// OptimalPath greedily follows the learned policy from the start state and returns the path taken plus whether the
// goal was reached. GIVEN.
func OptimalPath(e *env.Env, table *qlearn.Table, maxSteps int) ([]env.State, bool) {
	path := []env.State{e.Reset()}
	s := e.Reset()
	for i := 0; i < maxSteps; i++ {
		a := table.BestAction(s)
		r, _ := e.Step(s, a)
		s = r.Next
		path = append(path, s)
		if r.Done {
			return path, true
		}
	}
	return path, false
}
