# Step 02 — The agent and random policy: how to act

| | |
|---|---|
| **Build** | `internal/agent/agent.go` — `RandomPolicy`, `Rollout` |
| **Tests** | `internal/agent/agent_test.go` — 3 tests |
| **Run** | `go test ./internal/agent -run Step02 -v` |
| **You'll learn** | policies, episodes, rollouts, returns |
| **Time** | ~30 min |

## 1. The story

Before learning, an agent needs to **act** and **measure**. Step 02 builds a random policy and an episode rollout
that returns total reward, step count, and success.

## 2. Concept corner

- A **policy** maps state → action.
- An **episode** is a sequence of (s, a, r) until terminal or max steps.
- The **return** is the sum of rewards.

## 3. Requirements

`RandomPolicy` (seeded), `Act`, `Rollout(env, policy, maxSteps)`.

## 4. What the tests pin down

- `Act` returns a valid action (0..3).
- `Rollout` terminates within `maxSteps`.
- On a tiny grid, at least one rollout succeeds.

## 5. Done when

`go test ./internal/agent -run Step02 -race -v` → **ok**.

## 9. Next

Step 03 — tabular Q-learning: the Bellman update.
