# Scenario

**Feature**: detach run prints `session-id: <id>` on stdout and exits promptly

```
# long-lived child; harness waits for parent exit 0 within 5s
harness -> tty-watch run --detach sleep 120 -> session-id: session-N\n -> exit 0
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-detach-prints-session-id"
	req.RunCommand = []string{"sleep", "120"}
	return nil
}
```