// REFERENCE SOLUTION for Step 01. Try it yourself first!

// Package env implements the grid-world environment: the "problem" an RL agent must solve.
package env

import "errors"

// ErrInvalidAction: the action is not in the action space. GIVEN.
var ErrInvalidAction = errors.New("env: invalid action")

// State is an integer index into the grid (row-major: y*width + x). GIVEN.
type State int

// Action is one of the four cardinal moves. GIVEN.
type Action int

// The four actions. GIVEN.
const (
	ActionUp Action = iota
	ActionDown
	ActionLeft
	ActionRight
)

// StepResult is what Env.Step returns. GIVEN.
type StepResult struct {
	Next   State
	Reward float64
	Done   bool
}

// Env is a grid world: width x height cells, a start, a goal, and (optional) obstacles. GIVEN.
type Env struct {
	width, height int
	start, goal   State
	obstacles     map[State]bool
}

// New builds a grid world. GIVEN.
func New(width, height int, start, goal State, obstacles []State) *Env {
	obs := map[State]bool{}
	for _, s := range obstacles {
		obs[s] = true
	}
	return &Env{width: width, height: height, start: start, goal: goal, obstacles: obs}
}

// NumStates returns the number of states (width*height). GIVEN.
func (e *Env) NumStates() int { return e.width * e.height }

// NumActions returns the number of actions (4). GIVEN.
func (e *Env) NumActions() int { return 4 }

// Reset returns the start state. GIVEN.
func (e *Env) Reset() State { return e.start }

// Start returns the start state. GIVEN.
func (e *Env) Start() State { return e.start }

// Goal returns the goal state. GIVEN.
func (e *Env) Goal() State { return e.goal }

// Coord converts a state index to (x, y). GIVEN.
func (e *Env) Coord(s State) (int, int) {
	x := int(s) % e.width
	y := int(s) / e.width
	return x, y
}

// Index converts (x, y) to a state index. GIVEN.
func (e *Env) Index(x, y int) State { return State(y*e.width + x) }

// Step executes action a in state s and returns the next state, reward and done flag.
func (e *Env) Step(s State, a Action) (StepResult, error) {
	if a < ActionUp || a > ActionRight {
		return StepResult{}, ErrInvalidAction
	}
	x, y := e.Coord(s)
	nx, ny := x, y
	switch a {
	case ActionUp:
		ny--
	case ActionDown:
		ny++
	case ActionLeft:
		nx--
	case ActionRight:
		nx++
	}
	// bounds check
	if nx < 0 || ny < 0 || nx >= e.width || ny >= e.height {
		return StepResult{Next: s, Reward: -1.0, Done: false}, nil
	}
	next := e.Index(nx, ny)
	if e.obstacles[next] {
		return StepResult{Next: s, Reward: -1.0, Done: false}, nil
	}
	if next == e.goal {
		return StepResult{Next: next, Reward: 10.0, Done: true}, nil
	}
	return StepResult{Next: next, Reward: -0.1, Done: false}, nil
}

// Actions returns the list of actions (all four; invalid moves are penalised in Step). GIVEN.
func (e *Env) Actions(s State) []Action {
	return []Action{ActionUp, ActionDown, ActionLeft, ActionRight}
}
