// REFERENCE SOLUTION for Step 03. Try it yourself first!

// Package qlearn implements tabular Q-learning: a Q-table and the Bellman update.
package qlearn

import "rl/internal/env"

// Table holds Q(s,a) for every state-action pair. GIVEN.
type Table struct {
	Q [][]float64
}

// NewTable returns a zero-initialized table. GIVEN.
func NewTable(numStates, numActions int) *Table {
	Q := make([][]float64, numStates)
	for i := range Q {
		Q[i] = make([]float64, numActions)
	}
	return &Table{Q: Q}
}

// Get returns Q(s,a). GIVEN.
func (t *Table) Get(s env.State, a env.Action) float64 { return t.Q[int(s)][int(a)] }

// Set sets Q(s,a). GIVEN.
func (t *Table) Set(s env.State, a env.Action, v float64) { t.Q[int(s)][int(a)] = v }

// BestAction returns the action with the highest Q(s,a). GIVEN.
func (t *Table) BestAction(s env.State) env.Action {
	best := env.Action(0)
	bestVal := t.Q[int(s)][0]
	for a := 1; a < len(t.Q[int(s)]); a++ {
		if t.Q[int(s)][a] > bestVal {
			bestVal = t.Q[int(s)][a]
			best = env.Action(a)
		}
	}
	return best
}

// MaxQ returns max_a Q(s,a). GIVEN.
func (t *Table) MaxQ(s env.State) float64 {
	mx := t.Q[int(s)][0]
	for a := 1; a < len(t.Q[int(s)]); a++ {
		if t.Q[int(s)][a] > mx {
			mx = t.Q[int(s)][a]
		}
	}
	return mx
}

// Update applies one Bellman update. GIVEN.
func Update(table *Table, s env.State, a env.Action, sNext env.State, reward float64, alpha, gamma float64, done bool) {
	target := reward
	if !done {
		target = reward + gamma*table.MaxQ(sNext)
	}
	old := table.Get(s, a)
	table.Set(s, a, old+alpha*(target-old))
}
