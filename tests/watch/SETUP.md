# Scenario

**Feature**: `tty-watch watch` attaches as readonly observer

```
# observer mode: raw PTY bytes to stdout, no stdin forwarding
tty-watch watch <id> -> WS attach_mode=observer -> stdout bytes unchanged
```

## Preconditions

- Leaves start a detached session before invoking `watch`.
- `watch-stream` uses a looping echo command for deterministic markers.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if req.Bin == "" {
		t.Fatalf("watch setup: tty-watch binary not built")
	}
	return nil
}
```