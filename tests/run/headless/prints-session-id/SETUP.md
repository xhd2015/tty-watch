# Scenario

**Feature**: headless run prints `session-id: <id>` on stdout and skips attach

```
# long-lived child; harness reads first stdout line only
harness pipe -> tty-watch run --headless sleep 120 -> session-id: session-N\n
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-headless-prints-session-id"
	req.RunCommand = []string{"sleep", "120"}
	return nil
}
```