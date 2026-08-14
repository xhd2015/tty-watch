# Scenario

**Feature**: detach host stdout stays silent except for session-id line

```
# child echoes marker inside PTY; host must not show it
harness -> tty-watch run --detach sh -c 'echo DETACH_MARKER; sleep 60' -> exit 0
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-detach-no-attach-output"
	req.RunCommand = []string{"sh", "-c", "echo DETACH_MARKER; sleep 60"}
	return nil
}
```