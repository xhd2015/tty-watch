# Scenario

**Feature**: send on stale registry entry prunes and reports session not found

```
# stale registry with dead listen addr
harness seeds session-stale-1 -> tty-watch send session-stale-1 "hi" -> exit 1, registry pruned
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "send-stale"
	req.SendMessage = "hi"
	return nil
}
```