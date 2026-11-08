// REFERENCE SOLUTION for Step 04. Try it yourself first!

// Package train runs the Q-learning training loop over an environment.
package train

import (
	"math/rand"

	"rl/internal/env"
	"rl/internal/qlearn"
)

// Result is the outcome of training. GIVEN.
type Result struct {
	EpisodeRewards []float64
	SuccessCount   int
}

// Train runs Q-learning for `episodes` episodes. GIVEN.
func Train(e *env.Env, table *qlearn.Table, episodes, maxSteps int, alpha, gamma, epsilon float64, seed int64) Result {
	rng := rand.New(rand.NewSource(seed))
	res := Result{EpisodeRewards: make([]float64, episodes)}
	for ep := 0; ep < episodes; ep++ {
		s := e.Reset()
		total := 0.0
		for step := 0; step < maxSteps; step++ {
			a := table.BestAction(s)
			if rng.Float64() < epsilon {
				a = env.Action(rng.Intn(e.NumActions()))
			}
			r, _ := e.Step(s, a)
			qlearn.Update(table, s, a, r.Next, r.Reward, alpha, gamma, r.Done)
			total += r.Reward
			s = r.Next
			if r.Done {
				res.SuccessCount++
				break
			}
		}
		res.EpisodeRewards[ep] = total
	}
	return res
}
