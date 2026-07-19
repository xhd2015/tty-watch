# tty-watch SSoT Doctests (standalone module)

End-to-end + contract tests for **`github.com/xhd2015/tty-watch`**: CLI binary
(`./cmd/tty-watch` → `cli.Main` error-returning), `cli.Run` / `cli.ParseArgs`,
and `pkgs/ttywatch` pure helpers.

Migrated from agent-pro `script/tty-watch/tests` + `ttywatchtest` (P1). Overlapping
flag-validation / encode leaves keep the **newer in-process contract** (Mode=cli /
encode / run-config / parse-args). Live session phases keep the e2e harness
(`ttywatchtest`, build `./cmd/tty-watch`).

# DSN (Domain Specific Notion)

**Participants**

- **tty-watch binary** — built from `./cmd/tty-watch`; thin wrapper:
  `cli.Main(os.Args[1:], …)` → on error print `Error:` and `os.Exit(1)`.
- **cli.Main / ParseArgs / Run** — programmatic surface: returns `error` only;
  streams on args / `Config`; no package-level stream globals; no `os.Exit`.
- **Embedded ptywrap server** — HTTP/WS listener on `127.0.0.1:0` inside the
  owning `tty-watch` process for each session.
- **Registry** — `$TTY_WATCH_HOME/registry/` JSON files (`session-N.json` or
  `<custom-id>.json`) with flock-based id reservation.
- **PTY child** — command argv executed inside the embedded terminal session.
- **Host VT screen model** — server-side terminal emulator (cursor for CPR /
  `send --query-cursor`).
- **SGR mouse encoder** — pure `EncodeSGRClick(row,col,btn,release)` → CSI press
  (`M`) + optional release (`m`).
- **Test harness (`ttywatchtest`)** — isolated `TTY_WATCH_HOME`, PTY helpers,
  Ctrl-C / Ctrl-] detach, registry probes, send cat-capture, lockholder helper.

**Behaviors**

- Default `run` reserves `session-N`, starts serve, writer-attaches; **host stays
  silent** (no session-id on stdout/stderr).
- Registry exclusive flock timeout fails with multi-line stderr diagnostics
  (lock path, holders, process tree).
- `run --headless` prints `session-id: <id>`, blocks until child exit / Ctrl-C.
- `run --detach` prints session-id and exits 0 immediately; session survives.
- `--detach` and `--headless` mutually exclusive; flags stop at first command arg.
- Optional `run --session-id <id>`: pattern `[a-zA-Z0-9][a-zA-Z0-9._-]*`.
- `list` aligned table SESSION/UPTIME/WATCH/ATTACHED/COMMAND; prunes unreachable.
- `watch` observer (raw bytes, no stdin forward); `attach` multi-writer; `snapshot`
  sanitized; `kill` terminates / prunes.
- `send` text / `--click` / `--query-cursor` exclusive modes; flag validation
  before registry lookup.
- Contract Mode leaves call `cli.Main` / `cli.Run` / `cli.ParseArgs` /
  `EncodeSGRClick` **in-process** (no binary required for correctness; root still
  builds binary for shared setup).

## Version

0.0.2

## Decision Tree

```
[tty-watch SSoT tests]
 |
 +-- cli-main/                         # Mode=cli → cli.Main (contract)
 |    +-- help-root/                   (LEAF)
 |    +-- empty-args/                  (LEAF)
 |    +-- unknown-subcommand/          (LEAF)
 |
 +-- run-config/                       # Mode=run-config | parse-args (contract)
 |    +-- help/                        (LEAF)
 |    +-- unknown-command/             (LEAF)
 |    +-- send-click-missing-col/      (LEAF)
 |    +-- send-click-options/          (LEAF)
 |
 +-- run/                              # Phase e2e (live PTY / registry)
 |    +-- registers-session/ … headless/ detach/ custom-session-id/ …
 |
 +-- list/                             # Phase e2e
 +-- watch/                            # Phase e2e
 +-- attach/                           # Phase e2e
 +-- snapshot/                         # Phase e2e
 +-- kill/                             # Phase e2e
 |
 +-- send/
 |    +-- injects-verbatim/ no-suffix/ preserves-whitespace/ …
 |    +-- missing-args/{no-args,session-only}   # Mode=cli contract (split)
 |    +-- unknown-session/ terminal-unreachable/
 |    +-- click/inject/…               # Phase e2e cat-capture
 |    +-- click/validation/…           # Mode=cli contract (keep new)
 |    +-- query-cursor/{plain,json}-at-cup  # Phase e2e
 |    +-- query-cursor/{with-click,with-text} # Mode=cli contract
 |
 +-- unit/
      +-- encode-sgr-click/…           # Mode=encode contract
      +-- screen-snapshot-exit-marker/ normalize-tty-output/ …
```

