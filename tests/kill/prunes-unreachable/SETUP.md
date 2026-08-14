# Scenario

**Feature**: kill on unreachable server prunes stale registry idempotently

```
# stale registry with dead listen addr
harness seeds session-stale-1 -> tty-watch kill -> exit 0, registry pruned
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "kill-stale"
	return nil
}
```