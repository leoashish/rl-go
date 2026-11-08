# Step 01 — The grid-world environment: the problem an RL agent solves

| | |
|---|---|
| **Build** | `internal/env/env.go` — `Env`, `Step`, `Reset` |
| **Tests** | `internal/env/env_test.go` — 6 tests |
| **Run** | `go test ./internal/env -run Step01 -v` |
| **You'll learn** | MDPs, states, actions, transitions, rewards, terminal states |
| **Time** | ~45 min |

## 1. The story

Reinforcement learning needs a **problem**: an environment with states, actions, and rewards. Step 01 builds a grid world
— a 5×5 grid, start at (0,0), goal at (4,4), with movement penalties and a terminal reward.

## 2. Concept corner

- An **MDP** is (states, actions, transition, reward, discount).
- **Terminal state**: reaching the goal ends the episode with a positive reward.
- **Invalid moves** (walls/out of bounds) keep you in place with a penalty.

## 3. Requirements

`New`, `NumStates`, `NumActions`, `Reset`, `Start`, `Goal`, `Coord`, `Index`, `Step`, `Actions`.

## 4. What the tests pin down

- Grid dimensions and action count.
- Reset returns start.
- Stepping into the goal → done + positive reward.
- Out of bounds / obstacle → stay + penalty.
- Actions returns all four moves.

## 5. Done when

`go test ./internal/env -run Step01 -race -v` → **ok**.

## 9. Next

Step 02 — the agent and random policy: how to act and roll out episodes.