## Ownership / merge notes

| Source | Disposition |
|--------|-------------|
| agent-pro e2e phases (run/list/watch/attach/snapshot/kill/send inject) | Migrated |
| agent-pro `send/click/validation/*`, `query-cursor/with-*`, encode | **Dropped** — replaced by contract Mode leaves |
| agent-pro `send/missing-args` (single leaf) | **Replaced** by `no-args` + `session-only` contract |
| agent-pro `errors/unknown-subcommand` | **Dropped** — `cli-main/unknown-subcommand` |
| Existing `cli-main/`, `run-config/` | **Kept** |
| `pkgs/ttywatch/*_test.go` + `pkgs/ttywatch/tests/naming` | Ported separately |

## How to Run

```sh
cd external/tty-watch-master-2026-07-19   # module root
doctest vet ./tests
doctest test ./tests/...
doctest test ./tests/cli-main/...
doctest test ./tests/run-config/...
doctest test ./tests/send/...
doctest test ./tests/run/...
doctest test ./tests/unit/encode-sgr-click/...
go test ./...
doctest vet ./pkgs/ttywatch/tests/naming
doctest test ./pkgs/ttywatch/tests/naming/...
```

```go
import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/xhd2015/tty-watch/cli"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
	"github.com/xhd2015/tty-watch/ttywatchtest"
)

// Request is the shared hybrid input model.
//
// Mode selects the Run path when non-empty:
//   - "cli":        cli.Main(Args, stdin, stdout, stderr) error
//   - "run-config":  cli.Run(Config) error
//   - "parse-args":  cli.ParseArgs(Args) (Config, error)
//   - "encode":      ttywatch.EncodeSGRClick(...)
// Empty Mode → phase e2e via ttywatchtest (requires Phase).
type Request struct {
	Mode string

	// Contract argv / config / encode
	Args          []string
	Stdin         string
	Config        cli.Config
	EncodeRow     int
	EncodeCol     int
	EncodeBtn     int
	EncodeRelease bool

	// E2E phase fields (mirror ttywatchtest.Request)
	Phase                         string
	Bin, TTYWatchHome, SessionID  string
	CustomSessionID               string
	RunCommand                    []string
	Detach, SendCtrlC, Background bool
	WatchProbe, SnapshotID, KillID string
	SendID, SendMessage            string
	AttachID, AttachInput, AttachInputB, AttachProbe string

	Click       bool
	QueryCursor bool
	ClickRow, ClickCol       int
	HasClickRow, HasClickCol bool
	Mouse                    int
	HasMouse                 bool
	NoRelease                bool
	JSON                     bool
	SendTextArgs             []string
}

// Response captures both contract API results and e2e binary outcomes.
type Response struct {
	// Contract
	Err          error
	ErrMsg       string
	Bytes        []byte
	ParsedConfig cli.Config

	// Streams / e2e
	ExitCode int
	Stdout, Stderr, Combined string
	SessionID                string
	RegistryExists           bool
	RegistryIDs              []string
	ListOutput               string
	SessionRunning           bool
	SnapshotText             string
	ContainsEscape           bool
	TimedOut                 bool
	Elapsed                  time.Duration
	GrokModesSeen            bool
	TTYCleanupOnDetach       bool
	PostDetachOutput         string
	SourceCheckOK            bool
	SourceCheckNote          string
	StdinRestoredBeforeCleanup bool
	KittyPopCleanupInSrc     bool
	AttachStdoutWriterNormalizesRawTTY bool
	InjectedBytes            []byte
	AltExitCode              int
	AltStderr                string
	AttachOutput             string
	AttachBOutput            string
	WatchOutput              string
	RunOutput                string
	LockPath                 string
	LockHolderPID            int
	LockHolderMarker         string
}

func errMsgOf(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// Run executes the behavior under test.
func Run(t *testing.T, req *Request) (*Response, error) {
	t.Helper()
	mode := req.Mode
	switch mode {
	case "encode":
		got := ttywatch.EncodeSGRClick(req.EncodeRow, req.EncodeCol, req.EncodeBtn, req.EncodeRelease)
		return &Response{Bytes: got, InjectedBytes: got}, nil

	case "cli":
		var stdout, stderr bytes.Buffer
		var stdin io.Reader = strings.NewReader(req.Stdin)
		if req.Stdin == "" {
			stdin = strings.NewReader("")
		}
		err := cli.Main(req.Args, stdin, &stdout, &stderr)
		out := stdout.String()
		errStr := stderr.String()
		resp := &Response{
			Err:      err,
			ErrMsg:   errMsgOf(err),
			Stdout:   out,
			Stderr:   errStr,
			Combined: out + errStr,
		}
		if err != nil {
			resp.ExitCode = 1
		}
		return resp, nil

	case "run-config":
		var stdout, stderr bytes.Buffer
		cfg := req.Config
		if cfg.Stdout == nil {
			cfg.Stdout = &stdout
		}
		if cfg.Stderr == nil {
			cfg.Stderr = &stderr
		}
		if cfg.Stdin == nil {
			cfg.Stdin = strings.NewReader(req.Stdin)
		}
		err := cli.Run(cfg)
		out := stdout.String()
		errStr := stderr.String()
		resp := &Response{
			Err:      err,
			ErrMsg:   errMsgOf(err),
			Stdout:   out,
			Stderr:   errStr,
			Combined: out + errStr,
		}
		if err != nil {
			resp.ExitCode = 1
		}
		return resp, nil

	case "parse-args":
		cfg, err := cli.ParseArgs(req.Args)
		return &Response{
			Err:          err,
			ErrMsg:       errMsgOf(err),
			ParsedConfig: cfg,
		}, nil

	case "":
		// Phase e2e path (headless/detach special-cased like agent-pro root).
		switch req.Phase {
		case "run-headless-prints-session-id":
			return phaseRunHeadlessPrintsSessionID(t, req)
		case "run-headless-no-attach-output":
			return phaseRunHeadlessNoAttachOutput(t, req)
		case "run-headless-waits-until-child-exits":
			return phaseRunHeadlessWaitsUntilChildExits(t, req)
		case "run-headless-ctrl-c-forwards-exits":
			return phaseRunHeadlessCtrlCForwardsExits(t, req)
		case "run-headless-ctrl-c-waiting-logs":
			return phaseRunHeadlessCtrlCWaitingLogs(t, req)
		case "run-headless-session-live-while-waiting":
			return phaseRunHeadlessSessionLiveWhileWaiting(t, req)
		case "run-headless-with-custom-session-id":
			return phaseRunHeadlessWithCustomSessionID(t, req)
		case "run-headless-stop-on-first-arg":
			return phaseRunHeadlessStopOnFirstArg(t, req)
		case "run-detach-prints-session-id":
			return phaseRunDetachPrintsSessionID(t, req)
		case "run-detach-no-attach-output":
			return phaseRunDetachNoAttachOutput(t, req)
		case "run-detach-session-survives-in-list":
			return phaseRunDetachSessionSurvivesInList(t, req)
		case "run-detach-with-custom-session-id":
			return phaseRunDetachWithCustomSessionID(t, req)
		case "run-detach-mutually-exclusive-with-headless":
			return phaseRunDetachMutuallyExclusiveWithHeadless(t, req)
		case "run-detach-stop-on-first-arg":
			return phaseRunDetachStopOnFirstArg(t, req)
		default:
			return mapPhaseResp(ttywatchtest.Run(t, toPhaseReq(req)))
		}

	default:
		t.Fatalf("unknown Request.Mode %q", mode)
		return nil, fmt.Errorf("unknown Request.Mode %q", mode)
	}
}
```
