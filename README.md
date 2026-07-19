# tty-watch

Embedded **PTY session manager**: run a command under a real PTY, keep the session alive while detached, then **watch**, **attach**, **snapshot**, **send** input, or **kill** it.

This repo is the **single source of truth** for the CLI and core library. Consumers (e.g. agent-pro) depend on the module; they do not vendor a fork.

```text
agent-pro / go-pkgs tests / skills
        │
        │  require / binary on PATH
        ▼
github.com/xhd2015/tty-watch
  cmd/tty-watch   CLI
  cli             Main / Run / ParseArgs (programmatic)
  pkgs/ttywatch   registry, inject, snapshot, headless helpers
        │
        ▼
github.com/xhd2015/dot-pkgs/go-pkgs/shell/ptywrap
```

## Install

```bash
go install github.com/xhd2015/tty-watch/cmd/tty-watch@latest
```

Or from a checkout:

```bash
go build -o "$(go env GOPATH)/bin/tty-watch" ./cmd/tty-watch
```

Requires a recent Go toolchain (see `go.mod`).

## CLI

```text
tty-watch run [--headless] [--detach] [--session-id <id>] <command> [args...]
tty-watch list
tty-watch watch <session-id>
tty-watch attach <session-id>
tty-watch snapshot <session-id>
tty-watch kill <session-id>
tty-watch send <session-id> <msg>...
tty-watch send <session-id> --click --row r --col c [--mouse btn] [--no-release] [--json]
tty-watch send <session-id> --query-cursor [--json]
```

| Command | Behavior |
|---------|----------|
| `run` | Start a session; default attaches as interactive writer. `--detach` prints `session-id: …` and exits (session stays alive). `--headless` prints id and waits for the child to exit. |
| `list` | Live sessions (id, uptime, watch/attach counts, command). |
| `watch` | Read-only observer (PTY output to stdout; no stdin inject). |
| `attach` | Multi-writer attach (input + resize). |
| `snapshot` | Sanitized scrollback / screen text for asserts. |
| `kill` | Tear down session and registry entry. |
| `send` | Inject text or SGR mouse / host cursor query (see below). |

Environment:

| Variable | Meaning |
|----------|---------|
| `TTY_WATCH_HOME` | Registry root (default under the process home layout used by the library). Isolate tests with a temp dir. |
| `TTY_WATCH_REGISTRY_SUBDIR` | Optional registry subdirectory name under home. |

### `run` modes

```bash
# Interactive (writer attach; host mirrors the PTY)
tty-watch run -- bash

# Detach: keep session for later watch/send/kill
tty-watch run --detach --session-id demo -- sleep 300
# → session-id: demo

# Headless: print id, wait until the child exits
tty-watch run --headless -- true
```

Flags for `run` stop at the first command token (`StopOnFirstArg`).

### `send`

**Text** (default):

```bash
tty-watch send demo $'hello\r'
```

**SGR mouse click** (coordinates are **0-based** cell indices; wire encoding is 1-based SGR):

```bash
tty-watch send demo --click --row 10 --col 67
tty-watch send demo --click --row 10 --col 67 --no-release   # press only
tty-watch send demo --click --row 10 --col 67 --json        # JSON ack on stdout
```

Default is press + release; `--mouse` selects SGR button id (default `0` = left).

**Host VT cursor** (does **not** inject CSI 6n into the child; reads host screen model):

```bash
tty-watch send demo --query-cursor
# → row=R col=C

tty-watch send demo --query-cursor --json
# → {"row":R,"col":C}
```

### Example session

```bash
export TTY_WATCH_HOME=$(mktemp -d)
tty-watch run --detach --session-id s1 -- bash -c 'echo HI; sleep 60'
tty-watch list
tty-watch snapshot s1
tty-watch send s1 $'echo more\r'
tty-watch kill s1
```

## Programmatic API

Thin binary:

```go
// cmd/tty-watch/main.go
if err := cli.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
}
```

Library entry points (`github.com/xhd2015/tty-watch/cli`):

| Function | Role |
|----------|------|
| `ParseArgs(args) (Config, error)` | Flag/mode parsing only |
| `Run(cfg Config) error` | **Config-driven** execution (preferred for callers) |
| `Main(args, stdin, stdout, stderr) error` | `ParseArgs` + attach IO + `Run` |

Prefer `Run` when you already have structured options:

```go
err := cli.Run(cli.Config{
    Home:    home,
    Command: "send",
    Send: &cli.SendOptions{
        Session: sid,
        Mode:    cli.SendModeClick,
        Row:     10,
        Col:     67,
    },
})
```

Lower-level session helpers live in `github.com/xhd2015/tty-watch/pkgs/ttywatch` (registry, inject, snapshot, headless run, SGR encode, etc.).

## Layout

```text
cmd/tty-watch/     binary entry
cli/               CLI parse + dispatch (Main / Run / Config)
pkgs/ttywatch/     core library
pkgs/agentdriver/  optional driver interface for headless re-exec
tests/             CLI doctests
ttywatchtest/      test harness helpers
```

## Tests

```bash
# package units
go test ./...

# CLI doctests (doctest CLI required)
doctest vet ./tests
doctest test ./tests/...

# naming helpers
doctest test ./pkgs/ttywatch/tests/naming/...
```

Some leaves are labeled `slow` or `real-grok` and are skipped unless selected with `--label`.

## Consumers

| Consumer | How |
|----------|-----|
| **agent-pro** | `require github.com/xhd2015/tty-watch`; import `pkgs/ttywatch`; install CLI for skills |
| **go-pkgs** (mouse headless) | **Binary only** — exec `tty-watch` on PATH / build from source; no Go module cycle with `ptywrap` |
| **wrk** dashboard TTY leaves | Binary on PATH |

## License

See repository license file if present; otherwise follow the owning project’s policy.
