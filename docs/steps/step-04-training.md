# Step 04 — The training loop: epsilon-greedy exploration

| | |
|---|---|
| **Build** | `internal/train/train.go` — `Train` |
| **Tests** | `internal/train/train_test.go` — 2 tests |
| **Run** | `go test ./internal/train -run Step04 -v` |
| **You'll learn** | epsilon-greedy, episodes, learning curves |
| **Time** | ~45 min |

## 1. The story

Train by running many episodes: act with epsilon-greedy (explore vs exploit), update Q, track rewards. The agent
learns the optimal path.

## 2. Concept corner

- **Epsilon-greedy**: with prob ε pick a random action, else the best.
- **Learning curve**: reward per episode rises as Q-values converge.

## 3. Requirements

`Train(env, table, episodes, maxSteps, alpha, gamma, epsilon, seed) Result`.

## 4. What the tests pin down

- Runs the requested number of episodes; rewards finite.
- After training, success count >= 5 on a 5x5 grid.

## 5. Done when

`go test ./internal/train -run Step04 -race -v` → **ok**.

## 9. Next

Step 05 — evaluation and the optimal path (looking at the result).
