# Scenario

**Feature**: Ctrl-C (`\x03`) forwards interrupt to PTY child

```
# trap handler prints marker when child receives INT
harness PTY -> tty-watch sh trap -> write \x03 -> TTY_WATCH_INTERRUPTED
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-ctrl-c"
	req.SendCtrlC = true
	req.RunCommand = []string{"sh", "-c", `trap 'echo TTY_WATCH_INTERRUPTED' INT; sleep 300`}
	return nil
}
```