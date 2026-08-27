# Scenario

**Feature**: `tty-watch kill` terminates sessions and prunes registry

```
# kill detached session
tty-watch kill <id> -> DELETE ptywrap session -> SIGTERM owner -> remove registry json
```

## Preconditions

- `terminates-detached` starts a live detached session before kill.
- `prunes-unreachable` seeds a stale registry entry with dead listen addr.
- `serve-sigkill-no-orphan-child` detaches an HUP-immune sticky child, then
  SIGKILLs `__serve__` and asserts the PTY child is reaped (no orphan slot).

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if req.Bin == "" {
		t.Fatalf("kill setup: tty-watch binary not built")
	}
	return nil
}
```