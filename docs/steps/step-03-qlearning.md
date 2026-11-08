# Step 03 — Tabular Q-learning: the Bellman update

| | |
|---|---|
| **Build** | `internal/qlearn/qlearn.go` — `Table`, `Update` |
| **Tests** | `internal/qlearn/qlearn_test.go` — 4 tests |
| **Run** | `go test ./internal/qlearn -run Step03 -v` |
| **You'll learn** | Q-values, Bellman equation, learning rate, discount |
| **Time** | ~45 min |

## 1. The story

Q-learning keeps a table `Q(s,a)` of expected returns and updates it with the Bellman equation:
`Q(s,a) ← Q(s,a) + α (r + γ max_a' Q(s',a') − Q(s,a))`.

## 2. Concept corner

- **Q(s,a)**: expected discounted return of taking action `a` in state `s`.
- **α (learning rate)**: how fast to update.
- **γ (discount)**: how much future rewards matter.
- **Terminal update**: when done, target is just `r`.

## 3. Requirements

`NewTable`, `Get`, `Set`, `BestAction`, `MaxQ`, `Update`.

## 4. What the tests pin down

- Table zero-initialized.
- `BestAction` = argmax.
- Terminal update: `Q ← Q + α(r − Q)`.
- Non-terminal: `Q ← Q + α(r + γ maxQ(s') − Q)`.

## 5. Done when

`go test ./internal/qlearn -run Step03 -race -v` → **ok**.

## 9. Next

Step 04 — the training loop: epsilon-greedy exploration.
