# Scenario

**Feature**: headless host stdout stays silent except for session-id line

```
# child echoes marker inside PTY; host must not show it
harness pipe -> tty-watch run --headless sh -c 'echo HEADLESS_MARKER; sleep 60'
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-headless-no-attach-output"
	req.RunCommand = []string{"sh", "-c", "echo HEADLESS_MARKER; sleep 60"}
	return nil
}
```