// REFERENCE SOLUTION for Step 02. Try it yourself first!

// Package agent implements the random policy and episode rollout used to evaluate the environment.
package agent

import (
	"math/rand"

	"rl/internal/env"
)

// Policy is the interface every (possibly learned) policy implements. GIVEN.
type Policy interface {
	Act(s env.State) env.Action
}

// RandomPolicy acts uniformly at random. GIVEN.
type RandomPolicy struct {
	rng *rand.Rand
}

// NewRandomPolicy returns a seeded random policy. GIVEN.
func NewRandomPolicy(seed int64) *RandomPolicy {
	return &RandomPolicy{rng: rand.New(rand.NewSource(seed))}
}

// Act picks a uniformly random action. GIVEN.
func (p *RandomPolicy) Act(s env.State) env.Action {
	return env.Action(p.rng.Intn(4))
}

// Rollout runs one episode from the start state until done or maxSteps.
func Rollout(e *env.Env, policy Policy, maxSteps int) (totalReward float64, steps int, success bool) {
	s := e.Reset()
	for steps = 0; steps < maxSteps; steps++ {
		a := policy.Act(s)
		res, err := e.Step(s, a)
		if err != nil {
			return totalReward, steps, false
		}
		totalReward += res.Reward
		s = res.Next
		if res.Done {
			return totalReward, steps + 1, true
		}
	}
	return totalReward, steps, false
}
