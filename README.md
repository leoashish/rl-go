# RL (Go) — learn reinforcement learning from scratch by solving a grid world

A hands-on course in **reinforcement learning**, built the same way as [distkv-go](../distkv-go),
[blockchain-go](../blockchain-go), [kernel-go](../kernel-go) and [max-go](../max-go): **one step at a time, each with a
PRD, failing tests, stubs, collapsible hints, stretch challenges and a reference solution**.

Today you are building **RL** — a miniature reinforcement-learning agent that solves a grid world from first principles,
in Go. Step 01 builds the **environment** (the problem). Step 02 builds the **agent** (random policy + rollout). Step 03
builds **tabular Q-learning** (the Bellman update). Step 04 builds the **training loop** (epsilon-greedy). Step 05
builds **evaluation** (success rate + optimal path).

The goal: understand *how RL actually works* — from MDPs and Q-values to a trained agent that finds the optimal path —
by building miniature, tested pieces yourself. **And see the result:** `go run ./cmd/rl` prints the learning outcome and
the optimal path.

## What you are building, and why this shape

RL = agent + environment + reward. The smallest "real" RL system is tabular Q-learning on a grid world, and it contains
every lesson in order:

1. **Environment (MDP)** — states, actions, transitions, rewards. (Step 01.)
2. **Agent & policy** — how to act; episode rollouts. (Step 02.)
3. **Q-learning** — the Bellman update; learning Q(s,a). (Step 03.)
4. **Training loop** — epsilon-greedy exploration over many episodes. (Step 04.)
5. **Evaluation** — greedy success rate, average reward, optimal path. (Step 05.)

By the end you have an agent that **learns the optimal path** on a grid — **so you know where the textbooks get it
right and where they make trade-offs you haven't felt yet**.

## Start here

```bash
cd rl-go
go version                      # 1.22 or newer
go test ./...                   # everything fails CLEANLY at first: no panics, no hangs. That is the starting state.
go test ./internal/env -run Step01 -v
```

Work through a step:

1. Read the step's PRD (`docs/steps/step-01-environment.md`).
2. Open the tests (`internal/<package>/<step>_test.go`) — they are the spec.
3. Open the stub.
4. Run only that step: `go test ./internal/env -run Step01 -v`.
5. When green, read **Reflect**, try a **Stretch**, run `go vet ./...` and `gofmt -l .`, compare with `_solutions/stepNN/…`.
6. `git commit`, next step.

**Stuck?** `bash scripts/apply-solution.sh 01`.

## Roadmap

| # | Package | Feature | Tests |
|---|---------|---------|------:|
| 01 | `env` | Grid-world environment: states, actions, rewards | 6 ✅ |
| 02 | `agent` | Random policy + episode rollout | 3 ✅ |
| 03 | `qlearn` | Tabular Q-learning: Bellman update | 4 ✅ |
| 04 | `train` | Epsilon-greedy training loop | 2 ✅ |
| 05 | `eval` | Evaluation: success rate + optimal path | 2 ✅ |

**17 tests in total** — all steps solved.

## What you'll learn, in order

1. **Step 01 — environment:** MDPs, states, actions, transitions, rewards. ✅
2. **Step 02 — agent:** policies, episodes, rollouts, returns. ✅
3. **Step 03 — Q-learning:** Q-values, Bellman equation, learning rate, discount. ✅
4. **Step 04 — training:** epsilon-greedy, episodes, learning curves. ✅
5. **Step 05 — evaluation:** greedy success rate, average reward, optimal path. ✅

## See the result

```bash
cd rl-go
go run ./cmd/rl
```

Prints the learning outcome (success rate, avg reward) and the learned optimal path.

## Project layout

```
rl-go/
├── README.md
├── go.mod
├── cmd/rl/main.go          # prints the learning result
├── docs/steps/             # PRDs
├── internal/
│   ├── todo/               # ErrNotImplemented for stubs
│   ├── env/                # step 01: grid world
│   ├── agent/              # step 02: policy + rollout
│   ├── qlearn/             # step 03: Q-table + Bellman
│   ├── train/              # step 04: training loop
│   └── eval/               # step 05: evaluation
├── _solutions/             # reference solutions (ignored by Go tool)
│   └── step01/ … step05/
└── scripts/                # apply-solution.sh / .ps1
```

## How this project was built

- All **17 step tests pass** under `-race` with `go vet` and `gofmt` clean.
- `go run ./cmd/rl` prints a successful learning result (success rate ~1.00, optimal path found).

## License

MIT — learn stuff.
