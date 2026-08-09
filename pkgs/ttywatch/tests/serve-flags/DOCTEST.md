# ttywatch serve-flags (P2 flag-only re-exec)

Classic TDD **L2** doctests for pure parse/build/default helpers that make
`__serve_<slug>__` re-exec **flag-only**: session knobs travel after `sid` as
less-flags + `--` + pure command. Zero invent/clear of `TTY_WATCH_*` for
home/subdir/keep-alive/extra-paths on re-exec; policy B — those four are not
read ambiently for resolve (inject / pure defaults).

# DSN (Domain Specific Notion)

Flag-only serve re-exec: parent builds argv; serve child parses argv; child env
for the PTY agent is separate from re-exec knobs.

**Participants**

- **Caller / HeadlessRun** — builds re-exec remainder via `BuildServeArgv` from
  `HeadlessRunOptions` (home, subdir, keep-alive, extra-path, command env/unset,
  pure command).
- **ParseServeArgv** — pure: argv after serve token → `ServeArgv` (sid, flags,
  command after `--`).
- **BuildServeChildEnv** — re-exec child environ; **must not invent or clear**
  `TTY_WATCH_HOME`, `TTY_WATCH_REGISTRY_SUBDIR`, `TTY_WATCH_KEEP_ALIVE`,
  `TTY_WATCH_EXTRA_PATHS` (config is flags-only).
- **DefaultTTYWatchHome / DefaultRegistrySubdir** — pure defaults; inject
  `userHome`; no `os.Getenv` of the four knobs.
- **ServeOptionsFromArgv** — maps `ServeArgv` → `ServeOptions` including
  `CommandEnv` / `CommandUnset` for PTY spawn (`MergeProcessEnv` from P1,
  implementer replace).

**Behaviors**

- Wire: `<sid> [flags…] -- <command…>` (optional bare command without `--` when
  remaining tokens are non-flags).
- Flags: `--home`, `--registry-subdir`, `--keep-alive`, `--extra-path` (repeat),
  `--env KEY=VALUE` (repeat), `--unset-env KEY` (repeat).
- Omit home/subdir → empty fields; defaults applied via pure helpers only.
- Registry `command` = pure agent argv; env set/unset are separate fields.

## Version

0.0.2

## Decision Tree

```
pkgs/ttywatch/tests/serve-flags/          # L2 pure library APIs
├── parse/                                # ParseServeArgv
│   ├── ok/
│   │   ├── sid-command-after-dashdash    # sid + -- + command
│   │   ├── multi-env                     # repeated --env
│   │   ├── multi-unset-env               # repeated --unset-env
│   │   ├── all-serve-flags               # every flag together
│   │   └── bare-command-no-dashdash      # sid + bare command (no --)
│   └── err/
│       ├── empty-command                 # sid + -- with no command
│       └── bad-flag                      # unknown --flag
├── build/                                # BuildServeArgv / BuildServeChildEnv
│   ├── emits-flags-dashdash-command      # full opts → flags + -- + cmd
│   ├── omits-default-flags               # zero opts → sid -- cmd only
│   └── no-tty-watch-env-invent           # child env: no invent/clear of 4 keys
├── defaults/                             # pure default helpers (no getenv)
│   ├── home-from-user-base               # DefaultTTYWatchHome(userHome)
│   └── registry-subdir                   # DefaultRegistrySubdir == "registry"
└── child-env/                            # ServeArgv → ServeOptions mapping
    └── maps-to-serve-options             # CommandEnv/Unset + knobs map
```

Parameter ranking (most → least significant):

1. **Operation** — parse vs build vs defaults vs child-env mapping
2. **Outcome** (parse) — ok vs err
3. **Fixture shape** — which flags / error condition

## Assumed APIs (implementer)

```go
// pkgs/ttywatch

type ServeArgv struct {
    SessionID      string
    Home           string
    RegistrySubdir string
    KeepAlive      bool
    ExtraPaths     []string
    CommandEnv     []string // KEY=VALUE
    CommandUnset   []string // KEY
    Command        []string
}

func ParseServeArgv(args []string) (ServeArgv, error)

// BuildServeArgv: remainder after host binary + serve token:
//   <sid> [flags…] -- <command…>
func BuildServeArgv(sessionID string, opts HeadlessRunOptions) []string

// BuildServeChildEnv: re-exec environ. Must not invent or clear
// TTY_WATCH_HOME | TTY_WATCH_REGISTRY_SUBDIR | TTY_WATCH_KEEP_ALIVE | TTY_WATCH_EXTRA_PATHS.
func BuildServeChildEnv(base []string, opts HeadlessRunOptions) []string

func DefaultTTYWatchHome(userHome string) string // join userHome, ".tty-watch"
func DefaultRegistrySubdir() string              // "registry"

func ServeOptionsFromArgv(a ServeArgv) ServeOptions

// HeadlessRunOptions gains CommandEnv, CommandUnset []string
// ServeOptions gains CommandEnv, CommandUnset []string
```

