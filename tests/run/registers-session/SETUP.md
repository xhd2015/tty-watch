# Scenario

**Feature**: default run reserves session id and writes registry entry

```
# start sleep session, detach, registry persists
harness PTY -> tty-watch sleep -> registry session-N.json
```

## Steps

1. Set phase `run-registers`.
2. Harness detaches after registry appears.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-registers"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```