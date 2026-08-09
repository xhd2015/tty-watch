## Expected

- `err` is nil.
- Child env does **not** invent from opts:
  - no `TTY_WATCH_HOME=/from-opts-home`
  - no `TTY_WATCH_REGISTRY_SUBDIR=…`
  - no `TTY_WATCH_KEEP_ALIVE=…`
  - no `TTY_WATCH_EXTRA_PATHS=…`
- Ambient `TTY_WATCH_HOME=/ambient-home` is **not cleared** (no invent/clear policy).
- `PATH=/bin` still present.

## Errors

- None from `Run`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	if err != nil {
		t.Fatalf("BuildServeChildEnv: unexpected error: %v", err)
	}
	env := resp.ChildEnv

	// Must not invent opts-derived values for the four knobs.
	if v, ok := envValue(env, "TTY_WATCH_HOME"); ok && v == "/from-opts-home" {
		t.Fatalf("invented TTY_WATCH_HOME from opts: %v", env)
	}
	if hasEnvKey(env, "TTY_WATCH_REGISTRY_SUBDIR") {
		// Ambient base had none; invent would add it.
		t.Fatalf("must not invent TTY_WATCH_REGISTRY_SUBDIR: %v", env)
	}
	if hasEnvKey(env, "TTY_WATCH_KEEP_ALIVE") {
		t.Fatalf("must not invent TTY_WATCH_KEEP_ALIVE: %v", env)
	}
	if hasEnvKey(env, "TTY_WATCH_EXTRA_PATHS") {
		t.Fatalf("must not invent TTY_WATCH_EXTRA_PATHS: %v", env)
	}

	// Must not clear ambient TTY_WATCH_HOME.
	if v, ok := envValue(env, "TTY_WATCH_HOME"); !ok || v != "/ambient-home" {
		t.Fatalf("must not clear ambient TTY_WATCH_HOME; got ok=%v val=%q env=%v", ok, v, env)
	}
	if v, ok := envValue(env, "PATH"); !ok || v != "/bin" {
		t.Fatalf("PATH should pass through; got ok=%v val=%q env=%v", ok, v, env)
	}
}
```
