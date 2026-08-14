# Scenario

**Feature**: `tty-watch snapshot` fetches scrollback with sanitized print output

```
# snapshot mode via WS
tty-watch snapshot <id> -> attach_mode=snapshot -> sanitize ANSI/C0 -> stdout
```

## Preconditions

- `prints-sanitized` starts a detached session that emits ANSI-colored text.
- `unknown-session` uses a nonexistent session id.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if req.Bin == "" {
		t.Fatalf("snapshot setup: tty-watch binary not built")
	}
	return nil
}
```