## Test Index

| # | Leaf | Mode | Description |
|---|------|------|-------------|
| 1 | `parse/ok/sid-command-after-dashdash` | parse | Minimal: sid + `--` + command |
| 2 | `parse/ok/multi-env` | parse | Multiple `--env KEY=VALUE` preserved in order |
| 3 | `parse/ok/multi-unset-env` | parse | Multiple `--unset-env KEY` preserved |
| 4 | `parse/ok/all-serve-flags` | parse | All six flag kinds + command |
| 5 | `parse/ok/bare-command-no-dashdash` | parse | sid + bare non-flag command (no `--`) |
| 6 | `parse/err/empty-command` | parse | Error when command empty after `--` |
| 7 | `parse/err/bad-flag` | parse | Error on unknown flag |
| 8 | `build/emits-flags-dashdash-command` | build | Full opts emit flags + `--` + pure cmd |
| 9 | `build/omits-default-flags` | build | Empty knobs → only `sid -- cmd…` |
| 10 | `build/no-tty-watch-env-invent` | build-env | No invent/clear of four `TTY_WATCH_*` keys |
| 11 | `defaults/home-from-user-base` | defaults | Injected userHome → `…/.tty-watch` |
| 12 | `defaults/registry-subdir` | defaults | Default subdir is `"registry"` |
| 13 | `child-env/maps-to-serve-options` | child-env | Env/unset + knobs map onto ServeOptions |

## How to Run

```sh
cd external/tty-watch-master-2026-08-09   # module root
doctest vet ./pkgs/ttywatch/tests/serve-flags
doctest test ./pkgs/ttywatch/tests/serve-flags/...
doctest test -v ./pkgs/ttywatch/tests/serve-flags/parse/ok/sid-command-after-dashdash
```

Classic TDD: expect **RED** until implementer lands the APIs above.

```go
import (
	"fmt"
	"testing"

	"github.com/xhd2015/doctest/session"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

// Mode selects which pure surface Run exercises.
// parse | build | build-env | defaults | child-env
type Request struct {
	Mode string

	// parse
	Args []string

	// build / build-env
	SessionID string
	Opts      ttywatch.HeadlessRunOptions
	// BaseEnv is the parent environ for BuildServeChildEnv (build-env mode).
	BaseEnv []string

	// defaults
	UserHome string
}

type Response struct {
	// parse / child-env
	Parsed ttywatch.ServeArgv

	// build
	BuiltArgv []string

	// build-env
	ChildEnv []string

	// defaults
	DefaultHome   string
	DefaultSubdir string

	// child-env
	ServeOpts ttywatch.ServeOptions
}

func Run(t *testing.T, d *session.Doctest, req *Request) (*Response, error) {
	t.Helper()
	_ = d
	if req.Mode == "" {
		return nil, fmt.Errorf("req.Mode must be set by Setup")
	}
	resp := &Response{}

	switch req.Mode {
	case "parse":
		parsed, err := ttywatch.ParseServeArgv(req.Args)
		resp.Parsed = parsed
		return resp, err

	case "build":
		resp.BuiltArgv = ttywatch.BuildServeArgv(req.SessionID, req.Opts)
		return resp, nil

	case "build-env":
		resp.ChildEnv = ttywatch.BuildServeChildEnv(req.BaseEnv, req.Opts)
		return resp, nil

	case "defaults":
		resp.DefaultHome = ttywatch.DefaultTTYWatchHome(req.UserHome)
		resp.DefaultSubdir = ttywatch.DefaultRegistrySubdir()
		return resp, nil

	case "child-env":
		parsed, err := ttywatch.ParseServeArgv(req.Args)
		if err != nil {
			return resp, err
		}
		resp.Parsed = parsed
		resp.ServeOpts = ttywatch.ServeOptionsFromArgv(parsed)
		return resp, nil

	default:
		return nil, fmt.Errorf("unknown req.Mode %q", req.Mode)
	}
}
```
