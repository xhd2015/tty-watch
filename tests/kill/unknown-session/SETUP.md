# Scenario

**Feature**: kill on unknown session id is idempotent (exit 0)

```
# missing registry entry
tty-watch kill session-99999 -> exit 0
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "kill-missing"
	req.KillID = "session-99999"
	return nil
}
```