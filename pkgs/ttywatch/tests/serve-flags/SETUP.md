# Scenario

**Feature**: pure L2 parse/build/default APIs for flag-only `__serve_*__` re-exec

```
# leaf sets Mode + inputs
Setup -> Run(Mode) -> ParseServeArgv | BuildServeArgv | BuildServeChildEnv | defaults | ServeOptionsFromArgv
doctest <- Response fields; Assert checks wire + policy B (no TTY_WATCH_* invent)
```

## Preconditions

- Package `github.com/xhd2015/tty-watch/pkgs/ttywatch` exports the P2 pure APIs
  listed in root `DOCTEST.md` (implementer; tests RED until then).
- No PTY, registry, or subprocess in this tree.
- Parallel-safe: no `t.Setenv` / `os.Setenv` / `t.Chdir`; defaults inject
  `UserHome` string only.

## Steps

1. Root validates `req` is non-nil; leaves set `Mode` and inputs along the chain.
2. `Run` dispatches on `Mode` to the pure API under test.
3. Leaf `Assert` checks wire fields, error presence, or policy B env rules.

## Context

- Parse args are **after** the serve token (`__serve_<slug>__` is not included).
- Build returns remainder **after** host binary prefix and serve token: starts
  with `sid`.
- Flag emit order (build): `--home`, `--registry-subdir`, `--keep-alive`,
  `--extra-path`…, `--env`…, `--unset-env`…, then `--`, then command.

```go
import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if req == nil {
		return fmt.Errorf("nil Request")
	}
	return nil
}

// equalStrings is a tree-wide helper for slice equality in asserts.
func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// hasEnvKey reports whether KEY= appears in an environ slice.
func hasEnvKey(env []string, key string) bool {
	prefix := key + "="
	for _, e := range env {
		if strings.HasPrefix(e, prefix) || e == key {
			return true
		}
	}
	return false
}

// envValue returns the last KEY=value assignment value, or ("", false).
func envValue(env []string, key string) (string, bool) {
	prefix := key + "="
	found := false
	val := ""
	for _, e := range env {
		if strings.HasPrefix(e, prefix) {
			found = true
			val = e[len(prefix):]
		}
	}
	return val, found
}

// joinHome mirrors DefaultTTYWatchHome expected layout for asserts.
func joinHome(userHome string) string {
	return filepath.Join(userHome, ".tty-watch")
}

// reexecKnobKeys are the four TTY_WATCH_* keys that must not be invent/clear on re-exec.
var reexecKnobKeys = []string{
	"TTY_WATCH_HOME",
	"TTY_WATCH_REGISTRY_SUBDIR",
	"TTY_WATCH_KEEP_ALIVE",
	"TTY_WATCH_EXTRA_PATHS",
}
```
