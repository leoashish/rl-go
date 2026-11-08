# Step 05 — Evaluation and the optimal path: looking at the result

| | |
|---|---|
| **Build** | `internal/eval/eval.go` — `Evaluate`, `OptimalPath` |
| **Tests** | `internal/eval/eval_test.go` — 2 tests |
| **Run** | `go test ./internal/eval -run Step05 -v` |
| **You'll learn** | greedy evaluation, success rate, optimal path extraction |
| **Time** | ~30 min |

## 1. The story

After training, evaluate: run greedy episodes, measure success rate and average reward, and extract the learned
optimal path.

## 2. Concept corner

- **Greedy evaluation**: ε = 0 (always pick best action).
- **Optimal path**: follow BestAction from the start until the goal.

## 3. Requirements

`Evaluate(env, table, episodes, maxSteps) Result`, `OptimalPath(env, table, maxSteps) ([]State, bool)`.

## 4. What the tests pin down

- Success rate in [0,1].
- After training, the greedy path reaches the goal.

## 5. Done when

`go test ./internal/eval -run Step05 -race -v` → **ok**.

## 9. Next

Run `go run ./cmd/rl` to see the learning result (visible output).
